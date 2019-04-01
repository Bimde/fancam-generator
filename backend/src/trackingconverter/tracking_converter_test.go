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

	err = process(&rekSNSNotification{JobID: "1e9e70a679024b396e5b8145ba3b6f69a17a6e82810b268793d8df256994c4a4"})
	if err != nil {
		t.Error(err)
	}
}
