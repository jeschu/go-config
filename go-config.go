package go_config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func loadConfig(path string) (f *os.File, closer func(), err error) {
	closer = func() {
		// nothing
	}
	var d string
	if d, err = os.UserConfigDir(); err != nil {
		return
	}
	if f, err = os.Open(filepath.Join(d, path)); err != nil {
		closer = func() { _ = f.Close() }
	}
	return
}

//goland:noinspection GoUnusedExportedFunction
func LoadConfigYaml(path string, target any) (err error) {
	var (
		f      *os.File
		closer func()
	)
	if f, closer, err = loadConfig(path); err != nil {
		return
	}
	defer closer()
	return yaml.NewDecoder(f).Decode(target)
}

//goland:noinspection GoUnusedExportedFunction
func LoadConfigJson(path string, target any) (err error) {
	var (
		f      *os.File
		closer func()
	)
	if f, closer, err = loadConfig(path); err != nil {
		return
	}
	defer closer()
	return json.NewDecoder(f).Decode(target)
}
