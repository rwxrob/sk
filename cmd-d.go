package main

import (
	"fmt"

	"path/filepath"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("d")
	_cmd.Summary = "opens yaml data directory for the given member"
	_cmd.Parameters = `source schedule`
	_cmd.Usage = "[source|<id>|<first> [<last>]|schedule]"
	_cmd.Description = `
            Prints the directory matched from arguments. If called from the *sk*
            bashrc function also changes into the directory.
            
            If no argument is provided prints (and changes) into the *$SOILDIR*
            directory.
            
            If an <id> (beginning with the correct prefix) or the <first> and
            optionally <last> name of the member are provided will
            use *$SOILDIR/members/.../<id>* directory.

            If the *source* argument is passed will look for the configuration value
            of *SourceDir* inside of *$SOILDIR/config.yml* and change into that.
            This is useful if you have forked the *sk* utility and
            regularly want to customize your own source code.

            If the *schedule* argument is passed will print and change into the 
            *schedule* directory.
      `
	// TODO add completion based on name and/or id
	_cmd.Method = func(args []string) error {
		if len(args) == 0 {
			fmt.Println(Dir())
			return nil
		}
		if len(args) == 1 && args[0] == "source" {
			fmt.Println(LoadConfig().SourceDir)
			return nil
		}
		if len(args) == 1 && args[0] == "schedule" {
			fmt.Println(filepath.Join(Dir(), "schedule"))
			return nil
		}
		members := MembersMatching(args)
		if len(members) == 0 {
			return fmt.Errorf("Failed to find member matching %v", args)
		}
		m := members[0]
		fmt.Println(filepath.Join(Dir(), "members", m.Status, m.ID))
		return nil
	}
}
