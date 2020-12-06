package interfaces

import "github.com/julioc98/openbanking/pkg/entities"

// Provider is an interface for aobiebank gateway
type Provider interface {
	GetToken(user, pass string) (*entities.Token, error)
	GenerateAuthURL(user, prvtk, consentID string) (string, error)
	GetTokenByCode(user, pass, code string) (*entities.Token, error)
	CreateAccountAccessConsents(accountAccessConsents *entities.AccountAccessConsents, token string) (*entities.AccountAccessConsents, error)
}
