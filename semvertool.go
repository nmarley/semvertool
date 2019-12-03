package semvertool

import (
	"fmt"
	"regexp"
)

// Regexp from here:
// https://github.com/semver/semver/issues/232#issuecomment-405596809
var reSemVer = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
var rePreReleaseHead = regexp.MustCompile(`^([A-Za-z0-9]+)(?:\.)*.*$`)

// main lib + tests:
// - valid semver
// - permutation strings
// - prerelease head

// SemVerInfo contains the various fields defined in SemVer.
type SemVerInfo struct {
	Major         string
	Minor         string
	Patch         string
	PreRelease    string
	BuildMetadata string
}

// Parse parses string into semantic version info, and returns an error
// if invalid SemVer.
func Parse(version string) (*SemVerInfo, error) {
	if !reSemVer.MatchString(version) {
		return nil, fmt.Errorf("%s is not valid SemVer", version)
	}

	// parse version into map
	res := reSemVer.FindStringSubmatch(version)
	vMap := make(map[string]string)
	for i, name := range reSemVer.SubexpNames() {
		vMap[name] = res[i]
	}
	return &SemVerInfo{
		Major:         vMap["major"],
		Minor:         vMap["minor"],
		Patch:         vMap["patch"],
		PreRelease:    vMap["prerelease"],
		BuildMetadata: vMap["buildmetadata"],
	}, nil
}

// PreReleaseHead returns the first part of prerelease only -- the part before
// any `.` char. This is useful for prerelease versions like `dev.10` to return
// only the `dev` part.
func (i *SemVerInfo) PreReleaseHead() string {
	res := rePreReleaseHead.FindStringSubmatch(i.PreRelease)
	if len(res) == 2 && len(res[1]) > 0 {
		return res[1]
	}
	// unable to parse head or not found
	return i.PreRelease
}

// Permutations returns possible string permutations of Major / Minor / Patch
// with -PreReleaseHead if it exists. E.g. a version of "1.2.3" will return
// "1", "1.2" and "1.2.3".
//
// The main reason the `-show-permutations` flag exists is to use for Docker
// tagging, as they are commonly applied to Docker images. Tags like like
// `latest` are not included in this tool and should be considered separately,
// as this only deals with version numbers.
func (i *SemVerInfo) Permutations() []string {
	appendPreRelease := func(ver, pre string) string {
		if pre != "" {
			return fmt.Sprintf("%s-%s", ver, pre)
		}
		return ver
	}

	var permutations [3]string
	preReleaseHead := i.PreReleaseHead()

	majMin := fmt.Sprintf("%s.%s", i.Major, i.Minor)
	majMinPatch := fmt.Sprintf("%s.%s", majMin, i.Patch)

	permutations[0] = appendPreRelease(i.Major, preReleaseHead)
	permutations[1] = appendPreRelease(majMin, preReleaseHead)
	permutations[2] = appendPreRelease(majMinPatch, preReleaseHead)

	return permutations[:]
}
