package gateways

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/julioc98/openbanking/pkg/entities"
	"github.com/julioc98/openbanking/pkg/interfaces"
)

func TestObiebank_GetToken(t *testing.T) {
	type fields struct {
		authURL string
		baseURL string
		client  interfaces.Requester
	}
	type args struct {
		user string
		pass string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Token
		wantErr bool
	}{
		{
			name: "Request OK",
			fields: fields{
				authURL: "https://auth.obiebank.banfico.com",
				baseURL: "",
				client:  http.DefaultClient,
			},
			args: args{
				user: "PSDBR-NCA-AISP01",
				pass: "xpto",
			},
			want:    &entities.Token{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Obiebank{
				authURL: tt.fields.authURL,
				baseURL: tt.fields.baseURL,
				client:  tt.fields.client,
			}
			got, err := g.GetToken(tt.args.user, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("Obiebank.GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Obiebank.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObiebank_CreateAccountAccessConsents(t *testing.T) {
	type fields struct {
		authURL string
		baseURL string
		client  interfaces.Requester
	}
	type args struct {
		accountAccessConsents *entities.AccountAccessConsents
		token                 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.AccountAccessConsents
		wantErr bool
	}{
		{
			name: "Request OK",
			fields: fields{
				authURL: "",
				baseURL: "https://gw-dev.obiebank.banfico.com/obie-aisp/v3.1/aisp",
				client:  http.DefaultClient,
			},
			args: args{
				accountAccessConsents: &entities.AccountAccessConsents{
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
				},
				token: "jwt_token",
			},
			want:    &entities.AccountAccessConsents{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Obiebank{
				authURL: tt.fields.authURL,
				baseURL: tt.fields.baseURL,
				client:  tt.fields.client,
			}
			got, err := g.CreateAccountAccessConsents(tt.args.accountAccessConsents, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Obiebank.CreateAccountAccessConsents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Obiebank.CreateAccountAccessConsents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObiebank_GetTokenByCode(t *testing.T) {
	type fields struct {
		authURL     string
		baseURL     string
		redirectURI string
		client      interfaces.Requester
	}
	type args struct {
		user string
		pass string
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Token
		wantErr bool
	}{
		{
			name: "Request OK",
			fields: fields{
				authURL:     "https://auth.obiebank.banfico.com",
				baseURL:     "https://gw-dev.obiebank.banfico.com/obie-aisp/v3.1/aisp",
				client:      http.DefaultClient,
				redirectURI: "https://openbankinghacka.herokuapp.com/callback",
			},
			args: args{
				user: "PSDBR-NCA-AISP77",
				pass: "senha123",
				code: "138efc91-15cf-479a-af81-4720cb7f4b6f.i2qmM6i8CP8.67456a91-3f76-4170-b5f2-b08dce488abe",
			},
			want:    &entities.Token{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Obiebank{
				authURL:     tt.fields.authURL,
				baseURL:     tt.fields.baseURL,
				redirectURI: tt.fields.redirectURI,
				client:      tt.fields.client,
			}
			got, err := g.GetTokenByCode(tt.args.user, tt.args.pass, tt.args.code)
			if (err != nil) != tt.wantErr {

				t.Errorf("Obiebank.GetTokenByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Obiebank.GetTokenByCode() = %v", got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Obiebank.GetTokenByCode() = %v, want %v", got, tt.want)
			// }
		})
	}
}
