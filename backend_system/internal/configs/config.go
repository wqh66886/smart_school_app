package configs

type Config struct {
	Server struct {
		Port            string `yaml:"port"`
		ApplicationName string `mapstructure:"application-name"`
		Version         string `yaml:"version"`
		ContentPath     string `mapstructure:"content-path"`
	} `yaml:"server"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	Redis struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Password    string `yaml:"password"`
		Database    int    `yaml:"database"`
		Username    string `yaml:"username"`
		ConnectType string `mapstructure:"connect-type"`
	} `yaml:"redis"`
}
