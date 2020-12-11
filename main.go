package main

import (
	"fmt"
	"github.com/txn2/txeh"
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
)
func main(){
	// read the config file
	fmt.Println("Reading Rink Conf File Done!")

	// edit the /etc/host file
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		panic(err)
	}
	hosts.AddHost("127.0.0.1", "dev.devito.com")
	hosts.Save()


	// proxy traffic to the local web app
	fmt.Println("Proxing to the localhost running on PORT!")
	local, err := url.Parse("http://dev.devito.com:3000")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(local)
	http.Handle("/", &ProxyHandler{proxy})
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("X-Ben", "Rad")
	ph.p.ServeHTTP(w, r)
}
