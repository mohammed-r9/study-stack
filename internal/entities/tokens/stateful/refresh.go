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
		Hash:          hashFromPlainText(refreshPlain),
		CsrfHash:      hashFromPlainText(csrfPlain),
	}, nil
}
