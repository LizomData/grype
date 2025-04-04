package java

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"DIDTrustCore/grype/match"
	"DIDTrustCore/grype/pkg"
	"DIDTrustCore/internal/stringutil"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

func TestMatcherJava_matchUpstreamMavenPackage(t *testing.T) {
	newMatcher := func(searcher MavenSearcher) *Matcher {
		return &Matcher{
			cfg: MatcherConfig{
				ExternalSearchConfig: ExternalSearchConfig{
					SearchMavenUpstream: true,
				},
			},
			MavenSearcher: searcher,
		}
	}
	store := newMockProvider()

	p := pkg.Package{
		ID:       pkg.ID(uuid.NewString()),
		Name:     "org.springframework.spring-webmvc",
		Version:  "5.1.5.RELEASE",
		Language: syftPkg.Java,
		Type:     syftPkg.JavaPkg,
		Metadata: pkg.JavaMetadata{
			ArchiveDigests: []pkg.Digest{
				{
					Algorithm: "sha1",
					Value:     "236e3bfdbdc6c86629237a74f0f11414adb4e211",
				},
			},
		},
	}

	t.Run("matching from maven search results", func(t *testing.T) {
		matcher := newMatcher(mockMavenSearcher{
			pkg: p,
		})
		actual, _ := matcher.matchUpstreamMavenPackages(store, p)

		assert.Len(t, actual, 2, "unexpected matches count")

		foundCVEs := stringutil.NewStringSet()
		for _, v := range actual {
			foundCVEs.Add(v.Vulnerability.ID)

			require.NotEmpty(t, v.Details)
			for _, d := range v.Details {
				assert.Equal(t, match.ExactIndirectMatch, d.Type, "indirect match not indicated")
				assert.Equal(t, matcher.Type(), d.Matcher, "failed to capture matcher type")
			}
			assert.Equal(t, p.Name, v.Package.Name, "failed to capture original package name")
		}

		for _, id := range []string{"CVE-2014-fake-2", "CVE-2013-fake-3"} {
			if !foundCVEs.Contains(id) {
				t.Errorf("missing discovered CVE: %s", id)
			}
		}
		if t.Failed() {
			t.Logf("discovered CVES: %+v", foundCVEs)
		}
	})

	t.Run("handles maven rate limiting", func(t *testing.T) {
		matcher := newMatcher(mockMavenSearcher{simulateRateLimiting: true})

		_, err := matcher.matchUpstreamMavenPackages(store, p)

		assert.Errorf(t, err, "should have gotten an error from the rate limiting")
	})
}
