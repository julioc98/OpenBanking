package jwt

import (
	"testing"

	"github.com/julioc98/openbanking/tests"
)

var (
	privateKeyMock = []byte(tests.PrivateKeyStub)
	jsonClaimsMock = []byte(`{
            "aud": "https://auth.obiebank.banfico.com/auth/realms/provider",
            "iss": "PSDBR-NCA-AISP01",
            "client_id": "PSDBR-NCA-AISP01",
            "redirect_uri": "https://openbankinghacka.herokuapp.com/callback",
            "scope": "openid profile email accounts",
            "state": 1125107479227,
            "nonce": 1400476011498,
            "exp": 1607402832,
            "response_type": "code id_token",
            "claims": {
            "userinfo": {
                "openbanking_intent_id": {
                "value": "urn:obiebank:accounts:582c0c25-dfc4-4fda-8683-c73e08dabb63",
                "essential": true
                }
            },
            "id_token": {
                "openbanking_intent_id": {
                "value": "urn:obiebank:accounts:582c0c25-dfc4-4fda-8683-c73e08dabb63",
                "essential": true
                },
                "acr": {
                "essential": true,
                "values": [
                    "urn:openbanking:psd2:sca",
                    "urn:openbanking:psd2:ca"
                ]
                }
            }
            },
            "iat": 1607226432
        }`)
	wantToken = "eyJhbGciOiJSUzI1NiJ9.eyJhdWQiOiJodHRwczovL2F1dGgub2JpZWJhbmsuYmFuZmljby5jb20vYXV0aC9yZWFsbXMvcHJvdmlkZXIiLCJjbGFpbXMiOnsiaWRfdG9rZW4iOnsiYWNyIjp7ImVzc2VudGlhbCI6dHJ1ZSwidmFsdWVzIjpbInVybjpvcGVuYmFua2luZzpwc2QyOnNjYSIsInVybjpvcGVuYmFua2luZzpwc2QyOmNhIl19LCJvcGVuYmFua2luZ19pbnRlbnRfaWQiOnsiZXNzZW50aWFsIjp0cnVlLCJ2YWx1ZSI6InVybjpvYmllYmFuazphY2NvdW50czo1ODJjMGMyNS1kZmM0LTRmZGEtODY4My1jNzNlMDhkYWJiNjMifX0sInVzZXJpbmZvIjp7Im9wZW5iYW5raW5nX2ludGVudF9pZCI6eyJlc3NlbnRpYWwiOnRydWUsInZhbHVlIjoidXJuOm9iaWViYW5rOmFjY291bnRzOjU4MmMwYzI1LWRmYzQtNGZkYS04NjgzLWM3M2UwOGRhYmI2MyJ9fX0sImNsaWVudF9pZCI6IlBTREJSLU5DQS1BSVNQODciLCJleHAiOjE2MDc0MDI4MzIsImlhdCI6MTYwNzIyNjQzMiwiaXNzIjoiUFNEQlItTkNBLUFJU1A4NyIsIm5vbmNlIjoxNDAwNDc2MDExNDk4LCJyZWRpcmVjdF91cmkiOiJodHRwczovL3Bpc21vLWFwaS5oZXJva3VhcHAuY29tLy9jYWxsYmFjayIsInJlc3BvbnNlX3R5cGUiOiJjb2RlIGlkX3Rva2VuIiwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBhY2NvdW50cyIsInN0YXRlIjoxMTI1MTA3NDc5MjI3fQ.Z2L2F1obEnf_gJ40SGqUdsVDo-F53GHl9gwmwVXNckJ-JlMuNZKgnrey6jArxwhqFdBqMknP3oFO7CJRYqnrygBEjK0185eY9FWKOPnUvZFsofU1fEjH4AKnDXKgGwDJW4ASe3NNZ1ZkaKM5cuU5B6xvIfkqzmbPCvFd0s-3PTSFNmc9j6J12s_GHSV0S8L63E-px4Hp94BXHVhQ9bCRdT5oFCG-i_H56CpPqvMEgQv3-fba6cenfBqRsJ9Xl26VcByxGjRVQKpu8vgBFewHK5W06_WUsQBu73gOezATbIY8o-EGS_PX3q7Ma5or0HfoUtFIw7565oA_nTSMZtLPtQ"
)

func TestJWT_Generate(t *testing.T) {
	type fields struct {
		privateKey []byte
	}
	type args struct {
		jsonClaims []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Generate OK",
			fields: fields{
				privateKey: privateKeyMock,
			},
			args: args{
				jsonClaims: jsonClaimsMock,
			},
			want:    wantToken,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWT{
				privateKey: tt.fields.privateKey,
			}
			got, err := j.Generate(tt.args.jsonClaims)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWT.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JWT.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
