package config

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Timezone string `yaml:"timezone"`
}
