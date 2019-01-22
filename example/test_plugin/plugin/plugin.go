package plugin

import (
	"fmt"
	"net/http"
)

var version = "v1"

// Sample stores some plugin private metadata
type Sample struct{ Name string }

// Handler is a Sample method to processes HTTP requests
func (s *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my website", s.Name, version)
}

// NewSample returns a new sample handler function
func NewSample(name string) func(http.ResponseWriter, *http.Request) {
	s := &Sample{"test"}
	fmt.Println("in NewSample", name, version, s)
	return s.ServeHTTP
}

// NewSampleHandler returns a new sample handler function
func NewSampleHandler(name string) http.Handler {
	s := &Sample{"test"}
	fmt.Println("in NewSample", name, version, s)
	return s
}
