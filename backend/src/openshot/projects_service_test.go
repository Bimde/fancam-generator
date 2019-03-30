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

	project := createSampleProject(t, project)
	defer deleteSampleProject(t, project)

	projects := getProjects(t)

	if projects.Count < 1 {
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
	project := createSampleProject(t, project)
	defer deleteSampleProject(t, project)

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

func TestProjectCreatedAndDeleted(t *testing.T) {
	projects := getProjects(t)

	defer projectsSetup()()
	project = createSampleProject(t, project)

	newProjects := getProjects(t)
	if projects.Count != newProjects.Count-1 {
		t.Error("Project was not created")
	}

	deleteSampleProject(t, project)
	newProjects = getProjects(t)

	if projects.Count != newProjects.Count {
		t.Error("Project was not deleted")
	}
}

func projectsSetup() func() {
	project = &Project{Name: "Sample Name", JSON: "{}"}
	return projectsShutdown
}

func projectsShutdown() {
	project = nil
}

func getProjects(t *testing.T) *Projects {
	projects, err := openShot.GetProjects()
	if err != nil {
		t.Error(err)
	} else {
		log.Debug(projects)
	}

	return projects
}

func createSampleProject(t *testing.T, sampleProject *Project) *Project {
	res, err := openShot.CreateProject(sampleProject)
	if err != nil {
		t.Fatal(err)
	} else {
		log.Debug(res)
	}
	return res
}

func deleteSampleProject(t *testing.T, sampleProject *Project) {
	err := openShot.DeleteProject(sampleProject.ID)
	if err != nil {
		t.Error(err)
	}
}
