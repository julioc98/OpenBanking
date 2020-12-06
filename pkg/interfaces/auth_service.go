package interfaces

// Authenticator service interface
type Authenticator interface {
	GenerateAuthURL() (string, error)
	CallbackAuth(code string) (string, error)
}
