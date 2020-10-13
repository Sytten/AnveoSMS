package storage

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Assets struct {
	pulumi.ResourceState

	bucket *storage.Bucket
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

	return assets, nil
}
