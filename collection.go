package pexels

type CollectionParams struct {
	Type string
	PageParams
}

type CollectionResponse struct {
	TotalResults int           `json:"total_results"`
	Page         int           `json:"page"`
	PerPage      int           `json:"per_page"`
	Collections  []*Collection `json:"collections"`
	NextPage     string        `json:"next_page"`
	PrevPage     string        `json:"prev_page"`
}

type CollectionMediaResponse struct {
	ID           string  `json:"id"`
	Media        []Media `json:"media"`
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	TotalResults int     `json:"total_results"`
	NextPage     string  `json:"next_page"`
	PrevPage     string  `json:"prev_page"`
}

type Media struct {
	Type            string         `json:"type"`
	ID              int            `json:"id"`
	Width           int            `json:"width"`
	Height          int            `json:"height"`
	URL             string         `json:"url"`
	Photographer    string         `json:"photographer,omitempty"`
	PhotographerURL string         `json:"photographer_url,omitempty"`
	PhotographerID  int            `json:"photographer_id,omitempty"`
	AvgColor        string         `json:"avg_color"`
	Src             Source         `json:"src,omitempty"`
	Liked           bool           `json:"liked,omitempty"`
	Duration        int            `json:"duration,omitempty"`
	FullRes         interface{}    `json:"full_res,omitempty"`
	Tags            []interface{}  `json:"tags,omitempty"`
	Image           string         `json:"image,omitempty"`
	User            User           `json:"user,omitempty"`
	VideoFiles      []VideoFile    `json:"video_files,omitempty"`
	VideoPictures   []VideoPicture `json:"video_pictures,omitempty"`
}

type Collection struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	MediaCount  int    `json:"media_count"`
	PhotosCount int    `json:"photos_count"`
	VideosCount int    `json:"videos_count"`
}
