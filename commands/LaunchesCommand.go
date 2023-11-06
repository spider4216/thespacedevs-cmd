package commands

import (
	"ysirotenko.com/thespacedevs-cmd/utils"
	"strconv"
	"io/ioutil"
	"ysirotenko.com/thespacedevs-cmd/apis"
	"encoding/json"
	"fmt"
	"net/http"
)

func LaunchesCommand(page *int) {
	launchApi := utils.Init()
	launchApi.Page = *page * 10
	params := "?offset=" + strconv.Itoa(launchApi.Page)
 	response, _ := http.Get(launchApi.Domain + launchApi.Ver + launchApi.Url + params)
 	defer response.Body.Close()
 	// todo err
 	// todo status code
 	b, err := ioutil.ReadAll(response.Body)
 	
 	if err != nil {
 		panic("err 0")
 	}
 	
 	var structmy apis.LaunchUpcomingResponse
 	err = json.Unmarshal(b, &structmy)
 	
 	if err != nil {
 		panic("err")
 	}
 	// todo err
 	
 	for _, entity := range structmy.Results {
 		fmt.Printf("ID: %v\nName: %v\nDate:%v\nStatus: %v\nStatus Description: %v\n\n", entity.Id, entity.Name, entity.Date, entity.Status.Name, entity.Status.Description)
 	}
}