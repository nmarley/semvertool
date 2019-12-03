package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nmarley/semvertool"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "semvertool <version-string>",
	Short: "SemVerTool is used to verify SemVer and print versions",
	Long:  "",
	Run:   goGoGadgetSemVer,
}

// flags
var quiet bool
var showPermutations bool
var preRelease bool
var preReleaseHead bool

func init() {
	mainCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "do not show output")
	mainCmd.PersistentFlags().BoolVarP(&showPermutations, "show-permutations", "s", false, "print all major / minor / patch permutations as described in README")
	mainCmd.PersistentFlags().BoolVarP(&preRelease, "prerelease", "r", false, "print prerelease info")
	mainCmd.PersistentFlags().BoolVar(&preReleaseHead, "prerelease-head", false, "print prerelease HEAD - the first part only")
}

// Possible / ideas for other flags
// -major
// -minor
// -patch
// -buildmetadata

func goGoGadgetSemVer(c *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <version-string>\n", os.Args[0])
		os.Exit(1)
	}

	info, err := semvertool.Parse(args[0])
	if err != nil {
		if !quiet {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	// all further options are display only
	if quiet {
		return
	}

	if preRelease {
		fmt.Println(info.PreRelease)
	}

	if preReleaseHead {
		fmt.Println(info.PreReleaseHead())
	}

	if showPermutations {
		fmt.Println(strings.Join(info.Permutations(), " "))
	}
}

func main() {
	mainCmd.Execute()
}
