package openshot

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestModifyClip(t *testing.T) {
	defer clipsSetup(t)(t)
	createdClip := createSampleClip(t, project, sampleClip)
	defer deleteSampleClip(t, createdClip.ID)

	clip := getClip(t, createdClip.ID)
	openShot.SetScale(clip, 0)
	openShot.AddPropertyPoint(clip, LocationX, 1, 0.125)
	openShot.AddPropertyPoint(clip, LocationX, 7, 0.2673)
	localProperty := openShot.GetProperty(clip, LocationX)

	updateClip(t, clip)
	updatedClip := getClip(t, clip.ID)
	serverProperty := openShot.GetProperty(updatedClip, LocationX)

	if !cmp.Equal(localProperty, serverProperty) {
		t.Error("property not updated server")
	}
}
