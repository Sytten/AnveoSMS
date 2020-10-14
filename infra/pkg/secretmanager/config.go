package secretmanager

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/secretmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
	"github.com/sytten/anveosms/infra/pkg/storage"
	"gopkg.in/yaml.v2"
)

type Config struct {
	pulumi.ResourceState

	secret *secretmanager.Secret
}

type ConfigArgs struct {
	Assets *storage.Assets
}

func NewConfig(ctx *pulumi.Context, name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	// Create component
	c := &Config{}
	err := ctx.RegisterComponentResource("pkg:secretmanager:Config", name, c, opts...)
	if err != nil {
		return nil, err
	}

	// Create the secret
	secret, err := secretmanager.NewSecret(ctx,
		fmt.Sprintf("%s-config", name),
		&secretmanager.SecretArgs{
			Replication: &secretmanager.SecretReplicationArgs{
				Automatic: pulumi.Bool(true),
			},
			SecretId: pulumi.Sprintf("%s-config", name),
		}, pulumi.Parent(c))
	if err != nil {
		return nil, err
	}
	c.secret = secret

	// Store the secret data
	secretData, err := buildYamlConfiguration(ctx, args)
	if err != nil {
		return nil, err
	}
	_, err = secretmanager.NewSecretVersion(ctx, fmt.Sprintf("%s-config-version", name), &secretmanager.SecretVersionArgs{
		Secret:     secret.Name,
		SecretData: secretData,
	}, pulumi.Parent(c))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func buildYamlConfiguration(ctx *pulumi.Context, args *ConfigArgs) (pulumi.StringOutput, error) {
	// Load the data from the pulumi config
	c := config.New(ctx, "email")
	m := map[string]map[string]string{
		"email": {
			"sendgridApiKey": c.Require("sendgridApiKey"),
			"from":           c.Require("from"),
			"to":             c.Require("to"),
		},
		"hosting": {
			"bucketUrl": "%s",
		},
	}

	// Create the yaml
	data, err := yaml.Marshal(m)
	if err != nil {
		return pulumi.StringOutput{}, err
	}

	// Interpolate the string with the pulumi values
	// Note: this is a work around without any obvious path for fixing
	// until pulumi allows for output to be waited for.
	return pulumi.Sprintf(string(data[:]), args.Assets.GetUrl()), nil
}
