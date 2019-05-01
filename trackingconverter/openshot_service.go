package main

import (
	"fmt"
	"math"

	"github.com/Bimde/fancam-generator/config"

	"github.com/Bimde/openshot-sdk-go/openshot"

	log "github.com/sirupsen/logrus"
)

const (
	projectName     = "Test Project #1"
	defaultFileName = "BOSS.mp4"
	deafultS3Folder = "files/"
	deafultS3Bucket = "fancamgenerator"
	openshotURL     = "http://cloud.openshot.org/"
	scale           = 0
	fps             = 30
	height          = 1080
	width           = 1920
	frameWidth      = width / 5
)

var (
	openShot *openshot.OpenShot
	clients  map[int64]*OpenShot
)

func init() {
	openShot = openshot.New(openshotURL, config.GetString(config.Username), config.GetString(config.Password))
	clients = map[int64]*OpenShot{}
}

type OpenShot struct {
	project *openshot.Project
	file    *openshot.File
	clip    *openshot.Clip
}

// GetClient returns the OpenShot client associated with the particular id.
// Creates a new client and the mapping to the id if it doesn't exist. This
// function is purpose-built for dealing a large number of projects at once,
// i.e when dealing with aws rekognition people pathing results
func GetClient(ID int64) *OpenShot {
	if clients[ID] == nil {
		clients[ID] = newOpenShot()
	}
	return clients[ID]
}

// TriggerAllExports starts exporting all projects associated with any OpenShot clients
// created through GetClient.
func TriggerAllExports() *[]*openshot.Export {
	exports := make([]*openshot.Export, len(clients))
	for index, client := range clients {
		exports[index] = triggerExport(index, deafultExport(client.project), client)
	}
	return &exports
}

// TriggerAllExportsTrimmed provides the same functionality as TriggerAllExports,
// except with each export being trimmed to the range of entries for all properties.
func TriggerAllExportsTrimmed() *[]*openshot.Export {
	exports := make([]*openshot.Export, len(clients))
	for index, client := range clients {
		export := deafultExport(client.project)
		trim(export, client.clip)
		exports[index] = triggerExport(index, export, client)
	}
	return &exports
}

func triggerExport(index int64, export *openshot.Export, client *OpenShot) *openshot.Export {
	client.saveClip()
	export, err := client.createExport(export)
	if err != nil {
		log.WithField("index", index).Error("error exporting project ", err)
		export = deafultExport(client.project)
		export.URL = fmt.Sprintf("Export failed for projectID: %d, index: %d", client.project.ID, index)
	}
	return export
}

func newOpenShot() *OpenShot {
	project := createProject(defaultProject())
	file := createFile(project, defaultFile())
	clip := createClip(project, defaultClip(file, project))
	return &OpenShot{project: project, file: file, clip: clip}
}

func (o *OpenShot) AddTrackingFrame(timestamp int64, width float64, left float64) {
	openShot.AddPropertyPoint(o.clip, openshot.LocationX, o.getFrame(timestamp), o.getLocationX(left))
}

func (o *OpenShot) saveClip() error {
	var err error
	o.clip, err = openShot.UpdateClip(o.clip)
	return err
}

func (o *OpenShot) createExport(export *openshot.Export) (*openshot.Export, error) {
	export, err := openShot.CreateExport(o.project, export)
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

func createFile(project *openshot.Project, input *openshot.FileUploadS3) *openshot.File {
	file, err := openShot.CreateFile(project, input)
	if err != nil {
		log.Panic("error creating file ", err)
	}
	return file
}

// createClip creates a clip uses openshot, sets scale and clears LocationX
func createClip(project *openshot.Project, input *openshot.Clip) *openshot.Clip {
	clip, err := openShot.CreateClip(project, input)
	if err != nil {
		log.Panic("error creating clip ", err)
	}
	openShot.SetScale(clip, scale)
	openShot.ClearPropertyPoints(clip, openshot.LocationX)
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

func defaultFile() *openshot.FileUploadS3 {
	return openshot.CreateFileStruct(openshot.CreateFileS3InfoStruct(defaultFileName, deafultS3Folder, deafultS3Bucket))
}

func defaultClip(file *openshot.File, project *openshot.Project) *openshot.Clip {
	return openshot.CreateClipStruct(file, project)
}

func deafultExport(project *openshot.Project) *openshot.Export {
	o := openshot.CreateDefaultExportStruct(project)
	o.JSON["width"] = frameWidth
	return o
}

func trim(export *openshot.Export, clip *openshot.Clip) {
	property := openShot.GetProperty(clip, openshot.LocationX)
	export.StartFrame = getFirstFrame(property)
	export.EndFrame = getLastFrame(property)
	log.WithFields(log.Fields{
		"project":    clip.ProjectURL,
		"startFrame": export.StartFrame,
		"endFrame":   export.EndFrame,
	}).Infof(
		"seconds: %f",
		((float64(export.EndFrame) - float64(export.StartFrame)) / fps),
	)
}

func getFirstFrame(property *openshot.Property) int {
	min := math.MaxInt32
	for _, point := range property.Points {
		if point.Co.X < min {
			min = point.Co.X
		}
	}
	return min
}

func getLastFrame(property *openshot.Property) int {
	max := 0
	for _, point := range property.Points {
		if point.Co.X > max {
			max = point.Co.X
		}
	}
	return max
}
