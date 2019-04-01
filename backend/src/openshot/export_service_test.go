package openshot

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

var sampleExport *Export

func TestGetExports(t *testing.T) {
	defer exportsSetup(t)(t)
	createdExport := createSampleExport(t, project.ID, sampleExport)
	defer deleteSampleExport(t, createdExport.ID)

	exports := getExports(t, project.ID)
	if exports.Count < 1 {
		t.Error("No clips were returned")
	}
}

func TestExportsCreatedAndDeleted(t *testing.T) {
	defer exportsSetup(t)(t)
	exports := getExports(t, project.ID)
	createdExport := createSampleExport(t, project.ID, sampleExport)

	newExports := getExports(t, project.ID)
	if exports.Count != newExports.Count-1 {
		t.Error("clip was not created")
	}

	deleteSampleExport(t, createdExport.ID)
	newExports = getExports(t, project.ID)

	if exports.Count != newExports.Count {
		t.Error("clip was not deleted")
	}
}

func getExports(t *testing.T, projectID int) *Exports {
	exports, err := openShot.GetExports(projectID)
	if err != nil {
		t.Error("error getting exports ", err)
	}
	return exports
}

func createSampleExport(t *testing.T, projectID int, export *Export) *Export {
	res, err := openShot.CreateExport(projectID, export)
	if err != nil {
		t.Fatal(err)
	}
	log.Debug(res)
	return res
}

func deleteSampleExport(t *testing.T, exportID int) {
	err := openShot.DeleteExport(exportID)
	if err != nil {
		t.Error(err)
	}
}

func exportsSetup(t *testing.T) func(*testing.T) {
	clipsSetup(t)
	sampleClip = createSampleClip(t, project.ID, sampleClip)
	sampleExport = CreateExportStruct(project.ID)
	return exportsShutdown
}

func exportsShutdown(t *testing.T) {
	sampleExport = nil
	sampleClip = nil
	clipsShutdown(t)
}
