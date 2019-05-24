package openshot

import (
	"fmt"
)

const (
	exportsEndpoint = "/projects/%d/exports/"
	exportEndpoint  = "/exports/%d/"
)

// Values for default exports
const (
	exportType   = "video"      // video, audio, image, waveform
	videoFormat  = "mp4"        // mp4, avi, %05d.png, jpg, etc...
	videoCodec   = "libx264"    // h.264, mp4video, etc...
	videoBitrate = 8000000      // 8000000 = 7.6 MB / sec
	audioCodec   = "libfdk_aac" // libmp3lame, ac3, libfdk_aac, etc...
	audioBitrate = 1920000      // 1920000 = 192 KB / sec
	startFrame   = 1
	status       = "pending" // not sure why this is a mandatory input
)

// CreateExport triggers exporting the specified project with the given export settings
// on the OpenShot server
func (o *OpenShot) CreateExport(project *Project, input *Export) (*Export, error) {
	log := getLogger("CreateExport")
	log.Debug("Creating export ", *input)
	var export Export
	err := o.http.Post(log, o.exportsURL(project.ID), input, &export)
	if err != nil {
		return nil, err
	}
	return &export, nil
}

// GetExports returns a list of all exports created for a particular project
func (o *OpenShot) GetExports(projectID int) (*Exports, error) {
	log := getLogger("GetExports")
	var exports Exports
	err := o.http.Get(log, o.exportsURL(projectID), nil, &exports)
	if err != nil {
		return nil, err
	}
	return &exports, nil
}

// DeleteExport deletes the export from openshot
func (o *OpenShot) DeleteExport(exportID int) error {
	log := getLogger("DeleteExport")
	return o.http.Delete(log, o.exportURL(exportID), nil, nil)
}

// GetExport gets the server version of the specified export
func (o *OpenShot) GetExport(exportID int) (*Export, error) {
	log := getLogger("GetExport")
	var export Export
	err := o.http.Get(log, o.exportURL(exportID), nil, &export)
	if err != nil {
		return nil, err
	}
	return &export, nil
}

// CreateDefaultExportStruct creates an Export struct with default settings
func CreateDefaultExportStruct(project *Project) *Export {
	return &Export{
		ExportType:   exportType,
		VideoFormat:  videoFormat,
		VideoCodec:   videoCodec,
		VideoBitrate: 8000000,
		AudioCodec:   audioCodec,
		AudioBitrate: audioBitrate,
		StartFrame:   startFrame,
		ProjectURL:   project.URL,
		JSON:         map[string]interface{}{},
		Status:       status,
	}
}

func (o *OpenShot) exportsURL(projectID int) string {
	return fmt.Sprintf(o.BaseURL+exportsEndpoint, projectID)
}

func (o *OpenShot) exportURL(exportID int) string {
	return fmt.Sprintf(o.BaseURL+exportEndpoint, exportID)
}
