package gcp

import (
	"context"
	"fmt"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func ListInstances() {
	ctx := context.Background()

	computeService, err := compute.NewService(ctx, option.WithCredentialsFile("path/to/your/credentials.json"))
	if err != nil {
		fmt.Println("Error creating compute service:", err)
		return
	}

	projectID := "your-gcp-project-id"
	zone := "us-central1-a"

	req := computeService.Instances.List(projectID, zone)
	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			fmt.Printf("Instance ID: %s\n", instance.Id)
		}
		return nil
	}); err != nil {
		fmt.Println("Error listing instances:", err)
	}
}
