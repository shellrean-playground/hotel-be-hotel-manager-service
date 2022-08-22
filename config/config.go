package config

type Config struct {
	Server         ServerConfig   `yaml:"server"`
	Database       DatabaseConfig `yaml:"database"`
	ContextTimeout int            `yaml:"context_timeout"`
	Release        bool           `yaml:"release"`
}
