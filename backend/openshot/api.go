package openshot

type Projects struct {
	Count   int       `json:"count"`
	Results []Project `json:"results"`
}

// http://cloud.openshot.org/doc/api_endpoints.html#projects
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
	Files          []File      `json:"files"`
	Clips          []Clip      `json:"clips"`
	Exports        []Export    `json:"exports"`
	Actions        []string    `json:"actions"`
	DateCreated    string      `json:"date_created"`
	DateUpdated    string      `json:"date_updated"`
}

// http://cloud.openshot.org/doc/api_endpoints.html#files
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

// http://cloud.openshot.org/doc/api_endpoints.html#files
type FileUploadS3 struct {
	URL  string     `json:"url"`
	JSON FileS3Info `json:"json"`
}

// http://cloud.openshot.org/doc/api_endpoints.html#id20
type FileS3Info struct {
	URL    string `json:"url"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// http://cloud.openshot.org/doc/api_endpoints.html#clips
type Clip struct {
	URL      string      `json:"url"`
	ID       int         `json:"id"`
	JSON     interface{} `json:"json"`
	File     string      `json:"file"`
	Position float32     `json:"position"`
	Start    float32     `json:"start"`
	End      float32     `json:"end"`
	Actions  []string    `json:"actions"`
	Project  string      `json:"project"`
}

// http://cloud.openshot.org/doc/api_endpoints.html#exports
type Export struct {
	URL          string      `json:"url"`
	ID           int         `json:"id"`
	JSON         interface{} `json:"json"`
	Output       string      `json:"output"`
	ExportType   string      `json:"export_type"`
	VideoFormat  string      `json:"video_format"`
	VideoCodec   string      `json:"video_codec"`
	VideoBitrate int         `json:"video_bitrate"`
	AudioCodec   string      `json:"ac3"`
	AudioBitrate int         `json:"audio_bitrate"`
	StartFrame   int         `json:"start_frame"`
	EndFrame     int         `json:"end_frame"`
	Actions      []string    `json:"actions"`
	Project      string      `json:"project"`
	Webhook      string      `json:"webhook"`
	Progress     float32     `json:"progress"`
	Status       string      `json:"status"`
	DateCreated  string      `json:"date_created"`
	DateUpdated  string      `json:"date_updated"`
}
