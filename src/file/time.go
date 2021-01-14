package file

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func GetFileTime(file string) int64 {
	_, fileName := filepath.Split(file)
	reg := regexp.MustCompile("[0-9]+")
	if nil == reg {
		return 0
	}
	result := reg.FindAllStringSubmatch(fileName, -1)
	return matchTime(result)
}

//匹配时间戳
func matchTime(s [][]string) int64 {
	isTs, ts := isTimeStamp(s)
	if isTs {
		return ts
	}
	isDt6, ts := isDateTime6(s)
	if isDt6 {
		return ts
	}
	isDt2, ts := isDateTime2(s)
	if isDt2 {
		return ts
	}
	isDtAny, ts := isDateTimeAny(s)
	if isDtAny {
		return ts
	}
	fmt.Println(s)
	return 0
}

func isDateTimeAny(s [][]string) (bool, int64) {
	for _, j := range s {
		dt := j[0]
		if len(dt) == 14 {
			ts, err := time.Parse("20060102150405", dt)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return true, ts.Unix()
				}
			}
		}
	}
	return false, 0
}

func isDateTime2(s [][]string) (bool, int64) {
	if len(s) == 2 {
		dt := ""
		for _, j := range s {
			dt += j[0]
		}
		if len(dt) == 14 {
			ts, err := time.Parse("20060102150405", dt)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return true, ts.Unix()
				}
			}
		}
	}
	return false, 0
}

func isDateTime6(s [][]string) (bool, int64) {
	if len(s) == 6 {
		dt := ""
		for _, j := range s {
			dt += j[0]
		}
		if len(dt) == 14 {
			ts, err := time.Parse("20060102150405", dt)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return true, ts.Unix()
				}
			}
		}
	}
	return false, 0
}

//时间是否时间戳
func isTimeStamp(s [][]string) (bool, int64) {
	for _, j := range s {
		if len(j[0]) > 10 {
			ts, _ := strconv.ParseInt(j[0][0:10], 10, 64)
			if isAllowTs(ts) {
				return true, ts
			}
		}
	}
	return false, 0
}

//是否有效时间戳
func isAllowTs(ts int64) bool {
	year, _ := strconv.Atoi(time.Unix(ts, 0).Format("20060102150405"))
	now, _ := strconv.Atoi(time.Unix(time.Now().Unix(), 0).Format("20060102150405"))
	if 20000101000000 <= year && now > year { //预估数码相机在中国出现时间为2000年左右 与 现在时间
		return true
	} else {
		return false
	}
}
