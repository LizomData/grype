package platformcpe

import (
	"fmt"

	"DIDTrustCore/grype/pkg/qualifier"
	"DIDTrustCore/grype/pkg/qualifier/platformcpe"
)

type Qualifier struct {
	Kind string `json:"kind" mapstructure:"kind"`                   // Kind of qualifier
	CPE  string `json:"cpe,omitempty" mapstructure:"cpe,omitempty"` // CPE
}

func (q Qualifier) Parse() qualifier.Qualifier {
	return platformcpe.New(q.CPE)
}

func (q Qualifier) String() string {
	return fmt.Sprintf("kind: %s, cpe: %q", q.Kind, q.CPE)
}
