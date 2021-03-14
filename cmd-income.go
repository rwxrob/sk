package main

import (
	"fmt"
	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("income")
	_cmd.Summary = "prints income estmates"
	_cmd.Method = func(args []string) error {
		c := LoadConfig()
		members := LoadMembersByStatus("regular")
		sc := 0
		for _, m := range members {
			c := len(m.Reserved)
			sc += c
		}
		fmt.Printf("Reservations        : %v\n", sc)
		fmt.Printf("Session Rate        : %v\n", c.SessionRate)
		fmt.Printf("Block Count         : %v\n", c.BlockCount)
		fmt.Printf("Blocks Per Year     : %v\n", c.SessionsPerYear/c.BlockCount)
		fmt.Printf("Sessions Per Year   : %v\n", c.SessionsPerYear)
		fmt.Printf("Income Per Member   : %v\n", c.SessionsPerYear*c.SessionRate)
		fmt.Printf("Annual Gross Income : %v\n", (c.SessionsPerYear*c.SessionRate)*float64(sc))
		return nil
	}
}
