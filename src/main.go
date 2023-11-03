package main

import (
	"fmt"
	"net/http"
	"utils"
// 	"encoding/json"
	"io/ioutil"
	"apis"
	"encoding/json"
	"flag"
	"strconv"
)

const (
	FormatBold = "\033[1m"
	FormatReset = "\033[0m"
)

func main() {
	launchApi := utils.Init()
	
	action := flag.String("action", "", "Action list")
	page := flag.Int("page", 1, "Page number")
	flag.Parse()

	switch *action {
		case "launches":
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
}