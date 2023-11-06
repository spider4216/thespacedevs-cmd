package main

import (
	"flag"
	"ysirotenko.com/thespacedevs-cmd/commands"
)

const (
	FormatBold = "\033[1m"
	FormatReset = "\033[0m"
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