package watch

import (
	"os"
	"path/filepath"

	"github.com/xuender/fairy/meta"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

type Service struct {
	cfg *pb.Config
	ms  *meta.Service
}

func NewService(
	cfg *pb.Config,
	metaService *meta.Service,
) *Service {
	return &Service{
		cfg: cfg,
		ms:  metaService,
	}
}

func (p *Service) Run() {
	for _, path := range p.cfg.Path {
		dir := base.Must1(oss.Abs(path.Dir))
		logs.Debugw("run", "dir", dir)

		for _, entry := range base.Must1(os.ReadDir(dir)) {
			if p.cfg.IsIgnore(entry.Name()) {
				logs.Debugw("ignore", "name", entry.Name())

				continue
			}

			file := filepath.Join(dir, entry.Name())
			info := p.ms.Info(file)

			if target, has := path.Target[info.Meta.String()]; has {
				logs.Infow("mv", "path", file, "target", info.Target(target))
				Mv(file, info.Target(target))
			}
		}
	}
}

func (p *Service) Watch() {
	// TODO
}
