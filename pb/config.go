package pb

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfig(cmd *cobra.Command) *Config {
	cfg := &Config{}

	if viper.ConfigFileUsed() != "" {
		if file, err := os.Open(viper.ConfigFileUsed()); err == nil {
			defer file.Close()

			toml.NewDecoder(file).Decode(cfg)
		}
	}

	if len(cfg.Ignore) == 0 {
		cfg.Ignore = []string{"README.md"}
	}

	return cfg
}

func (p *Config) IsIgnore(path string) bool {
	for _, ignore := range p.Ignore {
		if path == ignore {
			return true
		}
	}

	return false
}
