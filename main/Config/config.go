package Config

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName   string    `json:"app_name"`
	AppModel  string    `json:"app_model"`
	AppHost   string    `json:"app_host"`
	AppPort   string    `json:"app_port"`
	JwtConfig JwtConfig `json:"jwt_config"`
}

type JwtConfig struct {
	Issuer    string `json:"issuer"`
	Audience  string `json:"audience"`
	Expires   int64  `json:"expires"`
	SecretKey string `json:"secret_key"`
}

func GetConfig() *Config {
	return cfg
}

var cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path) //读取文件
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader) //解析json
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
