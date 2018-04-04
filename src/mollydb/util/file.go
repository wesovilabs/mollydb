package util

import (
	"os"
	"path/filepath"
	"strings"
)

//IsYaml detects if filename is a yaml
func IsYaml(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".yaml" || ext == ".yml"
}

//GetFileName returns lowered case filename
func GetFileName(prefix string, filename string) (string, string, error) {
	ext := filepath.Ext(filename)
	absPrefix, err := filepath.Abs(prefix)
	if err != nil {
		return "", "", err
	}
	absFilename, err := filepath.Abs(filename)
	if err != nil {
		return "", "", err
	}
	document := strings.Replace(absFilename, absPrefix, "", 1)
	document = strings.ToLower(strings.Replace(document, ext, "", 1))[1:]
	return document, ext, nil
}

//GetFiles returns files from directory
func GetFiles(baseDir string) []string {
	fileList := []string{}
	err := filepath.Walk(baseDir, func(path string, f os.FileInfo, err error) error {
		if baseDir != path {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		return make([]string, 0)
	}
	return fileList
}

//IsDir returns true if path belongs to a directory
func IsDir(pth string) (bool, error) {
	fi, err := os.Stat(pth)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}
