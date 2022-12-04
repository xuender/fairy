package move_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/fairy/move"
)

func TestHash(t *testing.T) {
	t.Parallel()

	assert.Equal(t, uint64(0x2f01937146ffb926), move.Hash("utils.go"))
}

func TestEqual(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.True(move.Equal("test/a.txt", "test/a.txt"))
	ass.True(move.Equal("test/a.txt", "test/b.txt"))
	ass.True(move.Equal("test/a.txt", "test/b.txt"))
	ass.False(move.Equal("test/a.txt", "test/c.txt"))
	ass.False(move.Equal("test/a.txt", "test/c.txt"))
	ass.False(move.Equal("test/b.txt", "test/c.txt"))
	ass.False(move.Equal("test/a.txt", "test/d.txt"))
	ass.False(move.Equal("test/a.txt", "test/e.txt"))
}
