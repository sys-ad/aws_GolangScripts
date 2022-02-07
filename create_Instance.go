package main

import (
  ""
  ""
  "fmt"
  "github.com/aws/aws-sdk-go"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/session"
  "github.com/os"
)

func main() {
  svc := ec2.New(session.New())
  lambda.Start(Lambda_Handler)
}

func Lambda_Handler(ctx context, name event) (string, error) {
    message := event['message']
    init_script := `#/bin/bash
                    sudo yum update -y`
    Input :=      &ec2.RunInstancesInput{
    ImageId:      os.Getenv("AMI_ID"),
    InstanceType: os.Getenv("INSTANCE_TYPE"),
    KeyName:      os.Getenv("KEY_NAME"),
    MaxCount:     os.Getenv("MaxCount"),
    MinCount:     os.Getenv("MinCount"),
    
    instance_id := input['Instances'][0]['InstanceId']
    fmt.printLn(instance_id)         
    return instance_id 
 }          
    
}


