package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	initCmd := &cobra.Command{
		Use:     "init",
		Short:   "初始化",
		Long:    `根据输入初始化配置文件.`,
		Aliases: []string{"i"},
		Run: func(cmd *cobra.Command, args []string) {
			InitUI(cmd).Init()
		},
	}
	rootCmd.AddCommand(initCmd)
}
