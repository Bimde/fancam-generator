package openshot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	projectsEndpoint = "/projects/"
)

// GetProjects returns a list of all projects created
func (o *OpenShot) GetProjects() (*[]Project, error) {
	log := getLogger("GetProjects")

	req := o.createReqWithAuth("GET", projectsEndpoint, nil)
	resp, err := o.client.Do(req)

	if err != nil {
		log.Error("error executing request ", err)
		return nil, err
	}

	log.Info("Response: ", resp)
	log.Info("Response Body: ", resp.Body)

	var projects Projects
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &projects)

	if err != nil {
		log.Error("error unmarshalling projects ", err)
		return nil, err
	}

	return &projects.Results, nil
}

// CreateProject creates the given project
func (o *OpenShot) CreateProject(project *Project) (*Project, error) {
	log := getLogger("CreateProject").WithField("projectName", project.Name)

	log.Info("Creating project ", *project)

	data, err := json.Marshal(project)
	if err != nil {
		log.Error("error marshalling project ", err)
		return nil, err
	}

	req := o.createReqWithAuth("POST", projectsEndpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	res, err := o.client.Do(req)
	if err != nil {
		log.Error("error executing request ", err)
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("error reading response body ", err)
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Error("error request response status", res)
		log.Error("response body ", string(resBody))
		return nil, fmt.Errorf("Error creating project %s, OpenShot server response: %s", project.Name, string(resBody))
	}

	var createdProject Project
	err = json.Unmarshal(resBody, &createdProject)

	if err != nil {
		log.WithField("responseBody", string(resBody)).Error("error unmarshalling response ", err)
	}

	return &createdProject, nil
}
