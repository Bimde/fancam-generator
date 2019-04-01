package main

import (
	"fmt"
	"openshot"

	log "github.com/sirupsen/logrus"
)

const (
	projectName = "Test Project #1"
	fileName    = "IDOL.mp4"
	scale       = 0
	fps         = 30
	height      = 1080
	width       = 1920
	frameWidth  = width / 5
)

var (
	openShot *openshot.OpenShot
	clients  map[int64]*OpenShot
)

func init() {
	openShot = openshot.New()
	clients = map[int64]*OpenShot{}
}

type OpenShot struct {
	project *openshot.Project
	file    *openshot.File
	clip    *openshot.Clip
}

func GetClient(ID int64) *OpenShot {
	if clients[ID] == nil {
		clients[ID] = newOpenShot()
	}
	return clients[ID]
}

func TriggerAllExports() *[]*openshot.Export {
	exports := make([]*openshot.Export, len(clients))
	for i, o := range clients {
		o.saveClip()
		export, err := o.createExport(deafultExport(o.project.ID))
		if err != nil {
			log.WithField("index", i).Error("error exporting project ", err)
			export = deafultExport(o.project.ID)
			export.URL = fmt.Sprintf("Export failed for projectID: %d, index: %d", o.project.ID, i)
		}
		exports[i] = export
	}
	return &exports
}

func (o *OpenShot) AddTrackingFrame(timestamp int64, width float64, left float64) {
	openShot.AddPropertyPoint(o.clip, openshot.LocationX, o.getFrame(timestamp), o.getLocationX(left))
}

func newOpenShot() *OpenShot {
	project := createProject(defaultProject())
	file := createFile(project.ID, defaultFile(fileName))
	clip := createClip(project.ID, defaultClip(file.ID, project.ID))
	return &OpenShot{project: project, file: file, clip: clip}
}

func (o *OpenShot) saveClip() error {
	var err error
	o.clip, err = openShot.UpdateClip(o.clip)
	return err
}

func (o *OpenShot) createExport(export *openshot.Export) (*openshot.Export, error) {
	export, err := openShot.CreateExport(o.project.ID, export)
	if err != nil {
		return nil, err
	}
	return export, nil
}

func (o *OpenShot) getLocationX(left float64) float64 {
	// TODO change hardcoded formula values to be proportional to frame / video dimensions
	return (1-left)*5 - 2.75
}

func (o *OpenShot) getFrame(timestamp int64) int {
	return int((float64(timestamp) / 1000.0) * float64(fps))
}

func createProject(project *openshot.Project) *openshot.Project {
	project, err := openShot.CreateProject(project)
	if err != nil {
		log.Panic("error creating project ", err)
	}
	return project
}

func createFile(projectID int, input *openshot.FileUploadS3) *openshot.File {
	file, err := openShot.CreateFile(projectID, input)
	if err != nil {
		log.Panic("error creating file ", err)
	}
	return file
}

func createClip(projectID int, input *openshot.Clip) *openshot.Clip {
	clip, err := openShot.CreateClip(projectID, input)
	if err != nil {
		log.Panic("error creating clip ", err)
	}
	openShot.SetScale(clip, scale)
	return clip
}

func defaultProject() *openshot.Project {
	return &openshot.Project{
		Name:           projectName,
		FPSNumerator:   fps,
		FPSDenominator: 1,
		Height:         height,
		Width:          width,
	}
}

func defaultFile(fileName string) *openshot.FileUploadS3 {
	return openshot.CreateFileStruct(fileName)
}

func defaultClip(fileID int, projectID int) *openshot.Clip {
	return openshot.CreateClipStruct(fileID, projectID)
}

func deafultExport(projectID int) *openshot.Export {
	o := openshot.CreateExportStruct(projectID)
	o.JSON["width"] = frameWidth
	return o
}
