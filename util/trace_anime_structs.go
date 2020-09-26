package util

type TraceMoeResp struct {
	RawDocsCount      int            `json:"RawDocsCount"`
	RawDocsSearchTime int            `json:"RawDocsSearchTime"`
	ReRankSearchTime  int            `json:"ReRankSearchTime"`
	CacheHit          bool           `json:"CacheHit"`
	Trial             int            `json:"trial"`
	Limit             int            `json:"limit"`
	LimitTTL          int            `json:"limit_ttl"`
	Quota             int            `json:"quota"`
	QuotaTTL          int            `json:"quota_ttl"`
	Docs              []TraceMoeDocs `json:"docs"`
}
type TraceMoeDocs struct {
	From            float64       `json:"from"`
	To              float64       `json:"to"`
	AnilistID       int           `json:"anilist_id"`
	At              float64       `json:"at"`
	Season          string        `json:"season"`
	Anime           string        `json:"anime"`
	Filename        string        `json:"filename"`
	Episode         int           `json:"episode"`
	Tokenthumb      string        `json:"tokenthumb"`
	Similarity      float64       `json:"similarity"`
	Title           string        `json:"title"`
	TitleNative     string        `json:"title_native"`
	TitleChinese    string        `json:"title_chinese"`
	TitleEnglish    string        `json:"title_english"`
	TitleRomaji     string        `json:"title_romaji"`
	MalID           int           `json:"mal_id"`
	Synonyms        []string      `json:"synonyms"`
	SynonymsChinese []interface{} `json:"synonyms_chinese"`
	IsAdult         bool          `json:"is_adult"`
}
