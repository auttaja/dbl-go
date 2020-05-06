// Package provides a Golang interface to the top.gg API
package dbl_go

import "net/http"

const (
	// The current version of the package
	Version = "0.0.1"
	// The top.gg base URL
	BaseURL = "https://top.gg"
)

// New creates a new top.gg client using the given token and the default http client
func New(token string) *DBL {
	return NewDBL(token, &http.Client{})
}
