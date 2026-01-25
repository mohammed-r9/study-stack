package stateful

import "github.com/google/uuid"

type refreshToken struct {
	PlainText     string
	Hash          string
	CsrfPlainText string
	CsrfHash      string
}

type OpaqueToken struct {
	PlainText string
	Hash      string
	UserID    uuid.UUID
	Scope     OpaqueScope
}
