package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	scanCmd := &cobra.Command{
		Use:     "scan",
		Short:   "扫描指定目录",
		Long:    `扫描指定目录，并根据配置整理文件/目录.`,
		Aliases: []string{"s"},
		Run: func(cmd *cobra.Command, args []string) {
			InitMove(cmd).Scan(args)
		},
	}

	rootCmd.AddCommand(scanCmd)
}
