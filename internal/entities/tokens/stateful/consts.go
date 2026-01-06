package stateful

import "time"

type OpaqueScope string

const (
	PasswordReset     = OpaqueScope("Password_reset")
	EmailVerification = OpaqueScope("email_verification")
)

const (
	PasswordResetTTL     = 30 * time.Minute
	EmailVerificationTTL = 4 * time.Hour
)
