package pexels

type PhotoParams struct {
	Query       string
	Orientation string
	Size        string
	Color       string
	Locale      string
	Page        int
	PerPage     int
}

type PageParams struct {
	Page    int
	PerPage int
}

type SearchPhotoResponse struct {
	TotalResults int      `json:"total_results"`
	Page         int      `json:"page"`
	PerPage      int      `json:"per_page"`
	Photos       []*Photo `json:"photos"`
	NextPage     string   `json:"next_page"`
	PrevPage     string   `json:"prev_page"`
}

type CuratedResponse struct {
	Page     int      `json:"page"`
	PerPage  int      `json:"per_page"`
	Photos   []*Photo `json:"photos"`
	NextPage string   `json:"next_page"`
}

type PhotoResponse struct {
	Photos *Photo `json:"-"`
}

type Photo struct {
	ID              int    `json:"id"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	URL             string `json:"url"`
	Photographer    string `json:"photographer"`
	PhotographerURL string `json:"photographer_url"`
	PhotographerID  int    `json:"photographer_id"`
	AvgColor        string `json:"avg_color"`
	Liked           bool   `json:"liked"`
	Src             Source `json:"src"`
}

type Source struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}
