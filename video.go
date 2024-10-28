package pexels

type VideoParams struct {
	Query       string
	Orientation string
	Size        string
	Locale      string
	Page        int
	PerPage     int
}

type PopularParams struct {
	MinWidth    int
	MinHeight   int
	MinDuration int
	MaxDuration int
	Page        int
	PerPage     int
}

type SearchVideoResponse struct {
	TotalResults int      `json:"total_results"`
	Page         int      `json:"page"`
	PerPage      int      `json:"per_page"`
	Videos       []*Video `json:"videos"`
	NextPage     string   `json:"next_page"`
	PrevPage     string   `json:"prev_page"`
}

type PopularVideoResponse struct {
	SearchVideoResponse
}

type Video struct {
	ID            int            `json:"id"`
	Width         int            `json:"width"`
	Height        int            `json:"height"`
	URL           string         `json:"url"`
	Image         string         `json:"image"`
	FullRes       interface{}    `json:"full_res"`
	Tags          []interface{}  `json:"tags"`
	Duration      int            `json:"duration"`
	User          User           `json:"user"`
	VideoFiles    []VideoFile    `json:"video_files"`
	VideoPictures []VideoPicture `json:"video_pictures"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VideoFile struct {
	ID       int     `json:"id"`
	Quality  *string `json:"quality"`
	FileType string  `json:"file_type"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	FPS      float32 `json:"fps"`
	Link     string  `json:"link"`
	Size     int     `json:"size"`
}

type VideoPicture struct {
	ID      int    `json:"id"`
	Picture string `json:"picture"`
	Nr      int    `json:"nr"`
}
