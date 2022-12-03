package ui

import (
	"github.com/manifoldco/promptui"
	"github.com/xuender/oils/logs"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (p *Service) Init() {
	logs.Debug("init")
	for {
		logs.Debug("run")

		if !p.Next() {
			return
		}
	}
}

func (p *Service) Next() bool {
	prompt := promptui.Select{
		Label: "是否继续创建精灵目录?",
		Items: []string{"继续", "退出"},
	}

	index, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return index == 0
}
