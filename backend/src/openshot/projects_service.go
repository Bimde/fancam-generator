package openshot

import (
	"httputils"
)

const (
	projectsEndpoint = "/projects/"
)

// GetProjects returns a list of all projects created on the OpenShot server
func (o *OpenShot) GetProjects() (*[]Project, error) {
	log := getLogger("GetProjects")
	var projects Projects

	err := httputils.Get(log, baseURL+projectsEndpoint, nil, &projects)
	if err != nil {
		return nil, err
	}

	return &projects.Results, nil
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
