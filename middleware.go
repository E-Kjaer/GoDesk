package main

import "net/http"

// JSONWrapper is a middleware handler that adds content type = json on all requests.
type JSONWrapper struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *JSONWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	l.handler.ServeHTTP(w, r)
}

// NewJSONWrapper constructs a new Logger middleware handler
func NewJSONWrapper(handlerToWrap http.Handler) *JSONWrapper {
	return &JSONWrapper{handlerToWrap}
}
