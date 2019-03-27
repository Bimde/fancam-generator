package openshot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	projectsEndpoint = "/projects/"
)

// GetProjects returns a list of all projects created
func (o *OpenShot) GetProjects() (*[]Project, error) {
	req := o.createReqWithAuth("GET", projectsEndpoint, nil)
	resp, err := o.client.Do(req)

	if err != nil {
		log.Fatal("OpenShot#GetProjects error executing request ", err)
		return nil, err
	}

	log.Println("Response: ", resp)
	log.Println("Response Body: ", resp.Body)

	projects := &Projects{}
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, projects)

	if err != nil {
		log.Fatal("OpenShot#GetProjects error unmarshalling projects ", err)
		return nil, err
	}

	return &projects.Results, nil
}

// CreateProject creates the given project
func (o *OpenShot) CreateProject(project *Project) (*Project, error) {
	data, err := json.Marshal(project)
	if err != nil {
		log.Fatal("OpenShot#createProject error marshalling project ", err)
		return nil, err
	}

	req := o.createReqWithAuth("POST", baseURL+projectsEndpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	res, err := o.client.Do(req)
	if err != nil {
		log.Fatal("OpenShot#createProject error executing request ", err)
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("OpenShot#createProject Error reading response body ", err)
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Fatal("OpenShot#createProject request error response ", res)
		log.Fatal("OpenShot#createProject request error response body", string(resBody))
		return nil, fmt.Errorf("Error creating project %s, OpenShot server response: %s", project.Name, string(resBody))
	}

	var createdProject Project
	err = json.Unmarshal(resBody, &createdProject)

	if err != nil {
		log.Fatalf("OpenShot#createProject error unmarshalling response %s \n with error %s", string(resBody), err)

	}

	return &createdProject, nil
}
