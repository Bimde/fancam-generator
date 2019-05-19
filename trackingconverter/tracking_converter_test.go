package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	log "github.com/sirupsen/logrus"
)

// Rek job IDs -- these are uppercase to match the file names
const (
	DALLA_DALLA = "23ad7dce2baa000b3a29c1226d08e3eeca5338476e3ac95a149ddf25767abf1f"
	IDOL        = "_"
	LATATA      = "_"
	BOSS        = "acda4cfe4311f4dd0b18b4f1cb81109cf74e5b8a5676996f1ca8b36e6a9ecf26"
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

	err = process(&rekSNSNotification{JobID: DALLA_DALLA})
	if err != nil {
		t.Error(err)
	}
}