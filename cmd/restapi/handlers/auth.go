package handlers

import (
	"fmt"
	"net/http"

	"github.com/julioc98/openbanking/pkg/interfaces"
)

// AuthHandler ...
type AuthHandler struct {
	authService interfaces.Authenticator
}

// NewAuthHandler factory
func NewAuthHandler(authService interfaces.Authenticator) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Auth a Auth
func (ah *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {

	authURL, err := ah.authService.GenerateAuthURL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "authUrl": "%s" }`, authURL)))
}

// Callback Auth
func (ah *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	accessToken, err := ah.authService.CallbackAuth(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "accessToken": "%s" }`, accessToken)))
}
