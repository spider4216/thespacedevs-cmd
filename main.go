package main

import (
	"flag"
	"ysirotenko.com/thespacedevs-cmd/commands"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Action list")
	page := flag.Int("page", 1, "Page number")
	detailId := flag.String("launch_detail_id", "", "Launch detail id")
	flag.Parse()
	
	var err error
	var result string

	switch *action {
		case "launches":
			result, err = commands.LaunchesCommand(page)
		case "launch_detail":
			result, err = commands.LaunchDetailCommand(detailId)
	}
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(result)
}