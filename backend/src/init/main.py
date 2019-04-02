#Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#PDX-License-Identifier: MIT-0 (For details, see https://github.com/awsdocs/amazon-rekognition-developer-guide/blob/master/LICENSE-SAMPLECODE.)

import boto3
import json
import sys

class BoundingBox:
    def __init__(self, top, left, width, height, timestamp):
        self.top = top
        self.left = left
        self.width = width
        self.height = height
        self.timestamp = timestamp

class VideoDetect:
    jobId = '51a3a9bed1dca4015708e18b24c884ecde6212fb738870500bbd440ad284e2f1'
    rek = boto3.client('rekognition', 'us-east-1')
    queueUrl = 'https://sqs.us-east-1.amazonaws.com/744292932026/FullGroupVideos'
    roleArn = 'arn:aws:iam::744292932026:role/RekognitionServiceRole'
    topicArn = 'arn:aws:sns:us-east-1:744292932026:AmazonRekognitionFullGroupVideos'
    bucket = 'fancamgenerator'
    video = 'files/DALLA_DALLA.mp4'
    maxWidth = 0

    def main(self):
        self.triggerNewJob()
        #self.GetResultsPersons(self.jobId)

    def triggerNewJob(self):
        jobFound = False
        sqs = boto3.client('sqs', 'us-east-1')

        #=====================================
        response = self.rek.start_person_tracking(Video={'S3Object':{'Bucket':self.bucket,'Name':self.video}}, 
            NotificationChannel={'RoleArn':self.roleArn, 'SNSTopicArn':self.topicArn})
        #=====================================

        print(response)
        print('Start Job Id: ' + response['JobId'])
    #     dotLine = 0
    #     while jobFound == False:
    #         sqsResponse = sqs.receive_message(QueueUrl=self.queueUrl, MessageAttributeNames=['ALL'],
    #                                       MaxNumberOfMessages=10)

    #         if sqsResponse:
    #             if 'Messages' not in sqsResponse:
    #                 if dotLine<20:
    #                     print('.', end='')
    #                     dotLine=dotLine+1
    #                 else:
    #                     print()
    #                     dotLine=0    
    #                 sys.stdout.flush()
    #                 continue

    #             print(sqsResponse)
    #             for message in sqsResponse['Messages']:
    #                 notification = json.loads(message['Body'])
    #                 rekMessage = notification
    #                 # rekMessage = json.loads(notification['Message'])
    #                 print(rekMessage['JobId'])
    #                 print(rekMessage['Status'])
    #                 if str(rekMessage['JobId']) == response['JobId']:
    #                     print('Matching Job Found:' + rekMessage['JobId'])
    #                     jobFound = True
    #                     #=============================================
    #                     self.GetResultsPersons(rekMessage['JobId'])
    #                     #=============================================

    #                     sqs.delete_message(QueueUrl=self.queueUrl,
    #                                    ReceiptHandle=message['ReceiptHandle'])
    #                 else:
    #                     print("Job didn't match:" +
    #                           str(rekMessage['JobId']) + ' : ' + str(response['JobId']))
    #                 # Delete the unknown message. Consider sending to dead letter queue
    #                 sqs.delete_message(QueueUrl=self.queueUrl,
    #                                ReceiptHandle=message['ReceiptHandle'])

    #     print('done')

    # def foundWidth(self, width):
    #     if width > self.maxWidth:
    #         self.maxWidth = width

    # def GetResultsPersons(self, jobId):
    #     maxResults = 1000
    #     paginationToken = ''
    #     finished = False
    #     count = 0
    #     timestamps = set()

    #     while finished == False:
    #         response = self.rek.get_person_tracking(JobId=jobId,
    #                                         MaxResults=maxResults,
    #                                         NextToken=paginationToken)

    #         # print(response)
    #         print(response['VideoMetadata']['Codec'])
    #         print(str(response['VideoMetadata']['DurationMillis']))
    #         print(response['VideoMetadata']['Format'])
    #         print(response['VideoMetadata']['FrameRate'])

    #         for personDetection in response['Persons']:
    #             count += 1
    #             person = personDetection['Person']
    #             print('Index: ' + str(person['Index']))
    #             print('Timestamp: ' + str(personDetection['Timestamp']))
    #             timestamps.add(personDetection['Timestamp'])
    #             if 'BoundingBox' in person:
    #                 print ("      Bounding box")
    #                 print ("        Top: " + str(person['BoundingBox']['Top']))
    #                 print ("        Left: " + str(person['BoundingBox']['Left']))
    #                 print ("        Width: " +  str(person['BoundingBox']['Width']))
    #                 print ("        Height: " +  str(person['BoundingBox']['Height']))
    #                 self.foundWidth(person['BoundingBox']['Width'])
    #             print()

    #         if 'NextToken' in response:
    #             paginationToken = response['NextToken']
    #         else:
    #             finished = True
    #     print("Number of PersonDetection objects: " + str(count))
    #     print("Number of timestamps: " + str(len(timestamps)))
    #     print("Max width: " + str(self.maxWidth))
        
if __name__ == "__main__":

    analyzer = VideoDetect()
    analyzer.main()
