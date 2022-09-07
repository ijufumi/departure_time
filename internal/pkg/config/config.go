package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config is application configuration
type Config struct {
	Mail struct {
		Charset     string `env:"MAIL_CHARSET" envDefault:"UTF-8"`
		FromAddress string `env:"MAIL_FROM_ADDRESS"`

		SES struct {
			AwsAccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
			AwsSecretAccessKEY string `env:"AWS_SECRET_ACCESS_KEY"`
			AwsRegion          string `env:"AWS_REGION"`
		}
		SendGrid struct {
			SendGridAPIKEY string `env:"SENDGRID_API_KEY"`
		}
	}
}

// Load returns configuration made from environment variables
func Load() *Config {
	_ = godotenv.Load()
	c := Config{}
	_ = env.Parse(&c)

	return &c
}
