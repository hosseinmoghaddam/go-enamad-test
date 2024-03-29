package config

type Config struct {
	DB DB
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
