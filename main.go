package main

import (
	"flag"
	"ysirotenko.com/thespacedevs-cmd/commands"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Action list")
	page := flag.Int("page", 1, "Page number")
	flag.Parse()
	
	var err error
	var result string

	switch *action {
		case "launches":
			result, err = commands.LaunchesCommand(page)
	}
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(result)
}