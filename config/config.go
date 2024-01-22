package config

type ConfigDB struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Port     string
	DB       ConfigDB
	DebugMod bool
}

func InitConfig() *Config {
	cfgDb := ConfigDB{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	return &Config{
		Port: "8080",
		DB:   cfgDb}
}
