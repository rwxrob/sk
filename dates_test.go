package main

import (
	"fmt"
	"time"
)

func ExampleStartOfDay() {

	j, _ := time.Parse("2006-01-02 15:04", "2013-05-13 04:56")
	fmt.Println(StartOfDay(j))

	k, _ := time.Parse("2006-01-02 15:04", "2013-05-15 04:56")
	fmt.Println(StartOfDay(k))

	l, _ := time.Parse("2006-01-02 15:04", "2013-05-19 04:56")
	fmt.Println(StartOfDay(l))

	// Output:
	//
	// 2013-05-13 00:00:00 +0000 UTC
	// 2013-05-15 00:00:00 +0000 UTC
	// 2013-05-19 00:00:00 +0000 UTC
}
