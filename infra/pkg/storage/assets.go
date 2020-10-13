package storage

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Assets struct {
	pulumi.ResourceState

	bucket  *storage.Bucket
	objects []*storage.BucketObject
}

func NewAssets(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*Assets, error) {
	// Create the component
	assets := &Assets{}
	err := ctx.RegisterComponentResource("pkg:storage:Assets", name, assets, opts...)
	if err != nil {
		return nil, err
	}

	// Create the bucket
	bucket, err := storage.NewBucket(ctx, fmt.Sprintf("%s-assets", name), &storage.BucketArgs{
		UniformBucketLevelAccess: pulumi.Bool(true),
	}, pulumi.Parent(assets))
	if err != nil {
		return nil, err
	}
	_, err = storage.NewBucketIAMMember(ctx, fmt.Sprintf("%s-assets-iam-public", name), &storage.BucketIAMMemberArgs{
		Bucket: bucket.Name,
		Role:   pulumi.String("roles/storage.legacyObjectReader"),
		Member: pulumi.String("allUsers"),
	}, pulumi.Parent(assets))
	if err != nil {
		return nil, err
	}
	assets.bucket = bucket

	// Upload images
	image, err := storage.NewBucketObject(ctx, fmt.Sprintf("%s-assets-logo", name), &storage.BucketObjectArgs{
		Bucket:      bucket.Name,
		Name:        pulumi.String("logo.png"),
		ContentType: pulumi.String("image/png"),
		Source:      pulumi.NewFileAsset("./assets/logo.png"),
	})
	if err != nil {
		return nil, err
	}
	assets.objects = append(assets.objects, image)

	return assets, nil
}
