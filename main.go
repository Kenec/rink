package main

import (
	"github.com/Kenec/rink/Editor"
	"github.com/Kenec/rink/Proxy"
	"github.com/Kenec/rink/configReader"
)

func main() {
	// config
	config := configReader.Config()

	// edit the /etc/host file
	Editor.HostEditor(config.Domain)

	// proxy traffic to the local web app
	Proxy.ProxyRoute(config.Domain, config.Port, config.Tls)
}
