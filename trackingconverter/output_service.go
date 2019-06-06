package main

import (
	"fmt"
	"strings"

	"github.com/Bimde/fancam-generator/openshot"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	outputFormat = "[Length: %d seconds] ExportURL: %s\n"
	defaultTopicARN = "arn:aws:sns:us-east-1:744292932026:FancamGeneratorExportOutput"
)

// NotifyExportCompletion publishes the URLs of each given export to the specified SNS Topic
func NotifyExportCompletion(exports *[]*openshot.Export, topicARN *string) {
	if topicARN == nil {
		topicARN = aws.String(defaultTopicARN)
	}
	topic.Publish(&sns.PublishInput{
		Message:  aws.String(getMessage(exports)),
		TopicArn: topicARN,
	})
}

func getMessage(exports *[]*openshot.Export) string {
	var sb strings.Builder
	for _, export := range *exports {
		sb.WriteString(toString(export))
	}
	return sb.String()
}

func toString(export *openshot.Export) string {
	durationInSeconds := (float64(export.EndFrame) - float64(export.StartFrame)) / fps
	return fmt.Sprintf(outputFormat, int(durationInSeconds), export.URL)
}


