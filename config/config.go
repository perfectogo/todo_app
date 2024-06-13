package config

type Config struct {
	AppNeme    string
	HttpPort   string
	DbUser     string
	DbUserPass string
	DbHost     string
	DbPort     int
	DbName     string
}

func Load() *Config {

	var Config *Config = &Config{}

	Config.AppNeme = "todo_app"
	Config.HttpPort = ":8080"

	Config.DbName = "postgres"
	Config.DbUserPass = "postgres"
	Config.DbHost = "localhost"
	Config.DbPort = 5432
	Config.DbName = "postgres"

	return Config
}
