package licenser

import (
	"github.com/google/uuid"
)

// GenerateOSSLicense generates Active License of type OSS
func GenerateOSSLicense() License {
	return License{
		UUID:   uuid.New().String(),
		Type:   OSS,
		Status: Active,
	}
}
