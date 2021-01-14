package main

import (
	"file"
	"fmt"
	"os"
	"time"
)

func main() {
	s := file.List(os.Args[1])
	for _, j := range s {
		ft := file.GetFileTime(j)
		if ft > 0 {
			fmt.Println(time.Unix(ft, 0).Format("2006-01-02 15:04:05"),"->",j)
			file.ChangeFileTime(j,ft)
		}
	}
}
