package config

import "os"

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	URL string
}

func LoadConfig() (Config, error) {
	server := ServerConfig{
		Port: os.Getenv("PORT"),
	}

	db := DBConfig{
		URL: os.Getenv("DB_URL"),
	}

	config := Config{
		Server: server,
		DB:     db,
	}

	return config, nil
}
