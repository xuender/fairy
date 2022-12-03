package meta_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/fairy/meta"
	"github.com/xuender/oils/base"
)

func TestInfo_Target(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	info := &meta.Info{Created: base.Pass1(time.Parse("20060102150405", "20060102150405"))}

	ass.Equal("aa", info.Target("aa"))
	ass.Equal("[2006]", info.Target("[%yyyy]"))
	ass.Equal("[06]", info.Target("[%Yy]"))
	ass.Equal("[01]", info.Target("[%mm]"))
	ass.Equal("[2]", info.Target("[%d]"))
}
