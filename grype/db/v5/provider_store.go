package v5

import (
	"DIDTrustCore/grype/match"
	"DIDTrustCore/grype/vulnerability"
)

type ProviderStore struct {
	vulnerability.Provider
	match.ExclusionProvider
}
