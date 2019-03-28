package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	awsRegion            = "us-east-1"
	bucket               = "fancamgenerator"
	preSignExpiryMinutes = 15
	loggingName          = "lambda"
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
	log := getLogger("handle")
	log.Info("context ", ctx)
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
	}

	var body RequestBody
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           headers,
			MultiValueHeaders: nil,
			Body:              "Internal Server Error",
			IsBase64Encoded:   false,
		}, nil
	}

	videoID := getKey(body.UserId)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(videoID),
		Body:   strings.NewReader("EXPECTED CONTENTS"),
	})

	url, err := req.Presign(preSignExpiryMinutes * time.Minute)

	if err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           headers,
			MultiValueHeaders: nil,
			Body:              "Internal Server Error",
			IsBase64Encoded:   false,
		}, nil
	}

	log.Error("The URL is: ", url)

	code := 200
	response, err := json.Marshal(ResponseBody{PreSignedURL: url, VideoId: videoID})
	if err != nil {
		log.Error(err)
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

	log := getLogger("main")

	if err != nil {
		log.Panic("Error initiating session ", err)
	} else {
		log.Info("Successfully initiated session")
		lambda.Start(handle)
	}
}

func getLogger(method string) *log.Entry {
	return log.WithFields(log.Fields{
		"method": fmt.Sprintf("%s#%s", loggingName, method),
	})
}
