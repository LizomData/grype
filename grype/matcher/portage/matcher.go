package portage

import (
	"DIDTrustCore/grype/match"
	"DIDTrustCore/grype/matcher/internal"
	"DIDTrustCore/grype/pkg"
	"DIDTrustCore/grype/vulnerability"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Matcher struct {
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.PortagePkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.PortageMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, p pkg.Package) ([]match.Match, []match.IgnoredMatch, error) {
	return internal.MatchPackageByDistro(store, p, m.Type())
}
