package stateful

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
