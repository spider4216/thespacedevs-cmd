package commands

import (
	"net/http"
	"ysirotenko.com/thespacedevs-cmd/apis"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func ArticlesCommand(page *int) (string, error) {
	var result string = ""

	articlesRequest := apis.GetArticlesRequest()
	articlesRequest.Page = *page * 10
	params := "?offset=" + strconv.Itoa(articlesRequest.Page)

	response, errHttp := http.Get(articlesRequest.Domain + articlesRequest.Ver + articlesRequest.Url + params)
	defer response.Body.Close()
	
	if errHttp != nil {
		return result, errHttp
	}
	
	content, errReader := ioutil.ReadAll(response.Body)
	
	if errReader != nil {
		return result, errReader
	}
	
	var responseStruct apis.ArticleResponse
	errJson := json.Unmarshal(content, &responseStruct)
	
	if errJson != nil {
		return result, errJson
	}
	
	for _, entity := range responseStruct.Results {
		result += fmt.Sprintf("%vTitle:%v\n%v\n", FormatBold, FormatReset, entity.Title);
		result += fmt.Sprintf("%vId:%v\n%v\n", FormatBold, FormatReset, entity.Id);
		result += fmt.Sprintf("%vPublished:%v\n%v\n", FormatBold, FormatReset, entity.PublishedDate);
		result += fmt.Sprintf("%vDescription:%v\n%v\n\n", FormatBold, FormatReset, entity.Description);
	}
	
	return result, nil;
}