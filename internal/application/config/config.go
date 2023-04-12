package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type database struct {
	Host     string
	Name     string
	Password string
	Port     int
	Username string
}

type port struct {
	Grpc int
	Http int
}

type Config struct {
	Database database
	Port     port
}

func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("configs/.env")
	}
	if err != nil {
		err = godotenv.Load("../../configs/.env")
	}
	if err != nil {
		return nil, err
	}

	databasePort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		return nil, err
	}

	httpPort, err := strconv.Atoi(os.Getenv("PORT_HTTP"))
	if err != nil {
		return nil, err
	}

	grpcPort, err := strconv.Atoi(os.Getenv("PORT_GRPC"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: database{
			Host:     os.Getenv("DATABASE_HOST"),
			Name:     os.Getenv("DATABASE_NAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Port:     databasePort,
			Username: os.Getenv("DATABASE_USERNAME"),
		},
		Port: port{
			Grpc: grpcPort,
			Http: httpPort,
		},
	}, nil
}
