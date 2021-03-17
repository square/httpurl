package httpurl

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMustParse(t *testing.T) {
	panicF := func() {
		_ = MustParse("not a url")
	}
	require.Panics(t, panicF)

	u := MustParse("http://example.com/foo/bar?a=1&b=2")
	require.Equal(t, "http://example.com/foo/bar?a=1&b=2", u.String())
}

func TestAddQueryParam(t *testing.T) {
	u := MustParse("http://example.com/foo/bar?a=1&b=2")
	AddQueryParam(u, "b", 3)
	require.Equal(t, "http://example.com/foo/bar?a=1&b=2&b=3", u.String())
}

func TestSetQueryParam(t *testing.T) {
	u := MustParse("http://example.com/foo/bar?a=1&b=2")
	SetQueryParam(u, "b", 3)
	require.Equal(t, "http://example.com/foo/bar?a=1&b=3", u.String())
}

func TestAddPathSegment(t *testing.T) {
	u := MustParse("http://example.com/foo/bar?b=3")
	AddPathSegment(u, "../baz/meh")
	require.Equal(t, "http://example.com/foo/bar/..%252Fbaz%252Fmeh?b=3", u.String())

	u = MustParse("http://example.com/")
	AddPathSegment(u, "foo")
	AddPathSegment(u, "")
	AddPathSegment(u, "bar")

	require.Equal(t, "http://example.com/foo/bar", u.String())
}

func TestRemovePathSegment(t *testing.T) {
	u := MustParse("http://example.com/foo/bar/xyz?b=3")
	RemovePathSegment(u, 0)
	require.Equal(t, "http://example.com/bar/xyz?b=3", u.String())

	u = MustParse("http://example.com/foo/bar/xyz?b=3")
	RemovePathSegment(u, 1)
	require.Equal(t, "http://example.com/foo/xyz?b=3", u.String())

	u = MustParse("http://example.com/foo/bar/xyz?b=3")
	RemovePathSegment(u, 2)
	require.Equal(t, "http://example.com/foo/bar?b=3", u.String())

	u = MustParse("http://example.com/foo/bar/xyz?b=3")
	RemovePathSegment(u, 1000)
	require.Equal(t, "http://example.com/foo/bar/xyz?b=3", u.String())
}

func TestIsDomain(t *testing.T) {
	u := MustParse("http://www.example.com/foo/bar/xyz?b=3")
	require.False(t, IsDomain(u, "example"))
	require.False(t, IsDomain(u, "example.com"))
	require.True(t, IsDomain(u, "www.example.com"))
	require.False(t, IsDomain(u, "com"))
}

func TestIsSubdomainOf(t *testing.T) {
	u := MustParse("http://www.example.com/foo/bar/xyz?b=3")
	require.False(t, IsSubdomainOf(u, "example"))
	require.True(t, IsSubdomainOf(u, "example.com"))
	require.False(t, IsSubdomainOf(u, "www.example.com"))
	require.True(t, IsSubdomainOf(u, "com"))
}

func TestIsDomainOrSubdomainOf(t *testing.T) {
	u := MustParse("http://www.example.com/foo/bar/xyz?b=3")
	require.False(t, IsDomainOrSubdomainOf(u, "example"))
	require.True(t, IsDomainOrSubdomainOf(u, "example.com"))
	require.True(t, IsDomainOrSubdomainOf(u, "www.example.com"))
	require.True(t, IsDomainOrSubdomainOf(u, "com"))
}
