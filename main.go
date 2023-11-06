package main

import (
	"flag"
	"ysirotenko.com/thespacedevs-cmd/commands"
)

func main() {
	action := flag.String("action", "", "Action list")
	page := flag.Int("page", 1, "Page number")
	flag.Parse()

	switch *action {
		case "launches":
			commands.LaunchesCommand(page)
	}
}