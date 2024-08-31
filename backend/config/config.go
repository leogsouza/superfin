package config

type Config struct {
	AppConfig *AppConfig `env:", prefix=APP_"`
	DBConfig  *DBConfig  `env:", prefix=DB_"`
}

type AppConfig struct {
	Port int `env:"PORT"`
}

type DBConfig struct {
	Host     string `env:"HOST"`
	DbName   string `env:"DATABASE"`
	User     string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Port     int    `env:"PORT"`
}

type TestConfig struct {
	Config *Config `env:", prefix=TEST_"`
}
