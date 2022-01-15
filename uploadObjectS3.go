package main

import (
        "os"
        "io"
        "path/filepath"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
        "fmt"
        
)

const (
         S3_REGION = "us-east-1"
         S3_BUCKET = "music-of-pomm"
         S3_ACL    = "public-read"
      )

type S3Handler struct {
        Session *session.Session
        Bucket string
}


func main() {
        sess, err :=session.NewSessionWithOptions(session.Options{
         
        if err != nil {
                fmt.Printf("Failed to initialize new session: %v", err)
                return
                }
        }        
                uploader := s3manager.NewUploader(sess)
                f, err := os.Open(somepdf.pdf) // stream mode??
                if err != nil {
                    return fmt.Errorf("failed to open file %q
                
                result, err := uploader.Upload(&s3manager.UploadInput{
                    Bucket: aws.String(myBucket),
                    Key:    aws.String(myString),
                    Body:   f,
        })
                                      if err != nil {
                                              return.Errorf("failed to upload file, %v", err)
                                      }
                                      fmt.Printf("file uploaded to, %s\n", aws.Stringvalue(result.Location))
                       // to be continued... 12/13/21'
// -----------------------------------------------------------------------------------------
        
                                      
  func uploadDataToS3(dir string, svc *s3.S3) {
      fileList := []string{}
      filepath.          
                                
