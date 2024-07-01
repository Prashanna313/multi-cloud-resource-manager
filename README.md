# Multi-Cloud Resource Manager CLI

A CLI tool to manage resources across AWS and GCP.

## Project Structure

```plaintext
multi-cloud-resource-manager/
├── cmd/
│   ├── aws.go
│   ├── gcp.go
│   ├── root.go
├── pkg/
│   ├── aws/
│   │   └── manager.go
│   ├── gcp/
│   │   └── manager.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Usage

1. **Build the CLI:**

   ```sh
   go build -o multi-cloud-resource-manager
   ```

2. **Run AWS Commands:**

   ```sh
   ./multi-cloud-resource-manager aws list-instances
   ```

3. **Run GCP Commands:**

   ```sh
   ./multi-cloud-resource-manager gcp list-instances
   ```

## Dependencies

- [AWS SDK for Go](https://github.com/aws/aws-sdk-go)
- [Google Cloud SDK for Go](https://pkg.go.dev/google.golang.org/api/compute/v1)
- [Cobra CLI](https://github.com/spf13/cobra)
```

This `README.md` provides a basic overview of the project structure, usage instructions, and dependencies for your Multi-Cloud Resource Manager CLI.