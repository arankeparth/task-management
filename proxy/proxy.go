package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc(BasePath, handleRequest)
	http.ListenAndServe(ProxyPortHttp, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var backendURL string

	for route, targetPath := range Routes {
		if strings.HasPrefix(path, route) {
			backendURL = targetPath
			break
		}
	}
	// Parse the backend URL
	target, _ := url.Parse(backendURL)

	// Create a reverse proxy with the target URL
	proxy := httputil.NewSingleHostReverseProxy(target)

	fmt.Println("Proxying request to", target)
	// Update the request URL before forwarding
	r.URL.Path = path
	r.Host = target.Host
	// Forward the request to the backend
	proxy.ServeHTTP(w, r)
}
