package config

import (
	"bytes"
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/spf13/viper"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func loadFromSecret(secretName string) (*Configuration, error) {
	// Create client
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	// Download secret
	secret, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("%s/versions/latest", secretName),
	})
	if err != nil {
		return nil, err
	}

	// Load config
	v := viper.New()
	v.SetConfigType("yml")
	err = v.ReadConfig(bytes.NewBuffer(secret.Payload.Data))
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
