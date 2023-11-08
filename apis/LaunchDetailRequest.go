package apis

import (
	"ysirotenko.com/thespacedevs-cmd/config"
)

type LaunchDetailRequest struct {
	Method string
	Ver string
	Url string
	Domain string
}

func GetLaunchDetailRequest(id *string) LaunchDetailRequest {
	launchId := *id
	return LaunchDetailRequest{Method: "GET", Ver: "2.2.0", Url: "/launch/" + launchId, Domain: config.Config().ApiDomain};
}
