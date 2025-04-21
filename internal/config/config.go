package config

import (
	"fmt"
	"os"
)

type Config struct {
	MongoURL              string `json:"mongo_url"`
	Port                  string `json:"port"`
	Env                   string `json:"env"`
	Host                  string `json:"host"`
	Secret                string `json:"secret"`
	Timeout               int    `json:"timeout"`
	SSL                   bool   `json:"ssl"`
	SSLPort               string `json:"ssl_port"`
	SSLHost               string `json:"ssl_host"`
	SSLKey                string `json:"ssl_key"`
	SSLCert               string `json:"ssl_cert"`
	SSLCA                 string `json:"ssl_ca"`
	SSLMode               string `json:"ssl_mode"`
	SSLVerify             bool   `json:"ssl_verify"`
	SSLInsecureSkipVerify bool   `json:"ssl_insecure_skip_verify"`
	SSLClientAuth         bool   `json:"ssl_client_auth"`
	SSLClientCA           string `json:"ssl_client_ca"`
	SSLClientCert         string `json:"ssl_client_cert"`
	Key                   string `json:"key"`
	ServiceName           string `json:"service_name"`
	TokenTelegram         string `json:"token_telegram"`
	ChatID                string `json:"chat_id"`
	AuthSecret            string `json:"auth_secret"`
}

func NewConfig() *Config {
	return &Config{
		MongoURL:      getEnv("MONGO_URL", "mongodb://localhost:27017"),
		Env:           getEnv("ENV", "development"),
		Host:          getEnv("HOST", "localhost"),
		Secret:        getEnv("MONGO_SECRET", "password"),
		Key:           getEnv("MONGO_KEY", "admin"),
		ServiceName:   getEnv("SERVICE_NAME", "kiwify-webhook"),
		Port:          getEnv("PORT", ":3000"),
		TokenTelegram: getEnv("TOKEN_TELEGRAM", ""),
		ChatID:        getEnv("CHAT_ID", ""),
		AuthSecret:    getEnv("JWT_SECRET", "secret123"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (c *Config) CreateURIMongoDB() string {
	var uri string

	if c.Env == "development" || c.Env == "production" {
		uri = fmt.Sprintf("mongodb://%s:%s@localhost:27017", c.Key, c.Secret)
	} else {
		uri = c.MongoURL
	}

	return uri
}
