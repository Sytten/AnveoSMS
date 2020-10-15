package secretmanager

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/secretmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ConfigAccessorArgs struct {
	Config *Config
	Member pulumi.StringInput
}

func NewConfigAccessor(ctx *pulumi.Context, name string, args *ConfigAccessorArgs) error {
	_, err := secretmanager.NewSecretIamMember(ctx, fmt.Sprintf("%s-config-iam-%s", args.Config.name, name), &secretmanager.SecretIamMemberArgs{
		Role:     pulumi.String("roles/secretmanager.secretAccessor"),
		SecretId: args.Config.secret.SecretId,
		Member:   args.Member,
	}, pulumi.Parent(args.Config))
	return err
}
