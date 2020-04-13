package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "01/02/2006 15:04"
	)
	date := "3/23/2020 11:21"
	t, err := time.Parse(layoutISO, date)
	if err != nil {
		fmt.Println("ini", err)
	}
	fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
	fmt.Println(t.Format(layoutUS)) //
}
