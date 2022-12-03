package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	metaCmd := &cobra.Command{
		Use:   "meta",
		Short: "文件或目录识别",
		Long:  `识别文件或目录，判断其类型和时间`,
		Run: func(cmd *cobra.Command, args []string) {
			service := InitMeta(cmd)

			for _, arg := range args {
				info := service.Info(arg)
				info.Output(cmd.OutOrStderr())
			}
		},
	}

	rootCmd.AddCommand(metaCmd)
}
