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

// database errors

type DBError uint8

const (
	NoRowsAffected DBError = iota
	InvalidData
)

func (e DBError) Error() string {
	switch e {
	case NoRowsAffected:
		return "DB: Now rows affected"
	case InvalidData:
		return "Invalid data"
	}

	return "Unknown Error"
}

type ServiceError uint8

const (
	// General / system errors
	InternalError ServiceError = iota
	NotImplemented
	ServiceUnavailable

	// Client / user errors
	BadData
	Unauthorized
	Forbidden
	NotFound
	MethodNotAllowed
	Conflict
	TooManyRequests
	ValidationFailed
	InvalidCredentials
	FileTooLarge
)

func (e ServiceError) Error() string {
	switch e {
	case InternalError:
		return "service: internal error"
	case NotImplemented:
		return "service: not implemented"
	case ServiceUnavailable:
		return "service: unavailable"

	case BadData:
		return "service: bad request"
	case Unauthorized:
		return "service: unauthorized"
	case Forbidden:
		return "service: forbidden"
	case NotFound:
		return "service: not found"
	case MethodNotAllowed:
		return "service: method not allowed"
	case Conflict:
		return "service: conflict"
	case TooManyRequests:
		return "service: too many requests"
	case ValidationFailed:
		return "service: validation failed"
	case InvalidCredentials:
		return "service: invalid credentials"
	case FileTooLarge:
		return "service: file too large"

	default:
		return "service: unknown error"
	}
}
