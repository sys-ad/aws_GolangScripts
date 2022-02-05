package main

import (
  ""
  ""
  "fmt"
  "github.com/aws/aws-sdk-go"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/session"
)

func main() {
  svc := ec2.New(session.New())
  input := &ec2.RunInstancesInput{
    ImageId:      aws.String(AMI),
    InstanceType: aws.String(INSTANCE_TYPE),
    KeyName:      aws.String(KEY_NAME),
    MaxCount:     aws.Int64(1),
    MinCount:     aws.Int64(1),
 }          
  
}

func Lambda_Handler(ctx context, name event) (string, error) {
     return 
     
}


