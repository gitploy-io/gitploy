package vo

import "time"

const (
	TrialMemberLimit = 5
)

const (
	LicenseKindTrial    LicenseKind = "trial"
	LicenseKindStandard LicenseKind = "standard"
)

type (
	LicenseKind string

	License struct {
		Kind        LicenseKind `json:"kind"`
		MemberCount int         `json:"member_count"`
		MemberLimit int         `json:"memeber_limit,omitemtpy"`
		ExpiredAt   time.Time   `json:"expired_at,omitemtpy"`
	}

	// SigningData marshal and unmarshal the content of license.
	SigningData struct {
		MemberLimit int       `json:"memeber_limit"`
		ExpiredAt   time.Time `json:"expired_at"`
	}
)

func NewTrialLicense(cnt int) *License {
	return &License{
		Kind:        LicenseKindTrial,
		MemberCount: cnt,
	}
}

func NewStandardLicense(cnt int, d *SigningData) *License {
	return &License{
		Kind:        LicenseKindStandard,
		MemberCount: cnt,
		MemberLimit: d.MemberLimit,
		ExpiredAt:   d.ExpiredAt,
	}
}

func (l *License) IsTrial() bool {
	return l.Kind == LicenseKindTrial
}

func (l *License) IsOverLimit() bool {
	return l.MemberCount > TrialMemberLimit
}

func (l *License) IsExpired() bool {
	return l.ExpiredAt.Before(time.Now())
}
