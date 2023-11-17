package config

type MainConfig struct {
	ApiDomain string
	ApiNewsDomain string
}

func Config() MainConfig {
	return MainConfig{ApiDomain: "https://lldev.thespacedevs.com/", ApiNewsDomain: "https://api.spaceflightnewsapi.net/"}
}