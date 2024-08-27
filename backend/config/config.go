package config

type Config struct {
	DBConfig *DBConfig `env:", prefix=DB_"`
}

type DBConfig struct {
	Host     string `env:"HOST"`
	DbName   string `env:"DATABASE"`
	User     string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Port     int    `env:"PORT"`
}
