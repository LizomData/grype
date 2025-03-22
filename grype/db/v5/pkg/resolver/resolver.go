package resolver

import (
	"DIDTrustCore/grype/db/v5/pkg/resolver/java"
	"DIDTrustCore/grype/db/v5/pkg/resolver/python"
	"DIDTrustCore/grype/db/v5/pkg/resolver/stock"
	grypePkg "DIDTrustCore/grype/pkg"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Resolver interface {
	Normalize(string) string
	Resolve(p grypePkg.Package) []string
}

func FromLanguage(language syftPkg.Language) (Resolver, error) {
	var r Resolver

	switch language {
	case syftPkg.Python:
		r = &python.Resolver{}
	case syftPkg.Java:
		r = &java.Resolver{}
	default:
		r = &stock.Resolver{}
	}

	return r, nil
}

func PackageNames(p grypePkg.Package) []string {
	names := []string{p.Name}
	r, _ := FromLanguage(p.Language)
	if r != nil {
		parts := r.Resolve(p)
		if len(parts) > 0 {
			names = parts
		}
	}
	return names
}
