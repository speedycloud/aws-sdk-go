// +build example

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	ACCESS_KEY                     = "FIX_ME"
	SECRET_KEY                     = "FIX_ME"
	SPEEDYCLOUD_OS_ENDPOINT string = "https://oss-cn-shanghai.speedycloud.org"
)

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must specify a bucket")
		return
	}

	// 构建配置项
	cfg := &aws.Config{
		Region:                        aws.String("us-east-1"), //By default, Don't edit it
		Credentials:                   credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
		Endpoint:                      aws.String(SPEEDYCLOUD_OS_ENDPOINT),
		S3DisableContentMD5Validation: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
	}

	// 实例化 Session 对象
	sess := session.Must(session.NewSession(cfg))

	// 实例化 Service Client
	svc := s3.New(sess)

	i := 0
	// 调用 ListObjectsPages 接口查询对象列表
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &os.Args[1],
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
}
