package simdog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	Version   = "None"
	GitHash   = "NoneHash"
	GitBranch = "NoneBranch"
	COYPRIGHT = "©2018-%d %s"
	Owner     = "jialinwu"
	BuildTime = "None"
	BuildUser = "None"
	BuildHost = "None"
	GoVersion = runtime.Version()
)

var versionTemplate = `
Version:%s
Branch:%s
Hash:%s
GoVersion:%s
BuildContext:%s
Copyright %s

`

func ReadBuildInfoNoExit() {
	buildContext := fmt.Sprintf("%s@%s %s", BuildUser, BuildHost, BuildTime)
	if len(GitHash) > 0 {
		fmt.Printf(versionTemplate, Version, GitBranch, GitHash[:7], GoVersion, buildContext, fmt.Sprintf(COYPRIGHT, time.Now().Year(), Owner))
	} else {
		fmt.Printf(versionTemplate, Version, GitBranch, GitHash, GoVersion, buildContext, fmt.Sprintf(COYPRIGHT, time.Now().Year(), Owner))
	}
}

func ReadBuildInfo() {
	if len(os.Args) > 1 {
		man := strings.TrimSpace(os.Args[1])
		switch man {
		case "version", "buildtime", "hash", "info", "branch", "v", "i":
			ReadBuildInfoNoExit()
			os.Exit(0)
		}
	}
	ReadBuildInfoNoExit()
}

// NewVersionCollector returns a collector that exports metrics about current version information.
func NewVersionCollector(appName string) prometheus.Collector {
	return prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: appName,
			Name:      "build_info",
			Help: fmt.Sprintf(
				"A metric with a constant '1' value labeled by version, hash, branch, and goversion etc from which %s was built.",
				appName,
			),
			ConstLabels: prometheus.Labels{
				"hash":      GitHash,
				"branch":    GitBranch,
				"version":   Version,
				"goversion": GoVersion,
				"buildhost": BuildHost,
				"builduser": BuildUser,
				"buildtime": BuildTime,
				"platform":  runtime.GOOS + "/" + runtime.GOARCH,
			},
		},
		func() float64 { return 1 },
	)
}
