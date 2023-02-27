package move

import (
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/xuender/fairy/meta"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

// Service 文件移动服务.
type Service struct {
	cfg *pb.Config
	ms  *meta.Service
}

// NewService 新建文件移动服务.
func NewService(
	cfg *pb.Config,
	metaService *meta.Service,
) *Service {
	return &Service{
		cfg: cfg,
		ms:  metaService,
	}
}

// Move 根据分组配置移动文件.
func (p *Service) Move(paths []string) {
	for _, path := range paths {
		p.move(base.Must1(oss.Abs(path)))
	}
}

func (p *Service) move(path string) {
	if lo.Contains(p.cfg.Ignore, path) {
		return
	}

	info := p.ms.Info(path)

	if dir, has := p.cfg.Dirs[info.Meta.String()]; has {
		if Move(path, info.Target(dir)) == nil {
			logs.I.Println("mv", path, info.Target(dir))
		}
	}
}

// Scan 扫描目录目录.
func (p *Service) Scan(paths ...string) {
	for _, path := range paths {
		path = base.Must1(oss.Abs(path))

		for _, entry := range base.Must1(os.ReadDir(path)) {
			p.move(filepath.Join(path, entry.Name()))
		}
	}
}
