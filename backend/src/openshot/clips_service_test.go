package openshot

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

var (
	createdFile *File
	sampleClip  *Clip
)

func TestGetClips(t *testing.T) {
	defer clipsSetup(t)(t)
	createdClip := createSampleClip(t, project.ID, sampleClip)
	defer deleteSampleClip(t, createdClip.ID)

	clips := getClips(t, project.ID)
	if clips.Count < 1 {
		t.Error("No clips were returned")
	}
}

func TestClipsCreatedAndDeleted(t *testing.T) {
	defer clipsSetup(t)(t)
	clips := getClips(t, project.ID)
	createdClip := createSampleClip(t, project.ID, sampleClip)

	newClips := getClips(t, project.ID)
	if clips.Count != newClips.Count-1 {
		t.Error("clip was not created")
	}

	deleteSampleClip(t, createdClip.ID)
	newClips = getClips(t, project.ID)

	if clips.Count != newClips.Count {
		t.Error("clip was not deleted")
	}
}

func getClips(t *testing.T, projectID int) *Clips {
	clips, err := openShot.GetClips(project.ID)
	if err != nil {
		t.Error("error getting files ", err)
	}
	return clips
}

func createSampleClip(t *testing.T, projectID int, clip *Clip) *Clip {
	res, err := openShot.CreateClip(projectID, clip)
	if err != nil {
		t.Fatal(err)
	}
	log.Debug(res)
	return res
}

func deleteSampleClip(t *testing.T, clipID int) {
	err := openShot.DeleteClip(clipID)
	if err != nil {
		t.Error(err)
	}
}

func clipsSetup(t *testing.T) func(*testing.T) {
	filesSetup(t)
	createdFile = createSampleFile(t, project.ID, sampleFile)
	sampleClip = createClipStruct(createdFile.ID, project.ID)
	return clipsShutdown
}

func clipsShutdown(t *testing.T) {
	sampleClip = nil
	deleteSampleFile(t, createdFile.ID)
	createdFile = nil
	filesShutdown(t)
}
