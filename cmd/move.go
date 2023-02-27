package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	moveCmd := &cobra.Command{
		Use:     "move",
		Short:   "整理文件/目录",
		Long:    `根据选择的分组策略移动文件/目录到设置的位置.`,
		Aliases: []string{"m"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				InitMove(cmd).Move(args)
			}
		},
	}

	rootCmd.AddCommand(moveCmd)
}
