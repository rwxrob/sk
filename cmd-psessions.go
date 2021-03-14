package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("psessions")
	_cmd.Summary = "print a range of weekly sessions in YAML"
	_cmd.Usage = "[<hour> [<startdate> [<count>]]]"
	_cmd.Description = `
            This command is helpful when editing the YAML SOIL configuration
            files containing blocks of sessions. The <hour> defaults to the
            current hour. An optional <startdate> can be provided, and if so an
            option <count> of the number of sessions to print. The <count>
            defaults to 16.
            
            Keep in mind that blocks often span daylight savings time
            boundaries but thata the <hour> will never change in the output.
            This could provide some edge cases where the dates might be
            slightly off if for some reason the <hour> is 23 (which are never
            used in practice).
      `
	_cmd.Method = func(args []string) error {
		var err error

		hour := time.Now().Local().Hour()
		if len(args) > 0 {
			hour, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		}

		date := Today()
		if len(args) > 1 {
			date, err = time.Parse("2006-01-02", args[1])
			if err != nil {
				return err
			}
		}

		count := 16
		if len(args) > 2 {
			count, err = strconv.Atoi(args[2])
			if err != nil {
				return err
			}
		}

		for i := 0; i < count; i++ {
			fmt.Printf("  - { Date: '%v', Hour: %v, Status: 0 }\n", date.Format("2006-01-02"), hour)
			date = date.AddDate(0, 0, 7)
		}
		return nil
	}
}
