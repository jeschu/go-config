package go_config

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"testing"
)

func TestGetInfo(t *testing.T) {
	info := GetInfo()
	assert.Equalf(t, "v", info.Version[:1], "info.Version should start with 'v'")
	assert.Regexpf(t, "^v\\d+\\.\\d+\\.\\d+$", info.Version, "info.Version should match")
	assert.NotEmptyf(t, info.Build, "info.Build not empty")
	assert.Equalf(t, info.Version+"-b"+info.Build, info.VersionBuild, "info.VersionBuild")
	assert.Equalf(t, runtime.GOOS, info.BuildOs, "info.BuildOs")
	assert.Equalf(t, runtime.GOARCH, info.BuildArch, "info.BuildArch")
	assert.Equalf(t, readVersion(t), info.Version, "info.Version == building/VERSION")
}

func readVersion(t *testing.T) string {
	f, err := os.Open("building/VERSION")
	if err != nil {
		t.Fatalf("unable to read 'building/VERSION: %v'", err)
	}
	defer func(f *os.File) { _ = f.Close() }(f)
	version, _ := bufio.NewReader(f).ReadString('\n')
	return version
}
