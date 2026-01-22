package stateless

import (
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

// returns an asymetrically signed JWT with a serviceID when success,
// verification should happen on the reciever service using a public key
func NewServiceToken(id ServiceID) (string, error) {
	token := jwt.New()

	key, err := loadPrivateKey()
	if err != nil {
		return "", err
	}
	token.Set(jwt.SubjectKey, string(id))
	token.Set(jwt.IssuedAtKey, time.Now())
	token.Set(jwt.ExpirationKey, time.Now().Add(SERVICE_TOKEN_TTL))

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256(), key))
	if err != nil {
		return "", err
	}

	return string(signed), nil
}
