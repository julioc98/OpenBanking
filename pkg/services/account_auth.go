package services

import (
	"time"

	"github.com/julioc98/openbanking/pkg/entities"
	"github.com/julioc98/openbanking/pkg/interfaces"
)

// AccountAuth Use Case
type AccountAuth struct {
	provider   interfaces.Provider
	user       string
	pass       string
	privateKey string
}

// NewAccountAuth factory
func NewAccountAuth(provider interfaces.Provider, user, pass, privateKey string) *AccountAuth {
	return &AccountAuth{
		provider:   provider,
		user:       user,
		pass:       pass,
		privateKey: privateKey,
	}
}

// CreateAccessConsents greate account access
func (a *AccountAuth) CreateAccessConsents() (*entities.AccountAccessConsents, error) {
	token, err := a.provider.GetToken(a.user, a.pass)
	if err != nil {
		return nil, err
	}
	accountAccessConsents := &entities.AccountAccessConsents{
		Data: entities.Data{
			Permissions: []string{
				"ReadAccountsBasic",
				"ReadAccountsDetail",
				"ReadBalances",
				"ReadBeneficiariesDetail",
				"ReadDirectDebits",
				"ReadProducts",
				"ReadStandingOrdersDetail",
				"ReadTransactionsCredits",
				"ReadTransactionsDebits",
				"ReadTransactionsDetail",
				"ReadOffers",
				"ReadPAN",
				"ReadParty",
				"ReadPartyPSU",
				"ReadScheduledPaymentsDetail",
				"ReadStatementsDetail",
			},
			ExpirationDateTime:      time.Now().Add(30 * 24 * 60 * time.Minute),
			TransactionFromDateTime: time.Now(),
			TransactionToDateTime:   time.Now(),
		},
		Risk: entities.Risk{},
	}

	aac, err := a.provider.CreateAccountAccessConsents(accountAccessConsents, token.AccessToken)
	if err != nil {
		return nil, err

	}
	return aac, nil
}

// GenerateAuthURL return a URL auth
func (a *AccountAuth) GenerateAuthURL() (string, error) {
	aac, err := a.CreateAccessConsents()
	if err != nil {
		return "", err
	}
	return a.provider.GenerateAuthURL(a.user, a.privateKey, aac.Data.ConsentID)
}

// CallbackAuth return a token
func (a *AccountAuth) CallbackAuth(code string) (string, error) {
	token, err := a.provider.GetTokenByCode(a.user, a.pass, code)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
