package move

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/xuender/fairy/meta"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
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
func (p *Service) Move(num int, paths []string) {
	group := p.cfg.Group[num]

	for _, path := range paths {
		path = base.Must1(oss.Abs(path))
		info := p.ms.Info(path)

		if target, has := group.Meta[info.Meta.String()]; has {
			if Move(path, info.Target(target)) == nil {
				logs.Infow("mv", "path", path, "target", info.Target(target))
			}
		}
	}
}

// Scan 扫描分组目录.
func (p *Service) Scan() {
	for _, group := range p.cfg.Group {
		dir := base.Must1(oss.Abs(group.Watch))
		logs.Debugw("run", "dir", dir)

		for _, entry := range base.Must1(os.ReadDir(dir)) {
			if p.cfg.IsIgnore(entry.Name()) {
				logs.Debugw("ignore", "name", entry.Name())

				continue
			}

			file := filepath.Join(dir, entry.Name())
			info := p.ms.Info(file)

			if target, has := group.Meta[info.Meta.String()]; has {
				if Move(file, info.Target(target)) == nil {
					logs.Infow("mv", "path", file, "target", info.Target(target))
				}
			}
		}
	}
}

// Watch 监听分组目录.
func (p *Service) Watch() {
	logs.Info("watch")

	watcher := base.Must1(fsnotify.NewWatcher())
	defer watcher.Close()

	go p.toScan(watcher)

	for _, group := range p.cfg.Group {
		path := base.Must1(oss.Abs(group.Watch))

		base.Must(watcher.Add(path))
		logs.Debugw("watch", "path", path)
	}

	<-make(chan struct{})
}

func (p *Service) toScan(watcher *fsnotify.Watcher) {
	call, cancel := lo.NewDebounce(time.Second, p.Scan)

	defer cancel()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			logs.Debugw("watch", "event", event)

			if event.Has(fsnotify.Write) {
				logs.Info("modified file:", event.Name)
			}

			call()

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			logs.Error("error:", err)
		}
	}
}
