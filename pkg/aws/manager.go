package aws

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func ListInstances() {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"),
    }))

    svc := ec2.New(sess)

    result, err := svc.DescribeInstances(nil)
    if err != nil {
        fmt.Println("Error", err)
        return
    }

    for _, reservation := range result.Reservations {
        for _, instance := range reservation.Instances {
            fmt.Printf("Instance ID: %s\n", *instance.InstanceId)
        }
    }
}
