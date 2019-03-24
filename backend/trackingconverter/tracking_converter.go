package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"log"
)

const (
	awsRegion    string = "us-east-1"
	functionName string = "tracking_converter"
)

var svc *rekognition.Rekognition

type ResponseBody struct {
	Response string `json:"response"`
}

type RekSNSNotification struct {
	JobId  string `json:"JobId"`
	Status string `json:"Status"`
}

func process(notification *RekSNSNotification) error {
	var (
		maxResults      int64 = 1000
		paginationToken *string
		finished        = false
		count           = 0
	)
	for !finished {
		x := rekognition.GetPersonTrackingInput{
			JobId:      aws.String(notification.JobId),
			MaxResults: &maxResults,
			NextToken:  paginationToken,
		}
		results, err := svc.GetPersonTracking(&x)
		if err != nil {
			return err
		}

		log.Println(results.VideoMetadata)

		for _, p := range results.Persons {
			count++

			person := p.Person
			if person == nil {
				continue
			}
			log.Println("Person")

			boundingBox := person.BoundingBox
			if boundingBox == nil {
				continue
			}

			log.Println("	Bounding Box")
			log.Printf("		Top: %f", *boundingBox.Top)
			log.Printf("		Left: %f", *boundingBox.Left)
			log.Printf("		Width: %f", *boundingBox.Width)
			log.Printf("		Height: %f", *boundingBox.Height)
		}

		if results.NextToken == nil {
			finished = true
		} else {
			paginationToken = results.NextToken
		}
	}
	log.Println("Number of PersonDetection objects: ", count)

	return nil
}

func handle(ctx context.Context, snsEvent events.SNSEvent) (events.APIGatewayProxyResponse, error) {
	log.Println("Post: ", snsEvent)
	log.Println("context ", ctx)
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}

	var notification *RekSNSNotification
	jsonParseError := json.Unmarshal([]byte(snsEvent.Records[0].SNS.Message), notification)
	if jsonParseError != nil {
		log.Println(jsonParseError)
		return events.APIGatewayProxyResponse{500, headers, nil, "Internal Server Error", false}, nil
	}

	log.Println("SNS event received: ", notification)

	process(notification)

	code := 200

	// TODO change response
	response, jsonBuildError := json.Marshal(ResponseBody{Response: "TODO"})
	if jsonBuildError != nil {
		log.Println(jsonBuildError)
		response = []byte("Internal Server Error")
		code = 500
	}

	return events.APIGatewayProxyResponse{code, headers, nil, string(response), false}, nil
}

//func main() {
//	session, err := session.NewSession(&aws.Config{
//		Region: aws.String(awsRegion)},
//	)
//
//	// Create DynamoDB client
//	svc = rekognition.New(session)
//
//	if err != nil {
//		log.Println("Error initiating " + functionName + " lambda function ", err.Error())
//	} else {
//		log.Println("Successfully initiated " + functionName + " lambda function")
//		lambda.Start(handle)
//	}
//}

func main() {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Create DynamoDB client
	svc = rekognition.New(session)

	err = process(&RekSNSNotification{JobId: "51a3a9bed1dca4015708e18b24c884ecde6212fb738870500bbd440ad284e2f1"})
	if err != nil {
		log.Println(err)
	}
}
