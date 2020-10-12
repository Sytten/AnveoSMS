package config

import "github.com/spf13/viper"

type Configuration struct {
	Email EmailConfiguration
}

type EmailConfiguration struct {
	SendgridApiKey string
	From           string
	To             string
}

func NewConfiguration() (*Configuration, error) {
	v := viper.New()

	// Set the config file details
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(".")

	// Read the configuration
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Configuration
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
