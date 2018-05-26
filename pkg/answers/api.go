package answers

import (
	"fmt"
)

const (
	// BaseURL of the DuckDuckGo Instant Answers API
	BaseURL = "https://api.duckduckgo.com/?t=duck2go"
)

// Request to be sent
type Request struct {
	url      string
	Response Response
}

// Response to the API query
type Response struct {
}

// NewQuery creates a Request instance
func NewQuery(q string, json bool) *Request {
	return &Request{generateURL(q, json), Response{}}
}

// generateURL for the API query
func generateURL(s string, json bool) string {
	format := "xml"
	if json {
		format = "json"
	}
	return fmt.Sprintf("%s&q=%s&format=%s", BaseURL, s, format)
}
