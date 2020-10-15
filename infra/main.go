package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"

	"github.com/sytten/anveosms/infra/pkg/secretmanager"
	"github.com/sytten/anveosms/infra/pkg/service"
	"github.com/sytten/anveosms/infra/pkg/storage"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Storage
		assets, err := storage.NewAssets(ctx, "public")
		if err != nil {
			return err
		}

		// Secret Manager
		config, err := secretmanager.NewConfig(ctx, "app", &secretmanager.ConfigArgs{
			Assets: assets,
		})

		// Service
		_, err = service.NewApplication(ctx, "anveosms", &service.ApplicationArgs{
			Config: config,
		})

		return nil
	})
}
