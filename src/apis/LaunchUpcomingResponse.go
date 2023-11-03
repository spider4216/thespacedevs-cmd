package apis

type LaunchUpcomingResponse struct {
	Results []LaunchInfo `json:"results"`
}

type LaunchInfo struct {
	Id string            `json:"id"`
	Name string          `json:"name"`
	Date string          `json:"net"`
	Status LaunchStatus  `json:"status"`
}

type LaunchStatus struct {
	Name string        `json:"name"`
	Description string `json:"description"`
}