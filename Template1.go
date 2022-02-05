package main

import (
        "context"  //
        "flag" //
        "fmt"
        "log"
        
        "github.com/aws/aws-sdk-go-v2/config" // this package loads configuration
        "github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
        bucketName      string 
        objectPrefix    string
        objectDelimiter string
        maxKeys         int
)
        
