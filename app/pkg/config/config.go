package config

import (
	"os"

	"go.uber.org/zap"
)

type Configuration struct {
	Email   EmailConfiguration
	Hosting HostingConfiguration
}

type EmailConfiguration struct {
	SendgridApiKey string
	From           string
	To             string
}

type HostingConfiguration struct {
	BucketUrl string
}

func NewConfiguration(logger *zap.Logger) (*Configuration, error) {
	// Load from file
	config, err := loadFromFile()
	if config != nil {
		return config, nil
	}

	// Load from secret
	secretName := os.Getenv("APP_SECRET_NAME")
	if secretName != "" {
		config, err := loadFromSecret(secretName)
		if err != nil {
			logger.Warn("Unable to load secret", zap.Error(err))
		}
		if config != nil {
			return config, nil
		}
	}

	return nil, err
}
