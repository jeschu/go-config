package go_config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
)

type ConfigEncoder func(writer io.Writer, target any) error
type ConfigDecoder func(reader io.Reader, target any) error

func configPath(path string) (string, error) {
	if configDir, err := os.UserConfigDir(); err != nil {
		return "", err
	} else {
		return filepath.Join(configDir, path), nil
	}
}

func OpenConfig(path string, write bool) (f *os.File, closer func(), err error) {
	closer = func() { /* nothing */ }
	var cp string
	if cp, err = configPath(path); err != nil {
		return
	}
	var flag int
	if write {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	} else {
		flag = os.O_RDONLY
	}
	if f, err = os.OpenFile(cp, flag, 0666); err != nil {
		closer = func() { _ = f.Close() }
	}
	return
}

func ReadConfig(path string, target any, decode ConfigDecoder) (err error) {
	var (
		f      *os.File
		closer func()
	)
	if f, closer, err = OpenConfig(path, false); err != nil {
		return
	}
	defer closer()
	return decode(f, target)
}

func WriteConfig(path string, source any, encode ConfigEncoder) (err error) {
	var (
		f      *os.File
		closer func()
	)
	if f, closer, err = OpenConfig(path, true); err != nil {
		return
	}
	defer closer()
	return encode(f, source)
}

//goland:noinspection GoUnusedExportedFunction
func ReadConfigYaml(path string, target any) (err error) {
	return ReadConfig(path, target, func(r io.Reader, target any) error { return yaml.NewDecoder(r).Decode(target) })
}

//goland:noinspection GoUnusedExportedFunction
func ReadConfigJson(path string, target any) (err error) {
	return ReadConfig(path, target, func(r io.Reader, target any) error { return json.NewDecoder(r).Decode(target) })
}

//goland:noinspection GoUnusedExportedFunction
func WriteConfigYaml(path string, source any) (err error) {
	return WriteConfig(path, source, func(w io.Writer, source any) error { return yaml.NewEncoder(w).Encode(source) })
}

//goland:noinspection GoUnusedExportedFunction
func WriteConfigJson(path string, source any) (err error) {
	return WriteConfig(path, source, func(w io.Writer, source any) error { return json.NewEncoder(w).Encode(source) })
}
