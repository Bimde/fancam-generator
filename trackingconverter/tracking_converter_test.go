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
	DALLA_DALLA = "014b59c95c350f695c50531e44f73c564e6d281261b71eb844ec460d867b042b"
	IDOL        = "1e9e70a679024b396e5b8145ba3b6f69a17a6e82810b268793d8df256994c4a4"
	LATATA      = "b29b5fb116940d3328144d2f535067743d2a820fb8aa8d914977e4fed5993ab6"
	BOSS        = "f3cfae338f083bc01d1e074ceaa09665af8417e5d7eaa181b68372c96bcb5790"
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
