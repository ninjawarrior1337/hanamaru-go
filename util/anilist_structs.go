package util

var SearchQuery = `query {
  Media(search: "%v") {
    siteUrl
    title {
      english
    }
    coverImage {
      large
    }
    status
    duration
    seasonYear
    episodes
    description
    genres
    averageScore
  }
}`

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
	Status       string   `json:"status"`
	Duration     int      `json:"duration"`
	SeasonYear   int      `json:"seasonYear"`
	Episodes     int      `json:"episodes"`
	Description  string   `json:"description"`
	Genres       []string `json:"genres"`
	AverageScore int      `json:"averageScore"`
}
