package stateless

import (
	"os"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

func loadPrivateKey() (jwk.Key, error) {
	b, err := os.ReadFile("../../../../private.pem")
	if err != nil {
		return nil, err
	}
	return jwk.ParseKey(b, jwk.WithX509(true))
}
