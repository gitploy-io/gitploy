package extent

import (
	"math"
	"time"
)

const (
	TrialMemberLimit     = 5
	TrialDeploymentLimit = 5000
	InfiniteMemberLimit  = math.MaxInt
	InfiniteDeploymentLimit = math.MaxInt
)

const (
	// LicenseKindOSS is a license for the community edition.
	LicenseKindOSS LicenseKind = "oss"
	// LicenseKindTrial is a trial license of the enterprise edition.
	LicenseKindTrial LicenseKind = "trial"
	// LicenseKindStandard is a license of the enterprise edition.
	LicenseKindStandard LicenseKind = "standard"
)

type (
	LicenseKind string

	License struct {
		Kind            LicenseKind `json:"kind"`
		MemberCount     int         `json:"member_count"`
		MemberLimit     int         `json:"memeber_limit"`
		DeploymentCount int         `json:"deployment_count"`
		DeploymentLimit int         `json:"deployment_limit"`
		ExpiredAt       time.Time   `json:"expired_at"`
	}

	// SigningData marshal and unmarshal the content of license.
	SigningData struct {
		MemberLimit int       `json:"memeber_limit"`
		ExpiredAt   time.Time `json:"expired_at"`
	}
)

func NewOSSLicense() *License {
	return &License{
		Kind:            LicenseKindOSS,
		MemberCount:     -1,
		DeploymentCount: -1,
	}
}

func NewTrialLicense(memberCnt, deploymentCnt int) *License {
	return &License{
		Kind:            LicenseKindTrial,
		MemberCount:     memberCnt,
		MemberLimit:     InfiniteMemberLimit,
		DeploymentCount: deploymentCnt,
		DeploymentLimit: InfiniteMemberLimit,
	}
}

func NewStandardLicense(memberCnt int, d *SigningData) *License {
	return &License{
		Kind:            LicenseKindStandard,
		MemberCount:     memberCnt,
		MemberLimit:     InfiniteMemberLimit,
		DeploymentCount: -1,
		ExpiredAt:       d.ExpiredAt,
	}
}

func (l *License) IsOSS() bool {
	return l.Kind == LicenseKindOSS
}

func (l *License) IsTrial() bool {
	return l.Kind == LicenseKindTrial
}

func (l *License) IsStandard() bool {
	return l.Kind == LicenseKindStandard
}

// IsOverLimit verify it is over the limit of the license.
func (l *License) IsOverLimit() bool {
	return l.MemberCount > l.MemberLimit || l.DeploymentCount > l.DeploymentLimit
}

// IsExpired verify that the license is expired or not.
func (l *License) IsExpired() bool {
	return l.ExpiredAt.Before(time.Now())
}
