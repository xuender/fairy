package move

import (
	"bufio"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"

	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

// Move 移动文件.
func Move(path, dir string) error {
	dir, err := oss.Abs(dir)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, oss.DefaultDirFileMod); err != nil {
		return err
	}

	name := filepath.Base(path)
	ext := filepath.Ext(name)
	target := filepath.Join(dir, name)
	index := 2

	for {
		if !oss.Exist(target) {
			logs.Debugw("rename", "path", path, "target", target)

			return os.Rename(path, target)
		}

		if Equal(path, target) {
			logs.Debugw("remove", "path", path, "target", target)

			return os.Remove(path)
		}

		target = filepath.Join(dir, fmt.Sprintf("%s-%d%s", name[:len(name)-len(ext)], index, ext))
		index++
	}
}

func Equal(file1, file2 string) bool {
	if file1 == file2 {
		return true
	}

	info1, err1 := os.Stat(file1)
	info2, err2 := os.Stat(file2)

	if !errors.Is(err1, err2) || err1 != nil {
		return false
	}

	if info1.Size() != info2.Size() {
		return false
	}

	return Hash(file1) == Hash(file2)
}

func Hash(path string) uint64 {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	hash := fnv.New64a()

	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return hash.Sum64()
		}

		if num == 0 {
			return hash.Sum64()
		}

		hash.Write(buf[:num])
	}
}