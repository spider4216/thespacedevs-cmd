package apis

import (
	"ysirotenko.com/thespacedevs-cmd/config"
)

type ArticlesRequest struct {
	Method string
	Ver string
	Url string
	Domain string
	Page int
}

func GetArticlesRequest() ArticlesRequest {
	return ArticlesRequest{Method: "GET", Ver: "v4", Url: "/articles/", Domain: config.Config().ApiNewsDomain, Page: 1}
}