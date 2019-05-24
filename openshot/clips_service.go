package openshot

import (
	"fmt"
)

const (
	clipsEndpoint = "/projects/%d/clips/"
	clipEndpoint  = "/clips/%d/"
)

// GetClips returns a list of all clips created for a particular project
func (o *OpenShot) GetClips(projectID int) (*Clips, error) {
	log := getLogger("GetClips")
	var clips Clips
	err := o.http.Get(log, o.clipsURL(projectID), nil, &clips)
	if err != nil {
		return nil, err
	}
	return &clips, nil
}

// CreateClip creates a clip for the specified project
func (o *OpenShot) CreateClip(project *Project, clip *Clip) (*Clip, error) {
	log := getLogger("CreateClip")
	log.Debug("Creating clip ", *clip)
	var createdClip Clip
	err := o.http.Post(log, o.clipsURL(project.ID), clip, &createdClip)
	if err != nil {
		return nil, err
	}
	return &createdClip, nil
}

// UpdateClip updates a clip on the OpenShot server
func (o *OpenShot) UpdateClip(clip *Clip) (*Clip, error) {
	log := getLogger("UpdateClip")
	var updatedClip Clip
	err := o.http.Put(log, o.clipURL(clip.ID), clip, &updatedClip)
	if err != nil {
		return nil, err
	}
	return &updatedClip, nil
}

// GetClip gets the server version of the specified clip
func (o *OpenShot) GetClip(clipID int) (*Clip, error) {
	log := getLogger("GetClip")
	var clip Clip
	err := o.http.Get(log, o.clipURL(clipID), nil, &clip)
	if err != nil {
		return nil, err
	}
	return &clip, nil
}

// DeleteClip deletes the clip from openshot
func (o *OpenShot) DeleteClip(clipID int) error {
	log := getLogger("DeleteClip")
	return o.http.Delete(log, o.clipURL(clipID), nil, nil)
}

// CreateClipStruct creates a minimum Clip struct required for input to CreateClip
func CreateClipStruct(file *File, project *Project) *Clip {
	return &Clip{FileURL: file.URL, ProjectURL: project.URL, JSON: map[string]interface{}{}}
}

func (o *OpenShot) clipsURL(projectID int) string {
	return fmt.Sprintf(o.BaseURL+clipsEndpoint, projectID)
}

func (o *OpenShot) clipURL(clipID int) string {
	return fmt.Sprintf(o.BaseURL+clipEndpoint, clipID)
}
