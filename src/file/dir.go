package file

import (
	"io/ioutil"
	"log"
	"strings"
)

//FileList
func List(path string) []string {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var s []string
	for _, file := range fileInfoList {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if !IsAllowExt(file.Name()) {
			continue
		}
		if !file.IsDir() {
			s = append(s, path+file.Name())
		} else {
			s2 := List(path + file.Name() + "/")
			s = append(s, s2...)
		}
	}
	return s
}

//IsAllowExt
func IsAllowExt(fileName string) bool {
	sl := strings.Split(fileName, ".")
	ext := sl[len(sl)-1]
	extMap := map[string]bool{
		"jpg": true,
		"png": true,
		"mp4": true,
	}
	_, ok := extMap[ext]
	if ok {
		return true
	} else {
		return false
	}
}
