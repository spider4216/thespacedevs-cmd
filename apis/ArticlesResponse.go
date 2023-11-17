package apis

type ArticleResponse struct {
	Results []ArticleItem `json:"results"`
}

type ArticleItem struct {
	Id int               `json:"id"`
	Title string         `json:"title"`
	PublishedDate string `json:"published_at"`
	Description string   `json:"summary"`
}