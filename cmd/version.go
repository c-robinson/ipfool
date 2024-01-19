package cmd

import (
	"fmt"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var Version = ""

type Stamp struct {
	InfoGoVersion string
	InfoCompiler  string
	InfoGOARCH    string
	InfoGOOS      string
	InfoBuildTime string
	VCSRevision   string
}

var versionExtendedFlag bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print ipfool version",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			panic("could not read build info")
		}
		iplibVersion := "unknown"
		for _, mod := range info.Deps {
			if mod.Path == "github.com/c-robinson/iplib/v2" {
				iplibVersion = mod.Version
			}
		}

		fmt.Printf("ipfool %s, using iplib %s\n", retrieveAppVersion(info), iplibVersion)
		if versionExtendedFlag {
			viewExtendedVersion(info)
		}
	},
}

func retrieveAppVersion(info *debug.BuildInfo) string {
	if Version != "" {
		return Version
	}
	if info.Main.Version != "" {
		return info.Main.Version
	}
	return "(unknown)"
}

func retrieveDepends(info *debug.BuildInfo) []string {
	var name, ver string

	Depends := []string{}

	for _, module := range info.Deps {
		path := strings.Split(module.Path, "/")
		if len(path) < 3 {
			return Depends
		}
		name = path[2]
		if strings.Contains("golang.org", path[0]) {
			return Depends
		}

		if len(module.Version) == 0 {
			ver = module.Sum
		} else {
			ver = module.Version
		}
		Depends = append(Depends, fmt.Sprintf("%s %s", name, ver))
	}

	sort.Strings(Depends)
	return Depends
}

func retrieveStamp(info *debug.BuildInfo) *Stamp {
	stamp := Stamp{}
	for _, setting := range info.Settings {
		switch setting.Key {
		case "-compiler":
			stamp.InfoCompiler = setting.Value

		case "GOARCH":
			stamp.InfoGOARCH = setting.Value

		case "GOOS":
			stamp.InfoGOOS = setting.Value

		case "vcs.time":
			stamp.InfoBuildTime = setting.Value

		case "vcs.revision":
			stamp.VCSRevision = setting.Value
		}
	}

	return &stamp
}

func viewExtendedVersion(info *debug.BuildInfo) {
	stamp := retrieveStamp(info)
	depends := retrieveDepends(info)

	fmt.Printf("  Built with %s on %s\n", stamp.InfoCompiler, stamp.InfoBuildTime)
	fmt.Printf("  VCS revision: %s\n", stamp.VCSRevision)
	fmt.Printf("  Go version %s, GOOS %s, GOARCH %s\n", info.GoVersion, stamp.InfoGOOS, stamp.InfoGOARCH)
	fmt.Printf("  Dependencies:\n")
	for _, mod := range depends {
		fmt.Printf("    %s\n", mod)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVar(&versionExtendedFlag, "extended", false, "get expanded version information")
}
