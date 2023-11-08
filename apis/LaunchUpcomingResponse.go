package apis

type LaunchUpcomingResponse struct {
	Results []LaunchInfo `json:"results"`
}

type LaunchInfo struct {
	Id string             `json:"id"`
	Name string           `json:"name"`
	Date string           `json:"net"`
	Status LaunchStatus   `json:"status"`
	Mission LaunchMission `json:mission`
}

type LaunchStatus struct {
	Name string        `json:"name"`
	Description string `json:"description"`
}

type LaunchMission struct {
	Name string        `json:"name"`
	Description string `json:"description"`
}