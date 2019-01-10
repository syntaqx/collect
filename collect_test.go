package collect

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	vs := []string{"first", "second", "third"}

	assert.Equal(0, Index(vs, "first"))
	assert.Equal(1, Index(vs, "second"))
	assert.Equal(2, Index(vs, "third"))
	assert.Equal(-1, Index(vs, "fourth"))
}

func TestInclude(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	vs := []string{"foo", "bar"}

	assert.True(Include(vs, "foo"))
	assert.True(Include(vs, "bar"))
	assert.False(Include(vs, "baz"))
}

func TestAny(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	vs := []string{"foo", "bar"}

	assert.True(Any(vs, func(s string) bool {
		return true
	}))

	assert.True(Any(vs, func(s string) bool {
		if s == "bar" {
			return false
		}
		return true
	}))

	assert.False(Any(vs, func(s string) bool {
		return false
	}))
}

func TestAll(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	vs := []string{"foo", "bar"}

	assert.True(All(vs, func(s string) bool {
		return true
	}))

	assert.False(All(vs, func(s string) bool {
		if s == "bar" {
			return false
		}
		return true
	}))

	assert.False(All(vs, func(s string) bool {
		return false
	}))
}

func TestFilter(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	vs := []string{"apple", "banana", "cherry"}

	a := Filter(vs, func(s string) bool {
		return s[0] == 'a'
	})

	b := Filter(vs, func(s string) bool {
		return s[0] == 'b'
	})

	c := Filter(vs, func(s string) bool {
		return s[0] == 'c'
	})

	assert.Len(a, 1)
	assert.Equal("apple", a[0])
	assert.Len(b, 1)
	assert.Equal("banana", b[0])
	assert.Len(c, 1)
	assert.Equal("cherry", c[0])
}

func TestMap(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	vs := []string{"lorem", "ipsum", "dolor", "sit"}
	ex := `"lorem","ipsum","dolor","sit"`

	assert.Equal(ex, strings.Join(Map(vs, func(s string) string {
		return fmt.Sprintf(`"%s"`, s)
	}), ","))
}
