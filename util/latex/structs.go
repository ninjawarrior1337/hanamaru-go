package latex

type Status string

const (
	Success Status = "success"
	Error   Status = "error"
)

type LatexRequest struct {
	Code    string `json:"code"`
	Format  string `json:"format"`
	Quality int    `json:"quality,omitempty"`
	Density int    `json:"density,omitempty"`
}

type LatexResponse struct {
	Status      Status `json:"status"`
	Log         string `json:"log"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
}
