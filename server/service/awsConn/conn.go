package awsconn

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/xadhithiyan/videon/config"
)

var SVC *s3.S3 = Conn()

func Conn() *s3.S3 {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Env.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			config.Env.AWSAccessKey,
			config.Env.AWSSecrectKey,
			"",
		),
	})
	if err != nil {
		log.Fatal("AWS connection error")
	}

	return s3.New(sess)
}
