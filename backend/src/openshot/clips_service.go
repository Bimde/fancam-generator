package openshot

import (
	"fmt"
	"httputils"
)

const (
	clipsEndpoint = baseURL + "/projects/%d/clips/"
	clipEndpoint  = baseURL + "/clips/%d/"
)

// GetClips returns a list of all clips created for a particular project
func (o *OpenShot) GetClips(projectID int) (*Clips, error) {
	log := getLogger("GetClips")
	var clips Clips
	httputils.Get(log, fmt.Sprintf(clipsEndpoint, projectID), nil, &clips)
	return &clips, nil
}

// CreateClip creates a clip for the specified project
func (o *OpenShot) CreateClip(projectID int, clip *Clip) (*Clip, error) {
	log := getLogger("CreateClip")
	var createdClip Clip
	httputils.Post(log, fmt.Sprintf(clipsEndpoint, projectID), clip, &createdClip)
	return &createdClip, nil
}

// DeleteClip deletes the clip from openshot
func (o *OpenShot) DeleteClip(clipID int) error {
	log := getLogger("DeleteClip")
	return httputils.Delete(log, fmt.Sprintf(clipEndpoint, clipID), nil, nil)
}

/*
{
    "file": "http://cloud.openshot.org/files/135/",
    "project": "http://cloud.openshot.org/projects/146/",
    "json": {}
}
*/
func createClipStruct(fileID int, projectID int) *Clip {
	return &Clip{FileURL: fileURL(fileID), ProjectURL: projectURL(projectID), JSON: map[string]string{}}
}
