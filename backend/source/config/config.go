package config 

import (
	"fmt"
	"os"
)

type Config struct {
	Server 		Server		`json:"server"`
	Database	Database	`json:"database"`
}

func NewConfiguration() *Config {
	return &Config{
		Server :Server {
			Port: "8080",
		},
		Database: Database{
			Driver: 	"pgx",
			Name:		os.Getenv("DB_NAME"),
			UserName:	os.Getenv("DB_USER"),
			Password:	os.Getenv("DB_PASSWORD"),
			Protocol:	"tcp",
			Host:		os.Getenv("DB_HOST"),
			Port:		os.Getenv("DB_PORT"),
		},
	}
}

func (c *Config) GetStringDBConnection() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Database.UserName,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}

