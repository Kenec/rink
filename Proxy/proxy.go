package Proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type proxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("X-Ben", "Rad")
	ph.p.ServeHTTP(w, r)
}

func ProxyRoute(host string, port string, tls bool) {
	fmt.Println("Proxing to the localhost running on Port " + port)

	fullDomain := "http://" + host + ":" + port
	// TODO: check malformed url

	local, err := url.Parse(fullDomain)

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(local)
	http.Handle("/", &proxyHandler{proxy})

	if tls == true {
		err = http.ListenAndServeTLS(":443", "localhost.crt", "localhost.key", nil)
	} else {
		err = http.ListenAndServe(":80", nil)
	}

	if err != nil {
		panic(err)
	}
}
