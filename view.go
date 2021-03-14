package main

import (
	"fmt"
	"strings"
)

func Center(s string, w int) string {
	var space int // want floored division
	space = w - len(s)
	left := space / 2
	right := space - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

const pdoctable = `---- ------------- ------------- ------------- ------------- -------------`

func ViewHours(h Hours) {
	earliest := h.Earliest()
	latest := h.Latest()
	fmt.Println(pdoctable)
	for day := 0; day <= 5; day++ {
		if day == 0 {
			fmt.Printf("%4v", "")
			continue
		}
		fmt.Printf(" %-13v", Center(WDays[day], 13))
	}
	fmt.Println()
	fmt.Println(pdoctable)
	for hour := earliest; hour <= latest; hour++ {
		for day := 0; day <= 5; day++ {
			if day == 0 {
				fmt.Printf(" %2v ", hour)
				continue
			}
			val := ""
			hstr := fmt.Sprintf("%v%02v", day, hour)
			if v, has := h[hstr]; has {
				switch i := v.(type) {
				case string:
					val = i
				case nil:
					val = "OPEN"
				}
			}
			fmt.Printf(" %-13v", Center(val, 13))
		}
		fmt.Println()
	}
	fmt.Println(pdoctable)
}
