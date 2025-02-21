package config

type Config struct {
	DB          *DBConfig
	RedisConfig *RedisConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type RedisConfig struct {
	Host string
	Port string
}

func NewConfig() Config {
	return Config{
		DB: &DBConfig{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "aus1003",
			Name:     "go_clean_arch",
		},
		RedisConfig: &RedisConfig{
			Host: "localhost",
			Port: "6379",
		},
	}
}
