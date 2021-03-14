package main

import (
	"fmt"
	//"strconv"
	//"time"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("reserved")
	_cmd.Parameters = "byid pandoc"
	_cmd.Summary = "print weekly reserved session slots"
	_cmd.Usage = "[byid|pandoc]"
	_cmd.Description = `
            This command prints the weekly reserved session slots. Members may
            have multiple reservations but no reservation may be shared by any
            other member.
            
            If a resevation is duplicated, a fatal model violation
            error will occur until it is fixed in the YAML data. This is by
            design to prevent any attempt to have more than one person per
            session.  One-on-one mentoring sessions is a foundational principle
            to Shared Sustainable Open Independent Learning.
      `
	_cmd.Method = func(args []string) error {
		hours := LoadHours()
		members := LoadMembers()
		for _, member := range members {
			if member.Reserved != nil && len(member.Reserved) > 0 {
				for _, r := range member.Reserved {
					rs := fmt.Sprintf("%v", r)
					if val, has := hours[rs]; has && val != nil {
						cmdtab.ExitError("Duplicate reservation: %v", val, member.ID, rs)
					}
					if len(args) > 0 && args[0] == "pandoc" {
						hours[rs] = "FULL"
						continue
					}
					if len(args) > 0 && args[0] == "byid" {
						hours[rs] = member.ID
						continue
					}
					hours[rs] = member.ShortName()
				}
			}
		}
		ViewHours(hours)
		return nil
	}
}
