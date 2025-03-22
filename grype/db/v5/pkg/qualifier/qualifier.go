package qualifier

import (
	"fmt"

	"DIDTrustCore/grype/pkg/qualifier"
)

type Qualifier interface {
	fmt.Stringer
	Parse() qualifier.Qualifier
}
