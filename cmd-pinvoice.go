package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("pinvoice")
	_cmd.Summary = "print new invoice YAML with sessions"
	_cmd.Usage = "<hour> <startdate> [<number> <amount> <paid>]"
	_cmd.Description = `
            Prints the YAML for a new invoice entry in a member.yml file
            including the standard block of 16 weekly sessions. The <hour> and
            <startdate> are required. The <number>, true or false for <paid>,
            and and <amount> can also be optionally provided.
      `
	_cmd.Method = func(args []string) error {
		if len(args) < 2 {
			return _cmd.UsageError()
		}
		hour := args[0]
		start := args[1]
		num := ""
		cost := "800"
		paid := "false"
		if len(args) > 2 {
			if len(args) < 5 {
				return _cmd.UsageError()
			}
			num = args[2]
			cost = args[3]
			paid = args[4]
		}
		fmt.Printf("- Invoice:\n    Number: '%v'\n    Amount: %v\n    Paid:   %v\n  Sessions:\n", num, cost, paid)
		cmdtab.Call("psessions", []string{hour, start})
		return nil
	}
}
