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
         S3_BUCKET = ""
         S3_ACL    = ""
      )

type S3Handler struct {
        Session *session.Session
        Bucket string
}


func main() {
        
        if len(os.Args) != 3 {
                log.FatalF("usage: %s <filename> <s3-filepath>\n",
        }
        filename := os.Args[1]
        key      := os.Args[2]
                           
                           file, err := os.Open(filename)
                           if err != nil {
                                   log.Fatalf("os.Open - filename %v, err: %v", filename, err)
                           }
                           defer file.Close()
                           
                           
                           sess, err :=session.NewSession(&aws.Config{Region: aws.String(S3_REGION))}
                           if err != nil {
                                   log.FatalF("session.NewSession - filename: %v, err: %v", filename, err)
                           }
                                                          
                           handler := S3Handler {
                               v   Session: sess,
                                   Bucket: S3_BUCKET,
                           }
                           
                           err = handler.UploadFile(key, filename)
                           if err != nil {       
                                   log.FatalF("UploadFile - filename: %v, err: %v", filaname, err)
                           }
                           log.Info("UploadFile - success")
                           
}
                                                          
func (h S3Handler) UploadFile(key string, filename string) error {
        file, err := os.Open(filename)
        if err != nil {
                log.FatalF("os.Open - filename: %s, err: %v", filename, err)
        defer file.Close()
                
        _, err = s3.New(h.session).PutObject(&s3.PutObjectInput{
                Bucket:             aws.String(h.Bucket),
                Key:                aws.String(key),
                ACL:                aws.String(S3_ACL)
                Body:               file,
                ContentDisposition: aws.String("attachment"),
        })
        return err

func (h S3Handler) ReadFile(key string) (string, error
        results, err := s3.New(h.Session).GetObject(&s3.GetObjectInput{
                Bucket: aws.String(h.bucket),
                Key:    aws.String(key),
        })
        if err != nil {
                return "", err
        }
        defer results.Body.Close()   
        buf := bytes.NewBuffer(nil)
        if _, err := io.Copy(buf, results.Body); err != nil {
                return "", err
        }
        return string(buf.Bytes()), nil
 }
                                         
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                                          
                                            
                                   
            
                                                           
                                                           
                                                           
                                                          
         
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
                                
