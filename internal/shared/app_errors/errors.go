package appErrors

type userError uint8

const (
	PasswordMismatch userError = iota
	InvaliedAccessToken
)

func (e userError) Error() string {
	switch e {
	case PasswordMismatch:
		return "Password Mismatch"
	case InvaliedAccessToken:
		return "Invalid jwt"
	}

	return "Unknown Error"
}
