import boto3
import json
import sys

class VideoDetect:
    rek = boto3.client('rekognition', 'us-east-1')
    roleArn = 'arn:aws:iam::744292932026:role/RekognitionServiceRole'
    topicArn = 'arn:aws:sns:us-east-1:744292932026:AmazonRekognitionFullGroupVideos'
    bucket = 'fancamgenerator'
    video = 'files/DALLA_DALLA.mp4'

    def main(self):
        self.triggerNewJob()

    def triggerNewJob(self):
        response = self.rek.start_person_tracking(Video={'S3Object':{'Bucket':self.bucket,'Name':self.video}}, 
            NotificationChannel={'RoleArn':self.roleArn, 'SNSTopicArn':self.topicArn})
        print(response)
        print('Start Job Id: ' + response['JobId'])
        
if __name__ == "__main__":
    analyzer = VideoDetect()
    analyzer.main()
