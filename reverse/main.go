package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func main() {
	// set up our basic logging configuration
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logger := logrus.WithFields(logrus.Fields{
		"service": "go-reverse-proxy",
	})

	// create a new httprouter instance, we define the origin host (http://localhost:9000/)
	// and the ‘pattern’ we want httprouter to look out for (/*catchall, which is a special
	// syntax that represents a catchall wildcard/glob):
	router := httprouter.New()
	origin, _ := url.Parse("http://localhost:9000/")
	path := "/*catchall"

	// create a new reverse proxy instance, passing it the origin host
	reverseProxy := httputil.NewSingleHostReverseProxy(origin)

	// The director is simply a function that modifies the received incoming request,
	// while the response from the origin is copied back to the original client.
	reverseProxy.Director = func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host

		wildcardIndex := strings.IndexAny(path, "*")
		proxyPath := singleJoiningSlash(origin.Path, req.URL.Path[wildcardIndex:])
		if strings.HasSuffix(proxyPath, "/") && len(proxyPath) > 1 {
			proxyPath = proxyPath[:len(proxyPath)-1]
		}
		req.URL.Path = proxyPath
	}

	router.Handle("GET", path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		reverseProxy.ServeHTTP(w, r)
	})

	logger.Fatal(http.ListenAndServe(":9001", router))
}
