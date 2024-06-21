package version

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

// # build with verison infos
// projectName = "yourApp"
// versionDir = "git.code.oa.com/tme/version"
// gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
// gitCommit = $(shell git log --pretty=format:'%H' -n 1)
// gitBranch = $(shell git symbolic-ref --short -q HEAD)
// gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
// buildAuthor = $(shell if [ "`git config user.name`" != "" ]; then git config user.name; else hostname; fi)
// buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
// goVersion = $(shell go version | awk -F" " '{print $$3}')
// platform = $(shell go version | awk -F" " '{print $$4}')

// ldflags="-w -X ${versionDir}.name=${projectName} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.gitCommit=${gitCommit} \
// 		-X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.buildAuthor=${buildAuthor} \
// 		-X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.goVersion=${goVersion} \
// 		-X ${versionDir}.platform=${platform}"

// # go build with ldflags
// go build -ldflags ${ldflags}

var (
	name         = "unknown"              // name of software
	gitTag       = ""                     // git tag if have
	gitCommit    = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "not a git tree"       // state of git tree, either "clean" or "dirty"
	gitBranch    = "unknown"              // current git branch
	buildAuthor  = "user"                 // the author build the bin
	buildDate    = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	goVersion    = "go version"           // go version.
	platform     = "platform"             // platform
)

func init() {
	args := os.Args
	if nil == args || len(args) < 2 {
		fmt.Println(args)
		return
	}

	// 不用flag 防止引用的库使用了flag库导致参数无法读入
	if "-v" == args[1] || "-version" == args[1] || "version" == args[1] {
		fmt.Println(GetString())
		os.Exit(0)
	}
}

// Info contains versioning information.
type Info struct {
	Name         string `json:"name"`
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	GitBranch    string `json:"gitBranch"`
	BuildAuthor  string `json:"buildAuthor"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

// Get return version info.
func Get() Info {
	return Info{
		Name:         name,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		GitBranch:    gitBranch,
		BuildDate:    buildDate,
		BuildAuthor:  buildAuthor,
		GoVersion:    goVersion,
		Compiler:     runtime.Compiler,
		Platform:     platform,
	}
}

// GetString return version info string(json format).
func GetString() string {
	v := Get()
	marshalled, err := json.MarshalIndent(&v, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(marshalled)
}
