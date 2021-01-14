package file

import (
	"os"
	"time"
)

func ChangeFileTime(file string, ts int64) bool {
	ut := time.Unix(ts, 0)
	err := os.Chtimes(file, ut, ut)
	if nil == err {
		return false
	} else {
		return true
	}
}
