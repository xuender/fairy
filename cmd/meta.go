package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xuender/fairy/meta"
)

// nolint: gochecknoinits
func init() {
	metaCmd := &cobra.Command{
		Use:   "meta",
		Short: "文件或目录识别",
		Long:  `识别文件或目录，判断其类型和将要归档的位置`,
		Run: func(cmd *cobra.Command, args []string) {
			service := meta.NewService()

			for _, arg := range args {
				info := service.Info(arg)
				info.Output(cmd.OutOrStderr())
			}
		},
	}

	rootCmd.AddCommand(metaCmd)
}
