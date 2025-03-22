package qualifier

import (
	"DIDTrustCore/grype/pkg"
)

type Qualifier interface {
	Satisfied(p pkg.Package) (bool, error)
}
