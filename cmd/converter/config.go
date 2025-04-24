package main
import (
    "github.com/spf13/viper"
)
// Config holds converter configuration
type Config struct {
    OutputMD  string mapstructure:"output_md"
    OutputMMD string mapstructure:"output_mmd"
}
// LoadConfig reads configuration from converter.yaml
func LoadConfig(configPath string) (Config, error) {
    var cfg Config
    viper.SetConfigFile(configPath)
    viper.SetConfigType("yaml")
    viper.SetDefault("output_md", "usecases.md")
    viper.SetDefault("output_mmd", "diagrams.mmd")
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            return cfg, nil // Use defaults if file not found
        }
        return cfg, err
    }
    if err := viper.Unmarshal(&cfg); err != nil {
        return cfg, err
    }
    return cfg, nil
}
