package main

import (
	"fmt"
	"github.com/txn2/txeh"
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
	"flag"
	"os"
)
func main(){
	// read the config file
	fmt.Println("Reading Rink Conf File Done!")

	domain := flag.String("domain", "", "domain name to add to /etc/hosts")
	port := flag.String("port", "", "PORT on which the application is running on!")
	flag.Parse()

	if *domain == "" || *port == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// edit the /etc/host file
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		panic(err)
	}
	hosts.AddHost("127.0.0.1", *domain)
	hosts.Save()

	//TODO: check if the user set the SSL to TRUE. In which case, form the URL with https.
	fullDomain := "http://" + *domain + ":" + *port


	// proxy traffic to the local web app
	fmt.Println("Proxing to the localhost running on PORT!")
	local, err := url.Parse(fullDomain)
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
