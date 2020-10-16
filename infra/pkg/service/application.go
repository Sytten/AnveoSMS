package service

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/serviceaccount"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"

	"github.com/sytten/anveosms/infra/pkg/secretmanager"
)

type Application struct {
	pulumi.ResourceState

	serviceAccount *serviceaccount.Account
	service        *cloudrun.Service
}

type ApplicationArgs struct {
	Config *secretmanager.Config
	Image  *Image
}

func NewApplication(ctx *pulumi.Context, name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	// Create component
	app := &Application{}
	err := ctx.RegisterComponentResource("pkg:service:App", name, app, opts...)
	if err != nil {
		return nil, err
	}

	// Create service account
	serviceAccount, err := serviceaccount.NewAccount(ctx, fmt.Sprintf("%s-api-iam", name), &serviceaccount.AccountArgs{
		AccountId: pulumi.Sprintf("%s-api", name),
	}, pulumi.Parent(app))
	if err != nil {
		return nil, err
	}
	app.serviceAccount = serviceAccount

	err = secretmanager.NewConfigAccessor(ctx, "api", &secretmanager.ConfigAccessorArgs{
		Config: args.Config,
		Member: pulumi.Sprintf("serviceAccount:%s", serviceAccount.Email),
	})
	if err != nil {
		return nil, err
	}

	// Create service
	service, err := cloudrun.NewService(ctx, fmt.Sprintf("%s-api", name), &cloudrun.ServiceArgs{
		Location:                 pulumi.String(config.Require(ctx, "gcp:region")),
		AutogenerateRevisionName: pulumi.Bool(true),
		Traffics: cloudrun.ServiceTrafficArray{
			cloudrun.ServiceTrafficArgs{
				LatestRevision: pulumi.Bool(true),
				Percent:        pulumi.Int(100),
			},
		},
		Template: cloudrun.ServiceTemplateArgs{
			Metadata: cloudrun.ServiceTemplateMetadataArgs{
				Annotations: pulumi.StringMap{"autoscaling.knative.dev/maxScale": pulumi.String("5")},
			},
			Spec: cloudrun.ServiceTemplateSpecArgs{
				ContainerConcurrency: pulumi.Int(80),
				Containers: cloudrun.ServiceTemplateSpecContainerArray{
					cloudrun.ServiceTemplateSpecContainerArgs{
						Image: args.Image.GetName(),
						Ports: cloudrun.ServiceTemplateSpecContainerPortArray{
							cloudrun.ServiceTemplateSpecContainerPortArgs{
								ContainerPort: pulumi.Int(9000),
							},
						},
						Resources: cloudrun.ServiceTemplateSpecContainerResourcesArgs{
							Limits: pulumi.StringMap{"cpu": pulumi.String("1000m"), "memory": pulumi.String("1024Mi")},
						},
						Envs: cloudrun.ServiceTemplateSpecContainerEnvArray{
							cloudrun.ServiceTemplateSpecContainerEnvArgs{
								Name:  pulumi.String("APP_SECRET_NAME"),
								Value: args.Config.GetSecretName(),
							},
						},
					},
				},
				ServiceAccountName: serviceAccount.Email,
			},
		},
	}, pulumi.Parent(app))
	if err != nil {
		return nil, err
	}
	app.service = service

	// Create service access
	_, err = cloudrun.NewIamMember(ctx, fmt.Sprintf("%s-api-access"), &cloudrun.IamMemberArgs{
		Service:  service.Name,
		Location: pulumi.String(config.Require(ctx, "gcp:region")),
		Role:     pulumi.String("roles/run.invoker"),
		Member:   pulumi.String("allUsers"),
	}, pulumi.Parent(app))
	if err != nil {
		return nil, err
	}

	return app, nil
}
