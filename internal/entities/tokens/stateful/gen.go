package stateful

import "github.com/google/uuid"

func NewRefreshToken() (refreshToken, error) {
	refreshPlain, err := newOpaque()
	if err != nil {
		return refreshToken{}, err
	}
	csrfPlain, err := newOpaque()
	if err != nil {
		return refreshToken{}, err
	}

	return refreshToken{
		PlainText:     refreshPlain,
		CsrfPlainText: csrfPlain,
		Hash:          HashFromPlainText(refreshPlain),
		CsrfHash:      HashFromPlainText(csrfPlain),
	}, nil
}

func NewOpaqueToken(userID uuid.UUID, scope OpaqueScope) (OpaqueToken, error) {
	plainText, err := newOpaque()
	if err != nil {
		return OpaqueToken{}, err
	}

	return OpaqueToken{
		UserID:    userID,
		Hash:      HashFromPlainText(plainText),
		PlainText: plainText,
		Scope:     scope,
	}, nil
}
