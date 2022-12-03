package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	scanCmd := &cobra.Command{
		Use:   "scan",
		Short: "扫描所有分组监听目录",
		Long:  `扫描所有分组监听目录，并根据配置整理文件/目录.`,
		Run: func(cmd *cobra.Command, args []string) {
			InitMove(cmd).Scan()
		},
	}
	rootCmd.AddCommand(scanCmd)
}
