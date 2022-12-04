package cmd

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
)

// nolint: gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "fairy",
	Short: "文件整理精灵",
	Long:  `监听配置的目录，将文件移动到合适的位置.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !base.Must1(cmd.Flags().GetBool("debug")) {
			logs.RotateLog("/var", "tmp")
		}

		InitMove(cmd).Watch()
	},
}

// Execute 执行.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// nolint: gochecknoinits
func init() {
	var (
		cfgFile string
		debug   bool
	)

	cobra.OnInitialize(func() {
		if !debug {
			logs.SetInfoLevel()
		}

		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath(base.Must1(homedir.Dir()))
			viper.AddConfigPath(".")
			viper.SetConfigType("toml")
			viper.SetConfigName("fairy")
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			logs.Infow("加载配置文件", "file", viper.ConfigFileUsed())
		}
	})
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件 (默认: $HOME/fairy.toml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "调试模式")
}
