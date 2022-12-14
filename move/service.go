package move

import (
	"os"
	"path/filepath"
	"sync"
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
	cfg   *pb.Config
	ms    *meta.Service
	mutex sync.Mutex
}

// NewService 新建文件移动服务.
func NewService(
	cfg *pb.Config,
	metaService *meta.Service,
) *Service {
	return &Service{
		cfg:   cfg,
		ms:    metaService,
		mutex: sync.Mutex{},
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
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, group := range p.cfg.Group {
		dir := base.Must1(oss.Abs(group.Watch))
		logs.Debugw("scan", "dir", dir)

		entrys, err := os.ReadDir(dir)
		if err != nil {
			logs.Infow("不存在", "dir", dir, "error", err)

			continue
		}

		for _, entry := range entrys {
			if p.cfg.IsIgnore(entry.Name()) {
				logs.Debugw("ignore", "name", entry.Name())

				continue
			}

			file := filepath.Join(dir, entry.Name())
			info := p.ms.Info(file)

			if target, has := group.Meta[info.Meta.String()]; has {
				if err := Move(file, info.Target(target)); err == nil {
					logs.Infow("mv", "path", file, "target", info.Target(target))
				} else {
					logs.Warn(err)
				}
			}
		}
	}
}

// Watch 监听分组目录.
func (p *Service) Watch() {
	watcher := base.Must1(fsnotify.NewWatcher())
	defer watcher.Close()

	go p.toScan(watcher)

	paths := base.NewSet[string]()

	for _, group := range p.cfg.Group {
		path := base.Must1(oss.Abs(group.Watch))

		if _, err := os.Stat(path); err != nil {
			logs.Infow("不存在", "path", path, "err", err)
			paths.Add(path)

			continue
		}

		base.Must(watcher.Add(path))
		logs.Infow("watch", "path", path)
	}

	for range time.Tick(time.Second) {
		for path := range paths {
			if _, err := os.Stat(path); err == nil {
				base.Must(watcher.Add(path))
				logs.Infow("watch", "path", path)
				paths.Del(path)
			}
		}
	}
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

			if event.Has(fsnotify.Remove) {
				continue
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
