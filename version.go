package go_config

const (
  version   = "v1.2.0"
  build     = "2"
  branch    = "main"
  builduser = "jens"
  buildhost = "MacBook-Pro-von-Jens.local"
  buildos   = "darwin"
  buildarch = "arm64"
  buildtime = "2024-01-29T10:58:47+0100"
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
    VersionBuild: version + "-b" + build,
    Branch:       branch,
    BuildUser:    builduser,
    BuildHost:    buildhost,
    BuildOs:      buildos,
    BuildArch:    buildarch,
    BuildTime:    buildtime,
  }
}

