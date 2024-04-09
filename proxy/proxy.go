package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"task-management/server/spec/authspec"
	"task-management/server/spec/customerspec"
	"task-management/server/spec/taskspec"
)

// Define routing rules
var routes = map[string]string{
	taskspec.BasePath:     fmt.Sprintf("http://localhost%s", taskspec.Host),
	authspec.BasePath:     fmt.Sprintf("http://localhost%s", authspec.Host),
	customerspec.BasePath: fmt.Sprintf("http://localhost%s", customerspec.Host),
}

func main() {
	// Create a new HTTP server
	http.HandleFunc("/", handleRequest)

	// Start the server on port 80
	http.ListenAndServe(":80", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi")
	// Get the path from the incoming request
	path := r.URL.Path
	var backendURL string

	for route, targetPath := range routes {
		if strings.HasPrefix(path, route) {
			backendURL = targetPath
			break
		}
	}
	fmt.Println(backendURL)

	// Parse the backend URL
	target, _ := url.Parse(backendURL)

	// Create a reverse proxy with the target URL
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Update the request URL before forwarding
	r.URL.Path = path
	r.Host = target.Host
	fmt.Println(r.Host, r.URL.Path)
	// Forward the request to the backend
	proxy.ServeHTTP(w, r)
}
