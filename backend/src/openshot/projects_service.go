package openshot

import (
	"fmt"
	"httputils"
)

const (
	projectsEndpoint = "/projects/"
	projectEndpoint  = projectsEndpoint + "%d/"
)

const (
	defaultProjectWidth          = 1920
	defaultProjectHeight         = 1080
	defaultProjectFPSNumerator   = 30
	defaultProjectFPSDenominator = 1
	defaultAudioSampleRate       = 44100
	defaultAudioChannels         = 2 // 2=stereo
	defaultAudioChannelLayout    = 3 // 3=stereo, 4=mono, 7=Surround
)

// GetProjects returns a list of all projects created on the OpenShot server
func (o *OpenShot) GetProjects() (*Projects, error) {
	log := getLogger("GetProjects")
	var projects Projects

	err := httputils.Get(log, projectsURL(), nil, &projects)
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

// CreateProject creates the given project on the OpenShot server
func (o *OpenShot) CreateProject(project *Project) (*Project, error) {
	log := getLogger("CreateProject").WithField("projectName", project.Name)
	fillDefaults(project)
	log.Info("Creating project ", *project)
	var createdProject Project

	err := httputils.Post(log, projectsURL(), project, &createdProject)
	if err != nil {
		return nil, err
	}

	return &createdProject, nil
}

// DeleteProject deletes a project on the OpenShot server.
// Note that this deletion will trigger deletion of all associated files and clips.
// There is also no (easy) way to recover a deleted project so this method should
// only be exposed to trusted sources and through an "are you sure" or equivalent
// confirmation dialog.
func (o *OpenShot) DeleteProject(projectID int) error {
	log := getLogger("GetProjects")
	return httputils.Delete(log, projectURL(projectID), nil, nil)
}

func projectsURL() string {
	return baseURL + projectsEndpoint
}

func projectURL(projectID int) string {
	return fmt.Sprintf(baseURL+projectEndpoint, projectID)
}

func fillDefaults(project *Project) {
	if project.FPSNumerator == 0 {
		project.FPSNumerator = defaultProjectFPSNumerator
	}
	if project.FPSDenominator == 0 {
		project.FPSDenominator = defaultProjectFPSDenominator
	}
	if project.Width == 0 {
		project.Width = defaultProjectWidth
	}
	if project.Height == 0 {
		project.Height = defaultProjectHeight
	}
	if project.SampleRate == 0 {
		project.SampleRate = defaultAudioSampleRate
	}
	if project.Channels == 0 {
		project.Channels = defaultAudioChannels
	}
	if project.ChannelLayout == 0 {
		project.ChannelLayout = defaultAudioChannelLayout
	}
	if project.JSON == nil {
		project.JSON = map[string]interface{}{}
	}
}
