package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/manifoldco/promptui"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
)

type Service struct {
	cfg *pb.Config
}

func NewService(cfg *pb.Config) *Service {
	return &Service{cfg: cfg}
}

func (p *Service) Init() {
	p.cfg.Path = []*pb.Path{}

	for {
		path := &pb.Path{Target: map[string]string{}}
		dir := p.Prompt("请输入精灵目录")
		path.Dir = dir

		logs.Debugw("input", "dir", dir)

		p.cfg.Path = append(p.cfg.Path, path)

		for {
			meta := p.SelectMeta()
			target := p.Prompt(fmt.Sprintf("输入 %v 类型目标目录", meta))
			path.Target[meta.String()] = target

			if !p.Confirm("是否继续设置目标目录") {
				break
			}
		}

		if !p.Confirm("是否创建下一个精灵目录") {
			break
		}
	}

	logs.Info(p.cfg)

	p.Save()
}

func (p *Service) Save() {
	config := viper.ConfigFileUsed()

	if config == "" {
		home := base.Must1(homedir.Dir())
		config = filepath.Join(home, "fairy.toml")
	}

	file := base.Must1(os.Create(config))
	defer file.Close()

	encoder := toml.NewEncoder(file)
	encoder.Encode(p.cfg)
}

func (p *Service) SelectMeta() pb.Meta {
	keys := make([]int32, 0, len(pb.Meta_name))

	for k := range pb.Meta_name {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	items := make([]string, len(pb.Meta_name))

	for i, key := range keys {
		items[i] = pb.Meta_name[key]
	}

	prompt := promptui.Select{
		Label: "选择文件类型",
		Items: items,
	}

	index, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return pb.Meta(index)
}

func (p *Service) Prompt(label string) string {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if s == "" {
				return ErrDirEmpty
			}

			return nil
		},
	}

	return base.Must1(prompt.Run())
}

func (p *Service) Confirm(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	_, err := prompt.Run()

	return err == nil
}
