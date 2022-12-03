package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xuender/fairy/ui"
)

// nolint: gochecknoinits
func init() {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "初始化",
		Long:  `根据输入初始化配置文件.`,
		Run: func(cmd *cobra.Command, args []string) {
			ui.NewService().Init()
		},
	}
	rootCmd.AddCommand(initCmd)
}
