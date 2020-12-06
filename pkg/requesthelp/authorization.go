package requesthelp

import "net/http"

// SetAuthorization ...
func SetAuthorization(req *http.Request, token string) {
	var bearer = "Bearer " + token
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
}
