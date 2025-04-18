package ruby

import (
	"DIDTrustCore/grype/match"
	"DIDTrustCore/grype/matcher/internal"
	"DIDTrustCore/grype/pkg"
	"DIDTrustCore/grype/vulnerability"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Matcher struct {
	cfg MatcherConfig
}

type MatcherConfig struct {
	UseCPEs bool
}

func NewRubyMatcher(cfg MatcherConfig) *Matcher {
	return &Matcher{
		cfg: cfg,
	}
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.GemPkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.RubyGemMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, p pkg.Package) ([]match.Match, []match.IgnoredMatch, error) {
	return internal.MatchPackageByEcosystemAndCPEs(store, p, m.Type(), m.cfg.UseCPEs)
}
