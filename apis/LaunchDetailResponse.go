package apis

type LaunchDetailResponse struct {
	Name string                  `json:"name"`
	Updates []LaunchDetailUpdate `json:"updates"`
	Mission LaunchDetailMission  `json:"mission"`
}

type LaunchDetailUpdate struct {
	Comment string `json:"comment"`
	Date string    `json:"created_on"`
	InfoUrl string `json:"info_url"`
}

type LaunchDetailMission struct {
	Description string `json:"description"`
}