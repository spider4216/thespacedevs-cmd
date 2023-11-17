package commands

import (
	"strconv"
	"io/ioutil"
	"ysirotenko.com/thespacedevs-cmd/apis"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	FormatBold = "\033[1m"
	FormatReset = "\033[0m"
)

func LaunchesCommand(page *int) (string, error) {
	var result string = ""

	command := apis.GetLaunchUpcomingRequest()
	command.Page = *page * 10
	params := "?offset=" + strconv.Itoa(command.Page)
 	response, errHttp := http.Get(command.Domain + command.Ver + command.Url + params)
 	defer response.Body.Close()
 	
 	if errHttp != nil {
 		return result, errHttp
 	}

 	b, errReader := ioutil.ReadAll(response.Body)
 	
 	if errReader != nil {
 		return result, errReader
 	}
 	
 	var structmy apis.LaunchUpcomingResponse
 	errJson := json.Unmarshal(b, &structmy)
 	
 	if errJson != nil {
 		return result, errJson
 	}
 	
 	for _, entity := range structmy.Results {
 		result += fmt.Sprintf("%vID:%v %v\n", FormatBold, FormatReset, entity.Id)
 		result += fmt.Sprintf("%vName:%v %v\n", FormatBold, FormatReset, entity.Name)
 		result += fmt.Sprintf("%vDate:%v %v\n", FormatBold, FormatReset, entity.Date)
 		result += fmt.Sprintf("%vStatus:%v %v\n", FormatBold, FormatReset, entity.Status.Name)
 		result += fmt.Sprintf("%vStatus description:%v %v\n", FormatBold, FormatReset, entity.Status.Description)
 		result += fmt.Sprintf("%vAbout mission:%v %v\n", FormatBold, FormatReset, entity.Mission.Description)
 		result += fmt.Sprintf("*------------------*\n\n")
 	}
 	
 	return result, nil
}