package models

type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
}

type JWT struct {
	SecretKey   string `mapstructure:"SECRET_KEY"`
	ExpiresTime uint8  `mapstructure:"EXPIRES_TIME"`
}

type Config struct {
	Port     int `mapstructure:"PORT"`
	Postgres struct {
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
		Host     string `mapstructure:"HOST"`
		Database string `mapstructure:"DATABASE"`
		Port     string `mapstructure:"PORT"`
	} `mapstructure:"POSTGRES"`
	JWT JWT `mapstructure:"JWT"`
}
