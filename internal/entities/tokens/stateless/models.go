package stateless

import "github.com/google/uuid"

type UserClaims struct {
	UserID     uuid.UUID
	IsVerified bool
}
