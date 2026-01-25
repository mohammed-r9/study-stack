package stateless

import (
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/env"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func NewAcessToken(data UserClaims) (string, error) {
	token := jwt.New()

	token.Set(jwt.SubjectKey, data.UserID.String())
	token.Set(jwt.IssuedAtKey, time.Now())
	token.Set(jwt.ExpirationKey, time.Now().Add(ACCESS_TOKEN_TTL))

	token.Set("is_verified", data.IsVerified)

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), env.Config.JWT_KEY))
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

// returns an error if the token is invalid and the UserClaims otherwise
func VerifyAccessToken(tokenStr string) (UserClaims, error) {
	parsed, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKey(jwa.HS256(), env.Config.JWT_KEY),
	)

	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	err = jwt.Validate(parsed)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	sub, _ := parsed.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	var isVerified bool
	err = parsed.Get("is_verified", &isVerified)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	return UserClaims{
		UserID:     userID,
		IsVerified: isVerified,
	}, nil
}
