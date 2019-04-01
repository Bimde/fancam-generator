package openshot

import (
	"fmt"
	"httputils"
)

const (
	exportsEndpoint = baseURL + "/projects/%d/exports/"
	exportEndpoint  = baseURL + "/exports/%d/"
)

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
func (o *OpenShot) CreateExport(projectID int, input *Export) (*Export, error) {
	log := getLogger("CreateExport")
	var export Export
	err := httputils.Post(log, exportsURL(projectID), input, &export)
	if err != nil {
		return nil, err
	}
	return &export, nil
}

// GetExports returns a list of all exports created for a particular project
func (o *OpenShot) GetExports(projectID int) (*Exports, error) {
	log := getLogger("GetExports")
	var exports Exports
	err := httputils.Get(log, exportsURL(projectID), nil, &exports)
	if err != nil {
		return nil, err
	}
	return &exports, nil
}

// DeleteExport deletes the export from openshot
func (o *OpenShot) DeleteExport(exportID int) error {
	log := getLogger("DeleteExport")
	return httputils.Delete(log, exportURL(exportID), nil, nil)
}

func CreateExportStruct(projectID int) *Export {
	return &Export{
		ExportType:   exportType,
		VideoFormat:  videoFormat,
		VideoCodec:   videoCodec,
		VideoBitrate: 8000000,
		AudioCodec:   audioCodec,
		AudioBitrate: audioBitrate,
		StartFrame:   startFrame,
		ProjectURL:   projectURL(projectID),
		JSON:         map[string]interface{}{},
		Status:       status,
	}
}

func exportsURL(projectID int) string {
	return fmt.Sprintf(exportsEndpoint, projectID)
}

func exportURL(exportID int) string {
	return fmt.Sprintf(exportEndpoint, exportID)
}
