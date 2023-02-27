package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
)

// nolint: gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "fairy",
	Short: "文件整理精灵",
	Long:  `将文件移动到合适的位置.`,
	Run: func(cmd *cobra.Command, args []string) {
		// InitGUI(cmd).Run()
		if len(args) == 0 {
			InitMove(cmd).Scan(".")
		} else {
			InitMove(cmd).Scan(args...)
		}
	},
}

// Execute 执行.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logs.Log(err)
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
			logs.SetLevel(logs.Info)
		}

		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath(".")
			viper.AddConfigPath(base.Must1(os.UserHomeDir()))
			viper.SetConfigType("toml")
			viper.SetConfigName("fairy")
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			logs.I.Println("Config:", viper.ConfigFileUsed())
		}
	})
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件 (默认: $HOME/fairy.toml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "调试模式")
}
