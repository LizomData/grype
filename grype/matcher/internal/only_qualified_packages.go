package internal

import (
	"fmt"

	"DIDTrustCore/grype/pkg"
	"DIDTrustCore/grype/search"
	"DIDTrustCore/grype/vulnerability"
)

// onlyQualifiedPackages returns a criteria object that tests vulnerability qualifiers against the provided package
func onlyQualifiedPackages(p pkg.Package) vulnerability.Criteria {
	return search.ByFunc(func(vuln vulnerability.Vulnerability) (bool, string, error) {
		for _, qualifier := range vuln.PackageQualifiers {
			satisfied, err := qualifier.Satisfied(p)
			if err != nil {
				return satisfied, fmt.Sprintf("unable to evaluate qualifier: %s", err.Error()), err
			}
			if !satisfied {
				// TODO: qualifiers don't have a good string representation
				return false, fmt.Sprintf("package does not satisfy qualifier: %#v", qualifier), nil
			}
		}
		return true, "", nil // all qualifiers passed
	})
}
