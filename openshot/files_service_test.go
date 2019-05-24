package openshot

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	s3DefaultFilesFolder = "video_files/"
	testFileName         = "mini_test_file.mp4"
	s3DefaultBucket      = "openshot-sdk-go"
)

var sampleFile *FileUploadS3

func TestGetFiles(t *testing.T) {
	defer filesSetup(t)(t)
	createdFile := createSampleFile(t, project, sampleFile)
	defer deleteSampleFile(t, createdFile.ID)

	files := getFiles(t, project)
	if files.Count < 1 {
		t.Error("No files were returned")
	}
}

func TestFilesCreatedAndDeleted(t *testing.T) {
	defer filesSetup(t)(t)
	files := getFiles(t, project)
	createdFile := createSampleFile(t, project, sampleFile)

	newFiles := getFiles(t, project)
	if files.Count != newFiles.Count-1 {
		t.Error("file was not created")
	}

	deleteSampleFile(t, createdFile.ID)
	newFiles = getFiles(t, project)

	if files.Count != newFiles.Count {
		t.Error("file was not deleted")
	}
}

func getFiles(t *testing.T, project *Project) *Files {
	files, err := openShot.GetFiles(project)
	if err != nil {
		t.Error("error getting files ", err)
	}
	return files
}

func createSampleFile(t *testing.T, project *Project, file *FileUploadS3) *File {
	res, err := openShot.CreateFile(project, file)
	if err != nil {
		t.Fatal(err)
	}
	log.Debug(res)
	return res
}

func deleteSampleFile(t *testing.T, fileID int) {
	err := openShot.DeleteFile(fileID)
	if err != nil {
		t.Error(err)
	}
}

func filesSetup(t *testing.T) func(*testing.T) {
	projectsSetup()
	project = createSampleProject(t, project)
	sampleFile = CreateFileStruct(CreateFileS3InfoStruct(testFileName, s3DefaultFilesFolder, s3DefaultBucket))
	return filesShutdown
}

func filesShutdown(t *testing.T) {
	sampleFile = nil
	deleteSampleProject(t, project)
	projectsShutdown()
}
