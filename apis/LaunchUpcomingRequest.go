package apis

import (
	"ysirotenko.com/thespacedevs-cmd/config"
)

type LaunchUpcomingRequest struct {
	Method string
	Ver string
	Url string
	Domain string
	Page int
}

func GetLaunchUpcomingRequest() LaunchUpcomingRequest {
	return LaunchUpcomingRequest{Method: "GET", Ver: "2.2.0", Url: "/launch/upcoming/", Domain: config.Config().ApiDomain, Page: 1};
}
