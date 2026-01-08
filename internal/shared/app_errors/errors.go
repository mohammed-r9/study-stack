package appErrors

type userError uint8

const (
	PasswordMismatch userError = iota
	InvaliedAccessToken
	InvalidRefreshToken
	InvalidCsrfToken
	NoChange
	InvalidUserID
	InvalidEmail
	InvalidVerificationToken
)

func (e userError) Error() string {
	switch e {
	case PasswordMismatch:
		return "Password Mismatch"
	case InvaliedAccessToken:
		return "Invalid jwt"
	case InvalidRefreshToken:
		return "Invalid refresh token"
	case InvalidCsrfToken:
		return "Invalid csrf token"
	case NoChange:
		return "New field value is the same"
	case InvalidUserID:
		return "Invalid user ID"
	case InvalidEmail:
		return "Invalid email"
	case InvalidVerificationToken:
		return "Invalid email verifcation token"
	}

	return "Unknown Error"
}
