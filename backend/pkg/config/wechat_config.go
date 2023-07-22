package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v7"
	"github.com/gofiber/fiber/v2/log"
)

var (
	cfg *Config
)

// Config example config
type Config struct {
	Redis struct {
		Host        string `env:"REDIS_HOST"`
		Port        int    `env:"REDIS_PORT"`
		Password    string `env:"REDIS_PASSWORD"`
		Database    int    `env:"REDIS_DB_NUMBER"`
		MaxActive   int    `env:"maxActive"`
		MaxIdle     int    `env:"maxIdle"`
		IdleTimeout int    `env:"idleTimeout"`
	} `env:"redis"`
	OfficialAccountConfig `env:"wechat"`
}

// OfficialAccountConfig 公众号相关配置
type OfficialAccountConfig struct {
	AppID          string `env:"WECHAT_APPID"`
	AppSecret      string `env:"WECHAT_APPSECRET"`
	Token          string `env:"WECHAT_TOKEN"`
	EncodingAESKey string `env:"WECHAT_ENCODINGAESKEY"`
}

// GetConfig 获取配置
func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	log.Infof("we is %v", os.Getenv("WECHAT_APPID"))

	log.Infof("cfg is %v", cfg)
	return cfg
}
