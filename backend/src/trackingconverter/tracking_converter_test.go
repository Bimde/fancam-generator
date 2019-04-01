package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	log "github.com/sirupsen/logrus"
)

func TestProcess(t *testing.T) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	svc = rekognition.New(session)

	err = process(&rekSNSNotification{JobID: "9c12697cbba2cc8bfc2051836b60e5c34be3bf5313127f2736eaf9b159c90c51"})
	// err = process(&rekSNSNotification{JobID: "51a3a9bed1dca4015708e18b24c884ecde6212fb738870500bbd440ad284e2f1"})
	if err != nil {
		t.Error(err)
	}

	t.Logf("Project ID: %d", project.ID)
}
