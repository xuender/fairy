package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/startup"
)

// nolint: gochecknoinits
func init() {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "设置开机自启动",
		Long:  `设置开机自启动，实时监听分组目录.`,
		Run: func(cmd *cobra.Command, args []string) {
			base.Must(startup.Install())
			logs.Info("Fairy installed.")
		},
	}
	rootCmd.AddCommand(installCmd)
}
