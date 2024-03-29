package main

import (
	"./src/ptime"
	"./src/uitls"
	"fmt"
	"os"
	"time"
)

func main() {
	if !uitls.IsSet(os.Args, 1) {
		fmt.Println("未指定照片目录!")
		os.Exit(0)
	}
	if !ptime.IsDir(os.Args[1]) {
		println(os.Args[1], "不是有效的目录!")
		os.Exit(0)
	}

	s := ptime.List(os.Args[1])
	for _, j := range s {
		ft := ptime.GetFileTime(j)
		if ft > 0 {
			ptime.ChangeFileTime(j, ft)
			fmt.Println(time.Unix(ft, 0).Format("2006-01-02 15:04:05"), "->", j)
		} else {
			fmt.Println("未匹配成功: ", "->", j)
		}
	}
}
