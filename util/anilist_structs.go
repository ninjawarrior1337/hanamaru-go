package util

var SearchQueryTemplate = `query {
  Media(%v: %v, %v: %v) {
    siteUrl
    title {
      english
    }
    coverImage {
      large
    }
    type
    status
    seasonYear
    episodes
    duration
    chapters
    volumes
    description
    genres
    averageScore
  }
}`

type MediaType string

const (
	Manga MediaType = "MANGA"
	Anime           = "ANIME"
)

type ALRequest struct {
	Query string `json:"query"`
}

type ALResponse struct {
	Data struct {
		Media ALMedia `json:"media"`
	} `json:"data"`
}

type ALMedia struct {
	SiteURL string `json:"siteUrl"`
	Title   struct {
		English string `json:"english"`
	} `json:"title"`
	CoverImage struct {
		Large string `json:"large"`
	} `json:"coverImage"`
	Type         MediaType `json:"type"`
	Status       string    `json:"status"`
	SeasonYear   int       `json:"seasonYear"`
	Episodes     int       `json:"episodes"`
	Duration     int       `json:"duration"`
	Chapters     int       `json:"chapters"`
	Volumes      int       `json:"volumes"`
	Description  string    `json:"description"`
	Genres       []string  `json:"genres"`
	AverageScore int       `json:"averageScore"`
}
