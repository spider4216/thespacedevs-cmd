package apis

type LaunchDetailResponse struct {
	Name string                  `json:"name"`
	Updates []LaunchDetailUpdate `json:"updates"`
}

type LaunchDetailUpdate struct {
	Comment string `json:"comment"`
	Date string    `json:"created_on"`
	InfoUrl string `json:"info_url"`
}