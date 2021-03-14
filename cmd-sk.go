package main

import "github.com/rwxrob/cmdtab"

func init() {
	_cmd := cmdtab.New("sk", "psessions", "plog", "pinvoice", "reserved", "d", "genid", "bashrc", "roll", "email", "income")
	_cmd.Default = "d"
	_cmd.Summary = `manage a shared, sustainable, independent, open, learning community`
	_cmd.Description = `
            The *sk* command provide rudimentary assists for managing a SOIL
            mentored learning community. Rather than require a heavy database
            to maintain data is stored in simple YAML configuration files and
            can be edited with any text editor (although vim is strongly
            recommended). Therefore, most of the subcommands are to create data
            easily for these files and view the data in different
            human-friendly ways.
      `
}
