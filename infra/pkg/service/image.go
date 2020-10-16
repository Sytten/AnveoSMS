package service

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/config"
	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"

	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"

	"github.com/sytten/anveosms/infra/pkg/git"
)

type Image struct {
	pulumi.ResourceState

	registry *container.Registry
	image    *docker.Image
}

func NewImage(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*Image, error) {
	// Create component
	image := &Image{}
	err := ctx.RegisterComponentResource("pkg:service:Image", name, image, opts...)
	if err != nil {
		return nil, err
	}

	// Create registry
	registry, err := container.NewRegistry(ctx, fmt.Sprintf("%s-registry", name), &container.RegistryArgs{
		Location: pulumi.String("US"),
	}, pulumi.Parent(image))
	image.registry = registry

	// Find related tag
	version, err := git.Describe()
	if err != nil {
		version = "unknown"
	}

	// Get client credentials
	clientConfig, err := organizations.GetClientConfig(ctx)
	if err != nil {
		return nil, err
	}

	// Build & Push Image
	img, err := docker.NewImage(ctx, fmt.Sprintf("%s-image", name), &docker.ImageArgs{
		ImageName: pulumi.Sprintf("gcr.io/%s/%s:%s", config.GetProject(ctx), name, version),
		Build: docker.DockerBuildArgs{
			Context: pulumi.String("../app"),
			Args:    pulumi.StringMap{"BUILD_VERSION": pulumi.String(version)},
		},
		Registry: docker.ImageRegistryArgs{
			Server:   pulumi.String("gcr.io"),
			Username: pulumi.String("oauth2accesstoken"),
			Password: pulumi.String(clientConfig.AccessToken),
		},
	}, pulumi.Parent(image))
	if err != nil {
		return nil, err
	}
	image.image = img

	return image, nil
}

func (i *Image) GetName() pulumi.StringOutput {
	return i.image.ImageName
}
