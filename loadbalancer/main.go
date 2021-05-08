package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var severCount = 0

const (
	SERVER1 = "https://www.google.com"
	SERVER2 = "https://www.facebook.com"
	SERVER3 = "https://www.yahoo.com"
	PORT    = "1338"
)

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxyNewSingleHostReverseProxy is the function in the httputil package
	// which returns ReverseProxy struct and then calls ServeHTTP() where this package has the main
	// logic of reverse proxy to forward a request to backend servers.
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(w, r)
}

func logRequestPayload(proxyURL string) {
	log.Printf("proxy_url: %s\n", proxyURL)
}

// Round-robin algorithm
func getProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}

	server := servers[severCount]
	severCount++

	if severCount >= len(servers) {
		severCount = 0
	}

	return server
}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	url := getProxyURL()

	logRequestPayload(url)

	serveReverseProxy(url, w, r)
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
