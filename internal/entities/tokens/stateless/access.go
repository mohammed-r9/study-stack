package stateless

import (
	"strconv"
	appErrors "study-stack/internal/shared/app_errors"
	"study-stack/internal/shared/env"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func NewAcessToken(data UserClaims) (string, error) {
	token := jwt.New()

	token.Set(jwt.SubjectKey, data.UserID)
	token.Set(jwt.IssuedAtKey, time.Now())
	token.Set(jwt.ExpirationKey, ACCESS_TOKEN_TTL)

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

	if _, expired := parsed.Expiration(); expired {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	var userIDStr string
	var isVerifiedStr string
	err = parsed.Get("user_id", userIDStr)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	err = parsed.Get("is_verified", userIDStr)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}

	isVerified, err := strconv.ParseBool(isVerifiedStr)
	if err != nil {
		return UserClaims{}, appErrors.InvaliedAccessToken
	}
	return UserClaims{
		UserID:     userID,
		IsVerified: isVerified,
	}, nil
}
