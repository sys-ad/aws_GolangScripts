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
  input := &ec2.RunInstancesInput{
    ImageId:      os.Getenv("AMI_ID"),
    InstanceType: os.Getenv("INSTANCE_TYPE"),
    KeyName:      os.Getenv("KEY_NAME"),
    MaxCount:     os.Getenv("MaxCount"),
    MinCount:     os.Getenv("MinCount"),
 }          
  
}

func Lambda_Handler(ctx context, name event) (string, error) {
     return 
     
}


