package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	configuration Config
	once          sync.Once
)

const (
	EnvGoEnvironment = "GO_ENVIRONMENT"

	filePathFormat = "%s/config/yml/config.yml"
)

type Config struct {
	RestaurantCache   RestaurantCacheConfig   `yaml:"restaurant_cache"`
	WebSourceMetadata WebSourceMetadataConfig `yaml:"web_source_metadata"`
}

type RestaurantCacheConfig struct {
	CacheTTLSeconds int `yaml:"cache_ttl_seconds"`
}

type WebSourceMetadataConfig struct {
	RefreshPeriodSeconds int `yaml:"refresh_period_seconds"`
}

func GetConfig() Config {
	once.Do(func() {
		filePath := getYamlFilePath()

		configFile, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("yamlFile.Get err: %v", err)
		}

		configuration = Config{}
		if err := yaml.Unmarshal(configFile, &configuration); err != nil {
			log.Fatalf("yamlFile.Unmarshall err: %v", err)
		}
	})

	return configuration
}

func getYamlFilePath() string {
	env := os.Getenv(EnvGoEnvironment)

	var basePath string

	switch env {
	case "production":
		basePath, _ = os.Getwd()
	default:
		_, file, _, _ := runtime.Caller(1)
		basePath = strings.TrimSuffix(file, "/config/config.go")
	}

	filePath := fmt.Sprintf(filePathFormat, basePath)

	return filePath
}
