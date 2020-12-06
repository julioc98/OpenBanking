package services

import (
	"net/http"
	"testing"

	"github.com/julioc98/openbanking/pkg/gateways"
	"github.com/julioc98/openbanking/pkg/interfaces"
	"github.com/julioc98/openbanking/tests"
)

const (
	baseURL     = "https://gw-dev.obiebank.banfico.com/obie-aisp/v3.1/aisp"
	authURL     = "https://auth.obiebank.banfico.com"
	redirectURI = "https://openbankinghacka.herokuapp.com/callback"
	user        = "PSDBR-NCA-AISP01"
	pass        = "senha123"
	wantToken   = ""
)

var (
	client = http.DefaultClient
	prov   = gateways.NewObiebank(authURL, baseURL, redirectURI, client)
)

func TestAccountAuth_CreateAccessConsents(t *testing.T) {
	type fields struct {
		provider interfaces.Provider
		user     string
		pass     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Create OK",
			fields: fields{
				provider: prov,
				user:     user,
				pass:     pass,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AccountAuth{
				provider: tt.fields.provider,
				user:     tt.fields.user,
				pass:     tt.fields.pass,
			}
			if _, err := a.CreateAccessConsents(); (err != nil) != tt.wantErr {
				t.Errorf("AccountAuth.CreateAccessConsents() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountAuth_GenerateAuthURL(t *testing.T) {
	type fields struct {
		provider   interfaces.Provider
		user       string
		pass       string
		privateKey string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "GenerateAuthURL OK",
			fields: fields{
				provider:   prov,
				user:       user,
				pass:       pass,
				privateKey: tests.PrivateKeyStub,
			},
			want:    wantToken,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AccountAuth{
				provider:   tt.fields.provider,
				user:       tt.fields.user,
				pass:       tt.fields.pass,
				privateKey: tt.fields.privateKey,
			}
			got, err := a.GenerateAuthURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountAuth.GenerateAuthURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccountAuth.GenerateAuthURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
