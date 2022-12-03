package meta

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mholt/archiver/v3"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

type Service struct {
	times base.Set[string]
	dirs  map[string]pb.Meta
}

func NewService() *Service {
	return &Service{
		times: base.NewSet(
			"LICENSE",
			"typings.json",
		),
		dirs: map[string]pb.Meta{
			"go.mod":       pb.Meta_Golang,
			"pom.xml":      pb.Meta_Java,
			"package.json": pb.Meta_JavaScript,
		},
	}
}

// Info 文件目录信息获取.
func (p *Service) Info(path string) *Info {
	if newPath, err := oss.Abs(path); err == nil {
		path = newPath
	} else {
		return NewInfoError(path, err)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return NewInfoError(path, err)
	}

	if fileInfo.IsDir() {
		dirInfo := p.MatchDir(path)

		if dirInfo.Date.IsZero() {
			dirInfo.Date = fileInfo.ModTime()
		}

		return dirInfo
	}

	file, err := os.Open(path)
	if err != nil {
		return NewInfoError(path, err)
	}
	defer file.Close()

	meta, err := pb.GetMeta(path)
	if err != nil {
		return NewInfoError(path, err)
	}

	logs.Debugw("info", "path", path, "meta", meta)

	if meta == pb.Meta_Archive {
		return p.MatchArchive(path)
	}

	return &Info{Meta: meta, Date: fileInfo.ModTime(), Path: path}
}

func (p *Service) MatchArchive(path string) *Info {
	info := &Info{Meta: pb.Meta_Archive, Path: path}
	counts := map[pb.Meta]int{}

	if finfo, err := os.Stat(path); err == nil {
		info.Date = finfo.ModTime()
	}

	_ = archiver.Walk(path, func(entry archiver.File) error {
		if entry.IsDir() {
			return nil
		}

		if p.times.Has(entry.Name()) {
			info.Date = entry.ModTime()
		}

		if meta, has := p.dirs[entry.Name()]; has {
			info.Meta = meta
		}

		if meta, err := pb.GetMetaByReader(entry.ReadCloser); err == nil {
			logs.Debugw("MatchArchive", "path", path, "meta", meta)
			if count, has := counts[meta]; has {
				counts[meta] = count + 1
			} else {
				counts[meta] = 1
			}
		}

		return nil
	})

	if info.Meta != pb.Meta_Archive || len(counts) == 0 {
		return info
	}

	countMeta(counts, info)

	return info
}

func countMeta(counts map[pb.Meta]int, info *Info) {
	max := 0

	for meta, count := range counts {
		if meta != pb.Meta_Unknown && max < count {
			info.Meta = meta
			max = count
		}
	}
}

func (p *Service) MatchDir(dir string) *Info {
	info := &Info{Meta: pb.Meta_Unknown, Path: dir}

	entries, err := os.ReadDir(dir)
	if err != nil {
		info.Error = err

		return info
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if p.times.Has(entry.Name()) {
			if inf, err := entry.Info(); err == nil {
				info.Date = inf.ModTime()
			}
		}

		if meta, has := p.dirs[entry.Name()]; has {
			info.Meta = meta
		}
	}

	if info.Meta == pb.Meta_Unknown {
		return p.WalkDir(dir)
	}

	return info
}

func (p *Service) WalkDir(dir string) *Info {
	counts := map[pb.Meta]int{}

	info := &Info{Path: dir, Meta: pb.Meta_Unknown}
	info.Error = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if meta, err := pb.GetMeta(path); err == nil {
			if count, has := counts[meta]; has {
				counts[meta] = count + 1
			} else {
				counts[meta] = 1
			}
		}

		return nil
	})

	countMeta(counts, info)

	return info
}
