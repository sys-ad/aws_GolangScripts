package main

import (
        "os"
        
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ec2"
  
         "fmt"
)

func main() {
        sess, err :=session.NewSessionWithOptions(session.Options{
                Profile: "default"
                Config: aws.Config{
                        Region: aws.String{"us-east-1"),
                },
        })
         
        if err != nil {
                fmt.Printf("Failed to initialize new session: %v", err)
                return
        }
                
                uploader := s3manager.NewUploader(sess)
                
                result, err := uploader.Upload(&s3manager.UploadInput{
                    Bucket: aws.String(myBucket),
                    Key:    aws.String(myString),
                    Body:    
                
                                
