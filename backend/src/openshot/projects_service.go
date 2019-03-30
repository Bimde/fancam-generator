package openshot

import (
	"fmt"
	"httputils"
)

const (
	projectsEndpoint = "/projects/"
	projectEndpoint  = projectsEndpoint + "%d/"
)

// GetProjects returns a list of all projects created on the OpenShot server
func (o *OpenShot) GetProjects() (*Projects, error) {
	log := getLogger("GetProjects")
	var projects Projects

	err := httputils.Get(log, baseURL+projectsEndpoint, nil, &projects)
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

// CreateProject creates the given project on the OpenShot server
func (o *OpenShot) CreateProject(project *Project) (*Project, error) {
	log := getLogger("CreateProject").WithField("projectName", project.Name)
	log.Info("Creating project ", *project)
	var createdProject Project

	err := httputils.Post(log, baseURL+projectsEndpoint, project, &createdProject)
	if err != nil {
		return nil, err
	}

	return &createdProject, nil
}

// DeleteProject deletes a project on the OpenShot server.
// Note that this deletion will trigger deletion of all associated files and clips.
// There is also no (easy) way to recover a deleted project so this endpoints should
// only be accessible to trusted sources and through an "are you sure" or equivalent
// confirmation dialog.
func (o *OpenShot) DeleteProject(projectID int) error {
	log := getLogger("GetProjects")

	err := httputils.Delete(log, fmt.Sprintf(baseURL+projectEndpoint, projectID), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
