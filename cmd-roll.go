package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("roll")
	_cmd.Summary = "prints attendance roll"
	_cmd.Usage = "(<id> | <first> [<last>])"
	_cmd.Description = `
            Prints attendance roll for specified member.
      `
	// TODO add completion based on name and/or id
	_cmd.Method = func(args []string) error {
		members := MembersMatching(args)
		if len(members) == 0 {
			return fmt.Errorf("Failed to find member matching %v", args)
		}
		m := members[0]
		for _, current := range m.Enrollment {
			var num, last int
			for _, s := range current.Sessions {
				status := ""
				switch {
				case s.Past() && s.Status == 0:
					status = "MISS"
					num += 1
				case s.Status == 0:
					num += 1
				case s.Status == 1:
					status = "HERE"
					num += 1
				case s.Status == 2:
					status = "PUSH"
				}
				text := status
				if len(text) > 0 {
					text += " "
				}
				if len(s.Note) > 0 {
					text += s.Note
				}
				if len(s.X) > 0 {
					text += s.X
				}
				if num == 13 && s.Status != 2 {
					text += "(invoice)"
				}
				var sym interface{}
				if num == last {
					sym = " → "
				} else {
					sym = fmt.Sprintf("%2v.", num)
				}
				fmt.Printf("%v %v %v %v\n", sym, s.Date, s.Hour, text)
				last = num
			}
			fmt.Println()
		}
		return nil
	}
}
