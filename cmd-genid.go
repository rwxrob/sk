package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/uniq-go"
)

func init() {
	_cmd := cmdtab.New("genid")
	_cmd.Summary = `generate a unique member ID`
	_cmd.Method = func(args []string) error {
		config := LoadConfig()
		members := LoadAllMembers()
		var id string
	roll:
		id = uniq.Hex(3)
		for _, m := range members {
			if m.ID == id {
				goto roll
			}
		}
		fmt.Printf("%v%v\n", config.MemberIDPrefix, id)
		return nil
	}
}
