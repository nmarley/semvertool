package semvertool_test

import (
	"sort"
	"testing"

	"github.com/matryer/is"
	"github.com/nmarley/semvertool"
)

func TestSemVerInfo(t *testing.T) {
	tests := []struct {
		version        string
		shouldErr      bool
		major          string
		minor          string
		patch          string
		preRelease     string
		buildMetadata  string
		preReleaseHead string
		permutations   []string
	}{
		{
			version:        "3.4.0-dev.4+buildmeta1234",
			shouldErr:      false,
			major:          "3",
			minor:          "4",
			patch:          "0",
			preRelease:     "dev.4",
			buildMetadata:  "buildmeta1234",
			preReleaseHead: "dev",
			permutations:   []string{"3.4.0-dev", "3.4-dev", "3-dev"},
		},
		{
			version:   "v3.4.0",
			shouldErr: true,
		},
		{
			version:        "1.2.3",
			shouldErr:      false,
			major:          "1",
			minor:          "2",
			patch:          "3",
			preRelease:     "",
			buildMetadata:  "",
			preReleaseHead: "",
			permutations:   []string{"1.2.3", "1.2", "1"},
		},
		{
			version:        "2.4.5-beta",
			shouldErr:      false,
			major:          "2",
			minor:          "4",
			patch:          "5",
			preRelease:     "beta",
			buildMetadata:  "",
			preReleaseHead: "beta",
			permutations:   []string{"2.4.5-beta", "2.4-beta", "2-beta"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.version, func(st *testing.T) {
			is := is.New(st)
			info, err := semvertool.Parse(tt.version)
			if tt.shouldErr {
				is.True(err != nil)
				return
			}
			is.NoErr(err)
			is.Equal(info.Major, tt.major)
			is.Equal(info.Minor, tt.minor)
			is.Equal(info.Patch, tt.patch)
			is.Equal(info.PreRelease, tt.preRelease)
			is.Equal(info.BuildMetadata, tt.buildMetadata)
			is.Equal(info.PreReleaseHead(), tt.preReleaseHead)

			expected := tt.permutations
			got := info.Permutations()
			is.Equal(len(got), len(expected))
			sort.Strings(expected)
			sort.Strings(got)
			for i, str := range got {
				is.Equal(str, expected[i])
			}
		})
	}
}
