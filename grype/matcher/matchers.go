package matcher

import (
	"DIDTrustCore/grype/match"
	"DIDTrustCore/grype/matcher/apk"
	"DIDTrustCore/grype/matcher/dotnet"
	"DIDTrustCore/grype/matcher/dpkg"
	"DIDTrustCore/grype/matcher/golang"
	"DIDTrustCore/grype/matcher/java"
	"DIDTrustCore/grype/matcher/javascript"
	"DIDTrustCore/grype/matcher/msrc"
	"DIDTrustCore/grype/matcher/portage"
	"DIDTrustCore/grype/matcher/python"
	"DIDTrustCore/grype/matcher/rpm"
	"DIDTrustCore/grype/matcher/ruby"
	"DIDTrustCore/grype/matcher/rust"
	"DIDTrustCore/grype/matcher/stock"
)

// Config contains values used by individual matcher structs for advanced configuration
type Config struct {
	Java       java.MatcherConfig
	Ruby       ruby.MatcherConfig
	Python     python.MatcherConfig
	Dotnet     dotnet.MatcherConfig
	Javascript javascript.MatcherConfig
	Golang     golang.MatcherConfig
	Rust       rust.MatcherConfig
	Stock      stock.MatcherConfig
}

func NewDefaultMatchers(mc Config) []match.Matcher {
	return []match.Matcher{
		&dpkg.Matcher{},
		ruby.NewRubyMatcher(mc.Ruby),
		python.NewPythonMatcher(mc.Python),
		dotnet.NewDotnetMatcher(mc.Dotnet),
		&rpm.Matcher{},
		java.NewJavaMatcher(mc.Java),
		javascript.NewJavascriptMatcher(mc.Javascript),
		&apk.Matcher{},
		golang.NewGolangMatcher(mc.Golang),
		&msrc.Matcher{},
		&portage.Matcher{},
		rust.NewRustMatcher(mc.Rust),
		stock.NewStockMatcher(mc.Stock),
	}
}
