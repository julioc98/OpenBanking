package interfaces

import "net/http"

// Requester is an interface for &http.Client{}
type Requester interface {
	Do(req *http.Request) (*http.Response, error)
}
