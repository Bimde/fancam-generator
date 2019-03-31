package main

import "openshot"

const (
	projectName = "Test Project #1"
	fileName    = "ITZY.mp4"
	scale       = 0
	fps         = 30
	height      = 1080
	width       = 1920
	frameWidth  = width / 5
)

var (
	openShot = &openshot.OpenShot{}
	project  *openshot.Project
	file     *openshot.File
	clip     *openshot.Clip
)

func init() {
	log := getLogger("init")
	var err error
	project, err = openShot.CreateProject(
		&openshot.Project{
			Name:           projectName,
			FPSNumerator:   fps,
			FPSDenominator: 1,
			Height:         height,
			Width:          width,
		},
	)
	if err != nil {
		log.Error("error creating project ", err)
	}

	file, err = openShot.CreateFile(project.ID, openshot.CreateFileStruct(fileName))
	if err != nil {
		log.Error("error creating file ", err)
	}

	clip, err = openShot.CreateClip(project.ID, openshot.CreateClipStruct(file.ID, project.ID))
	if err != nil {
		log.Error("error creating clip ", err)
	}

	openShot.SetScale(clip, scale)
}

func addTrackingFrame(timestamp int64, width float64, left float64) {
	openShot.AddPropertyPoint(clip, openshot.LocationX, getFrame(timestamp), getLocationX(left))
}

func getLocationX(left float64) float64 {
	// TODO change hardcoded formula values to be proportional to frame / video dimensions
	return (1-left)*5 - 2.75
}

func getFrame(timestamp int64) int {
	return int((float64(timestamp) / 1000.0) * float64(fps))
}

func saveClip() error {
	var err error
	clip, err = openShot.UpdateClip(clip)
	return err
}
