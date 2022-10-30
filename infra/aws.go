package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"showcaseme/infra/core"
	"showcaseme/internal/utils"
)

func CreateAwsSession() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(core.AppConfig.AwsRegion),
			Credentials: credentials.NewStaticCredentials(
				core.AppConfig.AwsAccessKeyId,
				core.AppConfig.AwsAccessKeyPassword,
				"",
			)})

	utils.Check(err, "error connecting to AWS service")

	return sess
}
