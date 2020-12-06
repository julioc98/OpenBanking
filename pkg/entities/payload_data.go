package entities

// JWTPayload ...
type JWTPayload struct {
	Aud          string `json:"aud"`
	Iss          string `json:"iss"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        int64  `json:"state"`
	Nonce        int64  `json:"nonce"`
	Exp          int64  `json:"exp"`
	ResponseType string `json:"response_type"`
	Claims       Claims `json:"claims"`
	Iat          int64  `json:"iat"`
}

// OpenbankingIntentID ...
type OpenbankingIntentID struct {
	Value     string `json:"value"`
	Essential bool   `json:"essential"`
}

// Userinfo ...
type Userinfo struct {
	OpenbankingIntentID OpenbankingIntentID `json:"openbanking_intent_id"`
}

// Acr ...
type Acr struct {
	Essential bool     `json:"essential"`
	Values    []string `json:"values"`
}

// IDToken ...
type IDToken struct {
	OpenbankingIntentID OpenbankingIntentID `json:"openbanking_intent_id"`
	Acr                 Acr                 `json:"acr"`
}

// Claims ...
type Claims struct {
	Userinfo Userinfo `json:"userinfo"`
	IDToken  IDToken  `json:"id_token"`
}
