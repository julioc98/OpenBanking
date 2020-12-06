package jwt

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

// JWT struct
type JWT struct {
	privateKey []byte
}

// NewJWT factory
func NewJWT(privk string) *JWT {
	return &JWT{
		privateKey: []byte(privk),
	}
}

// Generate a new JWT token
func (j *JWT) Generate(jsonClaims []byte) (string, error) {

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(jsonClaims, &jsonMap)
	if err != nil {
		return "", err
	}

	c := jws.Claims(jsonMap)

	jwtJose := jws.NewJWT(c, crypto.SigningMethodRS256)

	block2, _ := pem.Decode(j.privateKey)

	rsaPriv, err := x509.ParsePKCS8PrivateKey(block2.Bytes)
	if err != nil {
		return "", err
	}

	token, err := jwtJose.Serialize(rsaPriv)
	if err != nil {
		return "", err
	}

	return string(token), nil
}
