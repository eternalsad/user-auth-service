package config

import "github.com/spf13/viper"

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

type Configuration struct {
	App *AppConfig
	DB  *DBConfig
}

type AppConfig struct {
	AppPort         string
	AppShutDownTime int
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Timeout  int
	SslMode  string
}

func NewConfiguration() (*Configuration, error) {
	err := initConfig()
	if err != nil {
		return nil, err
	}
	appConfig := &AppConfig{
		viper.GetString("server.port"),
		viper.GetInt("server.shutdown_time"),
	}
	dbConfig := &DBConfig{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.db_name"),
		viper.GetInt("db.timeout"),
		viper.GetString("db.ssl_mode"),
	}
	return &Configuration{appConfig, dbConfig}, nil
}
