package consts

import "time"

const (
	TTL_REFRESH        = time.Hour * 24 * 15
	TTL_CSRF           = time.Hour * 24 * 15
	UserDataContextKey = "user-data"
	// pagination page size
	PAGE_SIZE = 20

	// RateLimit
	RL_AUTH    = "authed_user"
	RL_REFRESH = "refresh"
	RL_PUBLIC  = "public"
)
