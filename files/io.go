package files

import (
	"io/ioutil"
	"os"
)

func ReadFileAsBytes(file *os.File) ([]byte, error) {
	if bytes, err := ioutil.ReadAll(file); err == nil {
		return bytes, nil
	} else {
		return nil, err
	}
}

func ReadFileAsString(file *os.File) (string, error) {
	z, err := ReadFileAsBytes(file)
	if err != nil {
		return "", err
	} else {
		return string(z), nil
	}
}

func ReadFileAsBytesByPath(path string) ([]byte, error) {
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		return ReadFileAsBytes(file)
	} else {
		return nil, err
	}
}

func ReadFileAsStringByPath(path string) (string, error) {
	z, err := ReadFileAsBytesByPath(path)
	if err != nil {
		return "", err
	} else {
		return string(z), nil
	}
}
