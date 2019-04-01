package openshot

// response contains standard fields for retrieving lists of objects from OpenShot
type response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// Projects is the response data for /projects/
type Projects struct {
	response
	Results []Project `json:"results"`
}

// Files is the response data for /projects/{projectID}/files/
type Files struct {
	response
	Results []File `json:"results"`
}

// Clips is the response data for /projects/{projectID}/clips/
type Clips struct {
	response
	Results []Clip `json:"results"`
}

// Exports is the response data for /projects/{projectID}/exports/
type Exports struct {
	response
	Results []Export `json:"results"`
}

// Project represents http://cloud.openshot.org/doc/api_endpoints.html#projects
type Project struct {
	URL            string      `json:"url"`
	ID             int         `json:"id"`
	JSON           interface{} `json:"json"`
	Name           string      `json:"name"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	FPSNumerator   int         `json:"fps_num"`
	FPSDenominator int         `json:"fps_den"`
	SampleRate     int         `json:"sample_rate"`
	Channels       int         `json:"channels"`
	ChannelLayout  int         `json:"channel_layout"`
	FileURLs       []string    `json:"files"`
	ClipURLs       []string    `json:"clips"`
	ExportURLs     []string    `json:"exports"`
	Actions        []string    `json:"actions"`
	DateCreated    string      `json:"date_created"`
	DateUpdated    string      `json:"date_updated"`
}

// File represents http://cloud.openshot.org/doc/api_endpoints.html#files
type File struct {
	URL         string      `json:"url"`
	ID          int         `json:"id"`
	JSON        interface{} `json:"json"`
	Media       string      `json:"media"`
	Project     string      `json:"project"`
	Actions     []string    `json:"actions"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
}

// FileUploadS3 represents http://cloud.openshot.org/doc/api_endpoints.html#files
type FileUploadS3 struct {
	ProjectURL string     `json:"project"`
	JSON       FileS3Info `json:"json"`
}

// FileS3Info represents http://cloud.openshot.org/doc/api_endpoints.html#id20
type FileS3Info struct {
	URL    string `json:"url"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// Clip represents http://cloud.openshot.org/doc/api_endpoints.html#clips
type Clip struct {
	URL        string                 `json:"url"`
	ID         int                    `json:"id"`
	JSON       map[string]interface{} `json:"json"`
	FileURL    string                 `json:"file"`
	Position   float32                `json:"position"`
	Start      float32                `json:"start"`
	End        float32                `json:"end"`
	Actions    []string               `json:"actions"`
	ProjectURL string                 `json:"project"`
}

// Export represents http://cloud.openshot.org/doc/api_endpoints.html#exports
type Export struct {
	URL          string                 `json:"url"`
	ID           int                    `json:"id"`
	JSON         map[string]interface{} `json:"json"`
	Output       string                 `json:"output"`
	ExportType   string                 `json:"export_type"`
	VideoFormat  string                 `json:"video_format"`
	VideoCodec   string                 `json:"video_codec"`
	VideoBitrate int                    `json:"video_bitrate"`
	AudioCodec   string                 `json:"ac3"`
	AudioBitrate int                    `json:"audio_bitrate"`
	StartFrame   int                    `json:"start_frame"`
	EndFrame     int                    `json:"end_frame"`
	Actions      []string               `json:"actions"`
	ProjectURL   string                 `json:"project"`
	Webhook      string                 `json:"webhook"`
	Progress     float32                `json:"progress"`
	Status       string                 `json:"status"`
	DateCreated  string                 `json:"date_created"`
	DateUpdated  string                 `json:"date_updated"`
}

// Property represents the JSON structure of points in animations
// More details at http://cloud.openshot.org/doc/animation.html#json-structure
type Property struct {
	Points []Point `json:"Points"`
}

type Point struct {
	Co            Cord `json:"co"`
	Interpolation int  `json:"interpolation"`
}

type Cord struct {
	X int     `json:"X"` // frame number
	Y float64 `json:"Y"` // value
}
