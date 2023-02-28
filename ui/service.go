package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/manifoldco/promptui"
	"github.com/pelletier/go-toml/v2"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

type Service struct {
	cfg *pb.Config
}

func NewService(cfg *pb.Config) *Service {
	return &Service{cfg: cfg}
}

func (p *Service) Init() {
	p.cfg.Dirs = map[string]string{}

	path := p.Prompt("请输入配置文件保存的目录", "./fairy.toml")

	logs.D.Println("input", "dir", path)

	for {
		meta := p.SelectMeta()
		p.cfg.Dirs[meta.String()] = p.Prompt(fmt.Sprintf("输入 %v 类型目录设置", meta), fmt.Sprintf("~/%v/$yyyy/$mm/$dd", meta))

		if !p.Confirm("继续设置新类型") {
			break
		}
	}

	logs.I.Println(p.cfg)

	p.Save(path)
}

func (p *Service) Save(path string) {
	config := viper.ConfigFileUsed()

	if config == "" {
		home := lo.Must1(os.UserHomeDir())
		config = filepath.Join(home, "fairy.toml")
	}

	if path != "" {
		config = lo.Must1(oss.Abs(path))
	}

	file := lo.Must1(os.Create(config))
	defer file.Close()

	encoder := toml.NewEncoder(file)
	_ = encoder.Encode(p.cfg)

	logs.I.Println("保存:", config)
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

	return lo.Must1(prompt.Run())
}

func (p *Service) Confirm(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	_, err := prompt.Run()

	return err == nil
}
