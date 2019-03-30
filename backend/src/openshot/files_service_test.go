package openshot

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const testFileName = "mini_test_file.mp4"

var sampleFile *FileUploadS3

func TestGetFiles(t *testing.T) {
	defer filesSetup(t)(t)
	createdFile := createSampleFile(t, project.ID, sampleFile)
	defer deleteSampleFile(t, createdFile.ID)

	files := getFiles(t, project.ID)
	if files.Count < 1 {
		t.Error("No files were returned")
	}
}

func TestFilesCreatedAndDeleted(t *testing.T) {
	defer filesSetup(t)(t)
	files := getFiles(t, project.ID)
	createdFile := createSampleFile(t, project.ID, sampleFile)

	newFiles := getFiles(t, project.ID)
	if files.Count != newFiles.Count-1 {
		t.Error("file was not created")
	}

	deleteSampleFile(t, createdFile.ID)
	newFiles = getFiles(t, project.ID)

	if newFiles.Count != newFiles.Count {
		t.Error("file was not deleted")
	}
}

func getFiles(t *testing.T, projectID int) *Files {
	files, err := openShot.GetFiles(project.ID)
	if err != nil {
		t.Error("error getting files ", err)
	}
	return files
}

func createSampleFile(t *testing.T, projectID int, file *FileUploadS3) *File {
	res, err := openShot.CreateFile(projectID, file)
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
	sampleFile = createFileStruct(testFileName)
	return filesShutdown
}

func filesShutdown(t *testing.T) {
	deleteSampleProject(t, project)
	projectsShutdown()
}
