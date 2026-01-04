package appErrors

type userError uint8

const (
	PasswordMismatch userError = iota
)

func (e userError) Error() string {
	switch e {
	case PasswordMismatch:
		return "Password Mismatch"
	}

	return "Unknown Error"
}
