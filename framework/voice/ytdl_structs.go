package voice

type Video struct {
	ID                 string             `json:"id"`
	Uploader           string             `json:"uploader"`
	UploaderID         string             `json:"uploader_id"`
	UploaderURL        string             `json:"uploader_url"`
	ChannelID          string             `json:"channel_id"`
	ChannelURL         string             `json:"channel_url"`
	UploadDate         string             `json:"upload_date"`
	License            interface{}        `json:"license"`
	Creator            string             `json:"creator"`
	Title              string             `json:"title"`
	AltTitle           string             `json:"alt_title"`
	Thumbnail          string             `json:"thumbnail"`
	Description        string             `json:"description"`
	Categories         []string           `json:"categories"`
	Tags               []string           `json:"tags"`
	Subtitles          Subtitles          `json:"subtitles"`
	AutomaticCaptions  AutomaticCaptions  `json:"automatic_captions"`
	Duration           int                `json:"duration"`
	AgeLimit           int                `json:"age_limit"`
	Annotations        interface{}        `json:"annotations"`
	Chapters           interface{}        `json:"chapters"`
	WebpageURL         string             `json:"webpage_url"`
	ViewCount          int                `json:"view_count"`
	LikeCount          int                `json:"like_count"`
	DislikeCount       int                `json:"dislike_count"`
	AverageRating      float64            `json:"average_rating"`
	Formats            []Formats          `json:"formats"`
	IsLive             interface{}        `json:"is_live"`
	StartTime          interface{}        `json:"start_time"`
	EndTime            interface{}        `json:"end_time"`
	Series             interface{}        `json:"series"`
	SeasonNumber       interface{}        `json:"season_number"`
	EpisodeNumber      interface{}        `json:"episode_number"`
	Track              string             `json:"track"`
	Artist             string             `json:"artist"`
	Album              string             `json:"album"`
	ReleaseDate        interface{}        `json:"release_date"`
	ReleaseYear        interface{}        `json:"release_year"`
	Extractor          string             `json:"extractor"`
	WebpageURLBasename string             `json:"webpage_url_basename"`
	ExtractorKey       string             `json:"extractor_key"`
	Playlist           interface{}        `json:"playlist"`
	PlaylistIndex      interface{}        `json:"playlist_index"`
	Thumbnails         []Thumbnails       `json:"thumbnails"`
	DisplayID          string             `json:"display_id"`
	RequestedSubtitles interface{}        `json:"requested_subtitles"`
	RequestedFormats   []RequestedFormats `json:"requested_formats"`
	Format             string             `json:"format"`
	FormatID           string             `json:"format_id"`
	Width              int                `json:"width"`
	Height             int                `json:"height"`
	Resolution         interface{}        `json:"resolution"`
	Fps                int                `json:"fps"`
	Vcodec             string             `json:"vcodec"`
	Vbr                interface{}        `json:"vbr"`
	StretchedRatio     interface{}        `json:"stretched_ratio"`
	Acodec             string             `json:"acodec"`
	Abr                int                `json:"abr"`
	Ext                string             `json:"ext"`
	Fulltitle          string             `json:"fulltitle"`
	Filename           string             `json:"_filename"`
}
type Subtitles struct {
}
type AutomaticCaptions struct {
}
type DownloaderOptions struct {
	HTTPChunkSize int `json:"http_chunk_size"`
}
type HTTPHeaders struct {
	UserAgent      string `json:"User-Agent"`
	AcceptCharset  string `json:"Accept-Charset"`
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	AcceptLanguage string `json:"Accept-Language"`
}
type Formats struct {
	FormatID          string            `json:"format_id"`
	URL               string            `json:"url"`
	PlayerURL         string            `json:"player_url"`
	Ext               string            `json:"ext"`
	FormatNote        string            `json:"format_note"`
	Acodec            string            `json:"acodec"`
	Abr               int               `json:"abr,omitempty"`
	Asr               int               `json:"asr"`
	Filesize          int               `json:"filesize"`
	Fps               interface{}       `json:"fps"`
	Height            interface{}       `json:"height"`
	Tbr               float64           `json:"tbr"`
	Width             interface{}       `json:"width"`
	Vcodec            string            `json:"vcodec"`
	DownloaderOptions DownloaderOptions `json:"downloader_options,omitempty"`
	Format            string            `json:"format"`
	Protocol          string            `json:"protocol"`
	HTTPHeaders       HTTPHeaders       `json:"http_headers"`
	Container         string            `json:"container,omitempty"`
}
type Thumbnails struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}
type RequestedFormats struct {
	FormatID          string            `json:"format_id"`
	URL               string            `json:"url"`
	PlayerURL         string            `json:"player_url"`
	Ext               string            `json:"ext"`
	Height            int               `json:"height"`
	FormatNote        string            `json:"format_note"`
	Vcodec            string            `json:"vcodec"`
	Asr               interface{}       `json:"asr"`
	Filesize          int               `json:"filesize"`
	Fps               int               `json:"fps"`
	Tbr               float64           `json:"tbr"`
	Width             int               `json:"width"`
	Acodec            string            `json:"acodec"`
	DownloaderOptions DownloaderOptions `json:"downloader_options"`
	Format            string            `json:"format"`
	Protocol          string            `json:"protocol"`
	HTTPHeaders       HTTPHeaders       `json:"http_headers"`
	Abr               int               `json:"abr,omitempty"`
}
