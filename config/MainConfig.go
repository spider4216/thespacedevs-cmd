package config

type MainConfig struct {
	ApiDomain string
}

func Config() MainConfig {
	return MainConfig{ApiDomain: "https://lldev.thespacedevs.com/"}
}