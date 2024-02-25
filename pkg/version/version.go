package version

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"

	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"

	// Release is a semantic version of current build
	Release = "unset"
)

type BuildInfo struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Release   string `json:"release"`
}

func NewBuildInfo() *BuildInfo {

	return &BuildInfo{
		BuildTime: BuildTime,
		Commit:    Commit,
		Release:   Release,
	}
}
