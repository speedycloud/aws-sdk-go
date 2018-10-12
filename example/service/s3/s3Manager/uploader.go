package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	ACCESS_KEY              = "FIX_ME"
	SECRET_KEY              = "FIX_ME"
	BUCKET                  = "FIX_ME"
	KEY                     = "FIX_ME"
	SPEEDYCLOUD_OS_ENDPOINT = "https://oss-cn-shanghai.speedycloud.org"
)

func initS3Client() *s3manager.Uploader {
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

	// 实例化 uploader 对象
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 16 * 1024 * 1024 // 16MB per part
	})

	return uploader
}

func main() {
	// 使用本地文件
	// dd if=/dev/zero of=/tmp/foo.file bs=1m count=50
	file, err := os.Open("/tmp/foo.file")
	if err != nil {
		fmt.Print("Error = ", err.Error())
		return
	}

	// 初始化
	uploader := initS3Client()

	// 初始化上传参数
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(BUCKET),
		Key:    aws.String(KEY),
		Body:   file,
	}

	// 启动上传，支持可配选项，如分片大小
	result, err := uploader.Upload(upParams, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 10MB part size
		u.LeavePartsOnError = true    // Don't delete the parts if the upload fails.
	})

	if err != nil {
		fmt.Print("上传出错，Error = ", err.Error())
		return
	}
	fmt.Print("上传完成, Upload ID = ", result.UploadID)

	return
}
