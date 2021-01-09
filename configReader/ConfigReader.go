package configReader

import (
	"flag"
	"log"
)

type configParam struct {
	Domain string
	Port   string
	Tls    bool
}

func validateConfig(configValue configParam) {
	switch true {
	case configValue.Domain == "":
		log.Fatal("Domain is required")
	case configValue.Port == "":
		log.Fatal("Port is required")
	}
}

func Config() configParam {
	domain := flag.String("domain", "", "domain name to add to /etc/hosts")
	port := flag.String("port", "", "PORT on which the application is running on!")
	tls := flag.Bool("tls", false, "Set TLS/SSL for the domain")
	flag.Parse()

	config := configParam{*domain, *port, *tls}
	validateConfig(config)

	return config
}
