package openshot

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

var (
	project *Project
)

func TestGetProjects(t *testing.T) {
	defer projectsSetup()()

	project, err := createSampleProject(project)
	defer deleteSampleProject(project)

	if err != nil {
		t.Error(err)
	} else {
		log.Debug(project)
	}

	projects, err := openShot.GetProjects()
	if err != nil {
		t.Error(err)
	} else {
		log.Debug(projects)
	}

	if len(*projects) < 1 {
		t.Error("No projects were listed")
	}
}

func TestCreateProject(t *testing.T) {
	defer projectsSetup()()

	const sampleName = "Hello"
	const sampleWidth = 3840
	const sampleHeight = 2160
	project.Name = sampleName
	project.Width = sampleWidth
	project.Height = sampleHeight
	project, err := createSampleProject(project)

	if err != nil {
		t.Fatal("Error creating project ", err)
	}
	if project.Name != sampleName {
		t.Error("Corret project name not set")
	}
	if project.Width != sampleWidth {
		t.Error("Corret project width not set")
	}
	if project.Height != sampleHeight {
		t.Error("Corret project height not set")
	}
}

func projectsSetup() func() {
	project = &Project{Name: "Sample Name", JSON: "{}"}
	return projectsShutdown
}

func projectsShutdown() {
	project = nil
}

func createSampleProject(sampleProject *Project) (*Project, error) {
	return openShot.CreateProject(sampleProject)
}

func deleteSampleProject(sampleProject *Project) {
	// TODO implement
}
