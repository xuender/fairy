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
	Long:  `根据配置文件将监听目录中的文件/目录移动到合适的位置.`,
	Run: func(cmd *cobra.Command, args []string) {
		InitWatch(cmd).Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// nolint: gochecknoinits
func init() {
	var cfgFile string

	cobra.OnInitialize(func() {
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
}
