package extent

import "time"

type (
	RateLimit struct {
		Limit     int       `json:"limit"`
		Remaining int       `json:"remaining"`
		Reset     time.Time `json:"reset"`
	}
)
