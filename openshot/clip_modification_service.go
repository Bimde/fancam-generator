package openshot

import (
	"github.com/mitchellh/mapstructure"
)

// Constants representing clip JSON property names
// More details at: http://cloud.openshot.org/doc/api_endpoints.html#clips
const (
	LocationX = "location_x"
)

const (
	// http://cloud.openshot.org/doc/animation.html#key-frames
	// 0=Bézier, 1=Linear, 2=Constant
	interpolationMode = 1
)

// SetScale sets the scale of the provided clip object
// DOES NOT set value on server, requires call to UpdateClip
func (o *OpenShot) SetScale(clip *Clip, scale int) {
	clip.JSON["scale"] = scale
}

// AddPropertyPoint sets a JSON property of the provided clip object at the specified frame.
// DOES NOT set value on server, requires call to UpdateClip
func (o *OpenShot) AddPropertyPoint(clip *Clip, key string, frame int, value float64) {
	property := o.GetProperty(clip, key)
	property.Points = append(property.Points, Point{Co: Cord{X: frame, Y: value}, Interpolation: interpolationMode})
	clip.JSON[key] = property
}

// ClearPropertyPoints clears all property point entires from clip for the specified
// property key.
func (o *OpenShot) ClearPropertyPoints(clip *Clip, key string) {
	property := o.GetProperty(clip, key)
	property.Points = []Point{}
	clip.JSON[key] = property
}

// GetProperty returns a json object type-asserted to an openshot.Property object
func (o *OpenShot) GetProperty(clip *Clip, key string) *Property {
	var property Property
	mapstructure.Decode(clip.JSON[key], &property)
	return &property
}
