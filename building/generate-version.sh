#!/usr/bin/env sh
echo ">- generate: ${0} -<"
SCRIPTPATH="$(realpath $(dirname $0))"
OUT="${SCRIPTPATH}/../version.go"
echo "SCRIPTPATH='$SCRIPTPATH'"
echo "OUT='$OUT'"
echo "package go_config" > ${OUT}
echo \
"
const (
  version   = \"$(cat ${SCRIPTPATH}/VERSION | head -1)\"
  build     = \"$(git rev-list --count --all)\"
  branch    = \"$(git rev-parse --abbrev-ref HEAD)\"
  builduser = \"${USER}\"
  buildhost = \"$(hostname -f)\"
  buildos   = \"$(go env GOOS)\"
  buildarch = \"$(go env GOARCH)\"
  buildtime = \"$(date +"%Y-%m-%dT%H:%M:%S%z")\"
)

type Info struct {
    Version      string
    Build        string
    VersionBuild string
    Branch       string
    BuildUser    string
    BuildHost    string
    BuildOs      string
    BuildArch    string
    BuildTime    string

}

func getInfo() Info {
  return Info{
    Version:      version,
    Build:        build,
    VersionBuild: version + \"-b\" + build,
    Branch:       branch,
    BuildUser:    builduser,
    BuildHost:    buildhost,
    BuildOs:      buildos,
    BuildArch:    buildarch,
    BuildTime:    buildtime,
  }
}
" >> ${OUT}
