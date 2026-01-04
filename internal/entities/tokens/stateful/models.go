package stateful

type refreshToken struct {
	PlainText     string
	Hash          string
	CsrfPlainText string
	CsrfHash      string
}
