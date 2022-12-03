package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xuender/oils/base"
)

// nolint: gochecknoinits
func init() {
	moveCmd := &cobra.Command{
		Use:   "move",
		Short: "整理文件/目录",
		Long:  `根据选择的分组策略移动文件/目录到设置的位置.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				InitMove(cmd).Move(base.Must1(cmd.Flags().GetInt("group")), args)
			}
		},
	}

	moveCmd.Flags().IntP("group", "g", 0, "使用分组设置")
	rootCmd.AddCommand(moveCmd)
}
