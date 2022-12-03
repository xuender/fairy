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
	p.cfg.Group = []*pb.Group{}

	for {
		group := &pb.Group{Meta: map[string]string{}}
		watch := p.Prompt("请输入要监听的精灵目录", "~/fairy")
		group.Watch = watch

		logs.Debugw("input", "dir", watch)

		p.cfg.Group = append(p.cfg.Group, group)

		for {
			meta := p.SelectMeta()
			target := p.Prompt(fmt.Sprintf("输入 %v 类型目录设置", meta), fmt.Sprintf("~/%v/$yyyy/$mm/$dd", meta))
			group.Meta[meta.String()] = target

			if !p.Confirm("继续设置新类型") {
				break
			}
		}

		if !p.Confirm("是否创建下一个分组") {
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
	_ = encoder.Encode(p.cfg)
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

func (p *Service) Prompt(label, def string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: def,
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
