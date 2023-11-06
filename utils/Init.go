package utils

import (
	"ysirotenko.com/thespacedevs-cmd/apis"
	"ysirotenko.com/thespacedevs-cmd/config"
)

func Init() apis.LaunchUpcoming {
	mainConfig := config.MainConfig{ApiDomain: "https://lldev.thespacedevs.com/"}
	launchUpcoming := apis.LaunchUpcoming{Method: "GET", Ver: "2.2.0", Url: "/launch/upcoming/", Domain: mainConfig.ApiDomain, Page: 1}
	
	return launchUpcoming
}