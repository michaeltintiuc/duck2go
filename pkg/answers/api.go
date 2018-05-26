package answers

import (
	"fmt"
)

const (
	// BaseURL of the DuckDuckGo Instant Answers API
	BaseURL = "https://api.duckduckgo.com/?t=duck2go"
)

type Query struct {
	url      string
	Response Response
}

type Response struct {
}

func newQuery(s string) *Query {
	q := &Query{generateURL(s, true, false), Response{}}
	return q
}

func generateURL(s string, json bool, prettyPrint bool) string {
	url := fmt.Sprintf("%s?q=%s", BaseURL, s)
	return url
}
