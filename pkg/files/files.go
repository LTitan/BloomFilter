package fileop

import (
	"io/ioutil"
	"os"
)

// CreateFileDir .
func CreateFileDir(path string) error {
	flag, err := PathExists(path)
	if err != nil {
		return err
	}
	if flag {
		return nil
	}
	return os.Mkdir(path, os.ModePerm)
}

// PathExists .
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// WtiteString .
func WtiteString(fileName, content string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return
}

// GetAllFile .
func GetAllFile(pathname string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			ret = append(ret, fi.Name())
		}
	}
	return ret, nil
}

// ReadFile .
func ReadFile(path, name string) (string, error) {
	ret, err := ioutil.ReadFile(path + "/" + name)
	if err != nil {
		return "", err
	}
	return string(ret), err
}
