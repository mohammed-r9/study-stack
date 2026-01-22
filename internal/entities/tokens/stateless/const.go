package stateless

import "time"

const (
	ACCESS_TOKEN_TTL  = time.Minute * 20
	SERVICE_TOKEN_TTL = time.Minute * 2
)

type ServiceID string

const (
	SERVICE_PDF = ServiceID("service__pdf")
)
