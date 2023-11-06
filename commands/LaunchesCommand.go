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

const (
	FormatBold = "\033[1m"
	FormatReset = "\033[0m"
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
 		fmt.Printf("%vID:%v %v\n", FormatBold, FormatReset, entity.Id)
 		fmt.Printf("%vName:%v %v\n", FormatBold, FormatReset, entity.Name)
 		fmt.Printf("%vDate:%v %v\n", FormatBold, FormatReset, entity.Date)
 		fmt.Printf("%vStatus:%v %v\n", FormatBold, FormatReset, entity.Status.Name)
 		fmt.Printf("%vDescription:%v %v\n", FormatBold, FormatReset, entity.Status.Description)
 		fmt.Printf("*------------------*\n\n")
 	}
}