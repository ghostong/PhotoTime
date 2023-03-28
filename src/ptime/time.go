package ptime

import (
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
	ts, isTs := isTimeStamp(s)
	if isTs {
		return ts
	}
	ts, isDt6 := isDateTime6(s)
	if isDt6 {
		return ts
	}
	ts, isDt2 := isDateTime2(s)
	if isDt2 {
		return ts
	}
	ts, isDt3 := isDateTime3(s)
	if isDt3 {
		return ts
	}
	ts, isDt4 := isDateTime4(s)
	if isDt4 {
		return ts
	}
	ts, isDtAny := isDateTimeAny(s)
	if isDtAny {
		return ts
	}
	return 0
}

//lv_6847479811028421902_20200928172242.mp4
func isDateTimeAny(s [][]string) (int64, bool) {
	for _, j := range s {
		dt := j[0]
		if len(dt) == 14 {
			loc, _ := time.LoadLocation("Local")
			ts, err := time.ParseInLocation("20060102150405", dt, loc)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return ts.Unix(), true
				}
			}
		}
	}
	return 0, false
}

//IMG_20210107_204835R.jpg
func isDateTime2(s [][]string) (int64, bool) {
	if len(s) == 2 {
		dt := ""
		for _, j := range s {
			dt += j[0]
		}
		if len(dt) == 14 {
			loc, _ := time.LoadLocation("Local")
			ts, err := time.ParseInLocation("20060102150405", dt, loc)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return ts.Unix(), true
				}
			}
		}
	}
	return 0, false
}

//img_20210508_190915r(2).jpg
func isDateTime3(s [][]string) (int64, bool) {
	if len(s) == 3 {
		s = s[0:2]
		dt := ""
		for _, j := range s {
			dt += j[0]
		}
		if len(dt) == 14 {
			loc, _ := time.LoadLocation("Local")
			ts, err := time.ParseInLocation("20060102150405", dt, loc)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return ts.Unix(), true
				}
			}
		}
	}
	return 0, false
}

func isDateTime4(s [][]string) (int64, bool) {
	if len(s) == 4 {
		s = s[0:2]
		dt := ""
		for _, j := range s {
			dt += j[0]
		}
		if len(dt) == 14 {
			loc, _ := time.LoadLocation("Local")
			ts, err := time.ParseInLocation("20060102150405", dt, loc)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return ts.Unix(), true
				}
			}
		}
	}
	return 0, false
}

//pt2021_01_02_19_23_38.jpg
func isDateTime6(s [][]string) (int64, bool) {
	if len(s) >= 6 {
		s := s[0:6]
		dt := ""
		for _, j := range s {
			dt += j[0]
		}
		if len(dt) == 14 {
			loc, _ := time.LoadLocation("Local")
			ts, err := time.ParseInLocation("20060102150405", dt, loc)
			if err == nil {
				if isAllowTs(ts.Unix()) {
					return ts.Unix(), true
				}
			}
		}
	}
	return 0, false
}

//159006351462438.PNG
//img_159006351462438_2894792112832.gif
func isTimeStamp(s [][]string) (int64, bool) {
	for _, j := range s {
		if len(j[0]) > 10 {
			ts, _ := strconv.ParseInt(j[0][0:10], 10, 64)
			if isAllowTs(ts) {
				return ts, true
			}
		}
	}
	return 0, false
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

//时间转本地时间戳
func dtToLocalTs(dt string) (time.Time, bool) {
	loc, _ := time.LoadLocation("Local")
	ts, err := time.ParseInLocation("20060102150405", dt, loc)
	if err == nil {
		return ts, true
	} else {
		return time.Time{}, false
	}
}
