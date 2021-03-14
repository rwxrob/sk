package main

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("plog")
	_cmd.Summary = "print session log entry markdown to be completed"
	_cmd.Usage = ""
	_cmd.Description = `
            Prints out the following updated with the exact, immediate time:

                ## Friday, September 27, 2019 - 5:24:10PM

                Report:

                * 

                Session:

                * 

                Home:

                * 

      `
	_cmd.Method = func(args []string) error {
		fmt.Println(time.Now().Format("## Monday, January 2, 2006 - 3:04:05 PM"))
		fmt.Print("\nReport:\n\n*\n\nSession:\n\n*\n\nHome:\n\n*\n\n")
		return nil
	}
}
