package handlers

import (
	"fmt"
	"log"
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	code := r.URL.Query().Get("code")
	log.Println("[DEBUG] Query Code:", code)
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := ah.authService.CallbackAuth(code)
	if err != nil {
		log.Println("[ERROR] message:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(fmt.Sprintf(`{ "accessToken": "%s" }`, accessToken)))
	w.Write([]byte(`
		<div style="display:flex; align-items: center;flex-direction: column;margin-top: 10%; font-weight: bold;">
		<div>Você foi logado com sucesso!</div>
		<div>Retorne para a aplicação.</div>
		</div>`))
}
