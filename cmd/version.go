package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionExtendedFlag bool

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "print ipfool version",
	Run: func(cmd *cobra.Command, args []string) {
		info, _ := debug.ReadBuildInfo()
		iplibVersion := "unknown"
		for _, mod := range info.Deps {
			if mod.Path == "github.com/c-robinson/iplib/v2" {
				iplibVersion = mod.Version
			}
		}

		fmt.Printf("ipfool version %s (using iplib %s)\n", info.Main.Version, iplibVersion)
		if versionExtendedFlag {
			getExtendedVersion(info)
		}
	},
}

func getExtendedVersion(info *debug.BuildInfo) {

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
			stamp.GameRevision = setting.Value
		}
	}

	fmt.Printf("Built on %s using %s\n", info, info.Main.Path)
	for _, mod := range info.Deps {
		fmt.Printf("%s %s\n", mod.Path, mod.Version)
	}
}

func init() {
	rootCmd.AddCommand(VersionCmd)
	VersionCmd.Flags().BoolVar(&versionExtendedFlag, "extended", false, "get expanded version information")
}

/*
func RetrieveStamp() *Stamp {
	info, _ := debug.ReadBuildInfo()

	stamp := Stamp{}
	stamp.Dependencies = []Dependency{}
	for _, mod := range info.Deps {
		dep, ok := retrieveDepends(mod)
		if ok {
			stamp.Dependencies = append(stamp.Dependencies, dep)
		}
	}
	sort.Slice(stamp.Dependencies, func(i, j int) bool {
		return stamp.Dependencies[i].Name < stamp.Dependencies[j].Name
	})

	stamp.InfoGoVersion = info.GoVersion
	stamp.GameVersion = info.Main.Version
	settings := info.Settings
	for _, setting := range settings {
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
			stamp.GameRevision = setting.Value
		}
	}
	return &stamp
}

func retrieveDepends(module *debug.Module) (Dependency, bool) {
	var name, ver string

	path := strings.Split(module.Path, "/")
	if len(path) < 3 {
		return Dependency{}, false
	}
	name = path[2]
	if strings.Contains("golang.org", path[0]) {
		return Dependency{}, false
	}

	if len(module.Version) == 0 {
		ver = module.Sum
	} else {
		ver = module.Version
	}
	dep := Dependency{
		Name:    name,
		Version: ver,
	}

	return dep,
}

*/
