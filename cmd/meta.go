package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	metaCmd := &cobra.Command{
		Use:   "meta",
		Short: "识别文件/目录",
		Long:  `识别文件/目录，判断其类型和时间`,
		Run: func(cmd *cobra.Command, args []string) {
			service := InitMeta(cmd)

			for _, arg := range args {
				service.Info(arg).Output(cmd.OutOrStderr())
			}
		},
	}

	rootCmd.AddCommand(metaCmd)
}
