#Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#PDX-License-Identifier: MIT-0 (For details, see https://github.com/awsdocs/amazon-rekognition-developer-guide/blob/master/LICENSE-SAMPLECODE.)

import boto3
import json
import sys


class VideoDetect:
    jobId = 'a2ad173ff2d498e84eea9810078597801b457ad8eb57e964a3c3690224e9900a'
    rek = boto3.client('rekognition', 'us-east-1')
    queueUrl = 'https://sqs.us-east-1.amazonaws.com/744292932026/FullGroupVideos'
    roleArn = 'arn:aws:iam::744292932026:role/RekognitionServiceRole'
    topicArn = 'arn:aws:sns:us-east-1:744292932026:AmazonRekognitionFullGroupVideos'
    bucket = 'fancamgenerator'
    video = 'ITZY.mp4'

    def main(self):
        # triggerNewJob()
        self.GetResultsPersons(self.jobId)

    def triggerNewJob(self):
        jobFound = False
        sqs = boto3.client('sqs', 'us-east-1')
       

        #=====================================
        response = self.rek.start_person_tracking(Video={'S3Object':{'Bucket':self.bucket,'Name':self.video}}, 
            NotificationChannel={'RoleArn':self.roleArn, 'SNSTopicArn':self.topicArn})
        #=====================================

        print(response)
        print('Start Job Id: ' + response['JobId'])
        dotLine=0
        while jobFound == False:
            sqsResponse = sqs.receive_message(QueueUrl=self.queueUrl, MessageAttributeNames=['ALL'],
                                          MaxNumberOfMessages=10)

            if sqsResponse:
                if 'Messages' not in sqsResponse:
                    if dotLine<20:
                        print('.', end='')
                        dotLine=dotLine+1
                    else:
                        print()
                        dotLine=0    
                    sys.stdout.flush()
                    continue

                print(sqsResponse)
                for message in sqsResponse['Messages']:
                    notification = json.loads(message['Body'])
                    rekMessage = notification
                    # rekMessage = json.loads(notification['Message'])
                    print(rekMessage['JobId'])
                    print(rekMessage['Status'])
                    if str(rekMessage['JobId']) == response['JobId']:
                        print('Matching Job Found:' + rekMessage['JobId'])
                        jobFound = True
                        #=============================================
                        self.GetResultsPersons(rekMessage['JobId'])
                        #=============================================

                        sqs.delete_message(QueueUrl=self.queueUrl,
                                       ReceiptHandle=message['ReceiptHandle'])
                    else:
                        print("Job didn't match:" +
                              str(rekMessage['JobId']) + ' : ' + str(response['JobId']))
                    # Delete the unknown message. Consider sending to dead letter queue
                    sqs.delete_message(QueueUrl=self.queueUrl,
                                   ReceiptHandle=message['ReceiptHandle'])

        print('done')


    def GetResultsLabels(self, jobId):
        maxResults = 10
        paginationToken = ''
        finished = False

        while finished == False:
            response = self.rek.get_label_detection(JobId=jobId,
                                            MaxResults=maxResults,
                                            NextToken=paginationToken,
                                            SortBy='TIMESTAMP')

            print(response['VideoMetadata']['Codec'])
            print(str(response['VideoMetadata']['DurationMillis']))
            print(response['VideoMetadata']['Format'])
            print(response['VideoMetadata']['FrameRate'])

            for labelDetection in response['Labels']:
                label=labelDetection['Label']

                print("Timestamp: " + str(labelDetection['Timestamp']))
                print("   Label: " + label['Name'])
                print("   Confidence: " +  str(label['Confidence']))
                if 'Instances' in label:
                    print("   Instances:")
                    for instance in label['Instances']:
                        print ("      Confidence: " + str(instance['Confidence']))
                        print ("      Bounding box")
                        print ("        Top: " + str(instance['BoundingBox']['Top']))
                        print ("        Left: " + str(instance['BoundingBox']['Left']))
                        print ("        Width: " +  str(instance['BoundingBox']['Width']))
                        print ("        Height: " +  str(instance['BoundingBox']['Height']))
                        print()
                print()
                print ("   Parents:")
                for parent in label['Parents']:
                    print ("      " + parent['Name'])
                print ()

                if 'NextToken' in response:
                    paginationToken = response['NextToken']
                else:
                    finished = True

    def GetResultsPersons(self, jobId):
        maxResults = 10
        paginationToken = ''
        finished = False

        while finished == False:
            response = self.rek.get_person_tracking(JobId=jobId,
                                            MaxResults=maxResults,
                                            NextToken=paginationToken)

            # print(response)
            print(response['VideoMetadata']['Codec'])
            print(str(response['VideoMetadata']['DurationMillis']))
            print(response['VideoMetadata']['Format'])
            print(response['VideoMetadata']['FrameRate'])

            for personDetection in response['Persons']:
                person = personDetection['Person']
                print('Index: ' + str(person['Index']))
                print('Timestamp: ' + str(personDetection['Timestamp']))
                if 'BoundingBox' in person:
                    print ("      Bounding box")
                    print ("        Top: " + str(person['BoundingBox']['Top']))
                    print ("        Left: " + str(person['BoundingBox']['Left']))
                    print ("        Width: " +  str(person['BoundingBox']['Width']))
                    print ("        Height: " +  str(person['BoundingBox']['Height']))
                print()

            if 'NextToken' in response:
                paginationToken = response['NextToken']
            else:
                finished = True
if __name__ == "__main__":

    analyzer=VideoDetect()
    analyzer.main()
