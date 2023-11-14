package commands

import (
	"io/ioutil"
	"ysirotenko.com/thespacedevs-cmd/apis"
	"encoding/json"
	"fmt"
	"net/http"
)

func LaunchDetailCommand(id *string) (string, error) {
	var result string = ""

	command := apis.GetLaunchDetailRequest(id)
 	response, errHttp := http.Get(command.Domain + command.Ver + command.Url)
 	defer response.Body.Close()
 	
 	if errHttp != nil {
 		return result, errHttp
 	}

 	b, errReader := ioutil.ReadAll(response.Body)
 	
 	if errReader != nil {
 		return result, errReader
 	}
 	
 	var entity apis.LaunchDetailResponse
 	errJson := json.Unmarshal(b, &entity)
 	
 	if errJson != nil {
 		return result, errJson
 	}

	result += fmt.Sprintf("%vName%v\n%v\n\n", FormatBold, FormatReset, entity.Mission.Name)
	result += fmt.Sprintf("%vDescription%v\n%v\n\n", FormatBold, FormatReset, entity.Mission.Description)
	result += fmt.Sprintf("%vUPDATES%v\n\n", FormatBold, FormatReset)
	
	for _, update := range entity.Updates {
		result += fmt.Sprintf("%vDate:%v %v\n", FormatBold, FormatReset, update.Date)
		result += fmt.Sprintf("%vComment:%v %v\n", FormatBold, FormatReset, update.Comment)
		result += fmt.Sprintf("%vMore:%v %v\n\n", FormatBold, FormatReset, update.InfoUrl)
	}
 	
 	return result, nil
}