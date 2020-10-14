package config

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
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
	projectId, err := getProjectId(ctx)
	if err != nil {
		return nil, err
	}
	secret, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectId, secretName),
	})
	if err != nil {
		return nil, err
	}

	// Load config
	v := viper.New()
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

func getProjectId(ctx context.Context) (string, error) {
	// Fetch from the credentials
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return "", err
	}
	if credentials.ProjectID != "" {
		return credentials.ProjectID, nil
	}

	// Fetch from the environment variable
	projectId := os.Getenv("APP_PROJECT_ID")
	if projectId != "" {
		return projectId, nil
	}

	return "", errors.New("unable to retrieve project id")
}
