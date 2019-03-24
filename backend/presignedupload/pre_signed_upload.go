package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"strings"
	"time"
)

const (
	awsRegion            = "us-east-1"
	bucket               = "fancamgenerator"
	preSignExpiryMinutes = 15
)

var (
	svc *s3.S3
)

type RequestBody struct {
	UserId string `json:"userId"`
}

type ResponseBody struct {
	PreSignedURL string `json:"preSignedURL"`
	VideoId      string `json:"videoId"`
}

func handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("context ", ctx)
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
	}

	var body RequestBody
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           headers,
			MultiValueHeaders: nil,
			Body:              "Internal Server Error",
			IsBase64Encoded:   false,
		}, nil
	}

	videoId := getKey(body.UserId)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(videoId),
		Body:   strings.NewReader("EXPECTED CONTENTS"),
	})

	url, err := req.Presign(preSignExpiryMinutes * time.Minute)

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           headers,
			MultiValueHeaders: nil,
			Body:              "Internal Server Error",
			IsBase64Encoded:   false,
		}, nil
	}

	log.Println("The URL is: ", url)

	code := 200
	response, err := json.Marshal(ResponseBody{PreSignedURL: url, VideoId: videoId})
	if err != nil {
		log.Println(err)
		response = []byte("Internal Server Error")
		code = 500
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        code,
		Headers:           headers,
		MultiValueHeaders: nil,
		Body:              string(response),
		IsBase64Encoded:   false,
	}, nil
}

func getKey(id string) string {
	return fmt.Sprint(id, '-', time.Now().Unix())
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)

	svc = s3.New(sess)

	if err != nil {
		log.Println("Error initiating pre_signed_upload lambda function ", err.Error())
	} else {
		log.Println("Successfully initiated pre_signed_upload lambda function")
		lambda.Start(handle)
	}
}
