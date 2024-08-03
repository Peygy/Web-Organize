package config

import (
	"os"

	"github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Gin *gin.GinConfig `yaml:"gin"`
	Grpc *grpc.GrpcConfig `yaml:"grpc"`
}

func InitConfig(filePath string) (*Config, *gin.GinConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, nil, err
	}

	return cfg, cfg.Gin, nil
}
