package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Mail struct {
		SES struct {
			AwsAccessKeyID string `env:"AWS_ACCESS_KEY_ID"`
			AwsSecretKEY   string `env:"AWS_SECRET_KEY"`
			AwsSesEndpoint string `env:"AWS_SES_ENDPOINT"`
			AwsRegion      string `env:"AWS_REGION"`
		}
		SendGrid struct {
			SendGridAPIKEY string `env:"SENDGRID_API_KEY"`
		}
	}
}

func Load() *Config {
	_ = godotenv.Load()
	c := Config{}
	_ = env.Parse(&c)

	return &c
}
