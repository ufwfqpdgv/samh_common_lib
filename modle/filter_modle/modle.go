package filter_modle

var (
	url     string
	timeOut int
)

type FilterByRegexRequest struct {
	Content string `form:"content" json:"content" binding:"required"`
}

type FilterByRegexResponse struct {
	Rate int `json:"rate"`
}
