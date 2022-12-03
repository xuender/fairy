package watch

import (
	"os"
	"path/filepath"

	"github.com/xuender/oils/oss"
)

func Mv(path, dir string) error {
	dir, err := oss.Abs(dir)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, oss.DefaultDirFileMod); err != nil {
		return err
	}

	return os.Rename(path, filepath.Join(dir, filepath.Base(path)))
}
