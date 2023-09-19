// EC2 using go

package main

import (
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    var (
        instanceID string
        err error
    )
    if instanceID, err = createEC2("us-east-1"); err != nil {
        fmt.Printf("Create EC2 ERROR ")
        os.Exit(1)
    }
    fmt.Printf("InstanceId: %s\n", InstanceId)
}

func createEC2(region string) (string, error) {
    sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region),
    return "", nil
}