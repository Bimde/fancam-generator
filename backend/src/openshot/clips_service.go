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
	err := httputils.Get(log, clipsURL(projectID), nil, &clips)
	if err != nil {
		return nil, err
	}
	return &clips, nil
}

// CreateClip creates a clip for the specified project
func (o *OpenShot) CreateClip(projectID int, clip *Clip) (*Clip, error) {
	log := getLogger("CreateClip")
	var createdClip Clip
	err := httputils.Post(log, clipsURL(projectID), clip, &createdClip)
	if err != nil {
		return nil, err
	}
	return &createdClip, nil
}

// UpdateClip updates a clip on the OpenShot server
func (o *OpenShot) UpdateClip(clip *Clip) (*Clip, error) {
	log := getLogger("UpdateClip")
	var updatedClip Clip
	err := httputils.Put(log, clipURL(clip.ID), clip, &updatedClip)
	if err != nil {
		return nil, err
	}
	return &updatedClip, nil
}

// GetClip gets the server version of the specified clip
func (o *OpenShot) GetClip(clipID int) (*Clip, error) {
	log := getLogger("GetClip")
	var clip Clip
	err := httputils.Get(log, clipURL(clipID), nil, &clip)
	if err != nil {
		return nil, err
	}
	return &clip, nil
}

// DeleteClip deletes the clip from openshot
func (o *OpenShot) DeleteClip(clipID int) error {
	log := getLogger("DeleteClip")
	return httputils.Delete(log, clipURL(clipID), nil, nil)
}

func CreateClipStruct(fileID int, projectID int) *Clip {
	return &Clip{FileURL: fileURL(fileID), ProjectURL: projectURL(projectID), JSON: map[string]interface{}{}}
}

func clipsURL(projectID int) string {
	return fmt.Sprintf(clipsEndpoint, projectID)
}

func clipURL(clipID int) string {
	return fmt.Sprintf(clipEndpoint, clipID)
}
