package main

import (
	"fmt"
	"io/ioutil"
	fp "path/filepath"
	"strings"

	"github.com/rwxrob/cmdtab"
)

func init() {
	_cmd := cmdtab.New("email")
	_cmd.Summary = "sends <soildir>/email.md to member and member contacts"
	_cmd.Usage = "(<id>|<first> [<last>])"
	_cmd.Description = `
      `
	// TODO add completion based on name and/or id
	_cmd.Method = func(args []string) error {

		// read from standard form mail markdown file
		file := fp.Join(Dir(), "email.md")
		byt, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		body := string(byt)

		// TODO convert to multipart HTML+plain with Pandoc
		// TODO replace template variables in the form

		// extract the subject first line
		i := strings.Index(body, "\n\n")
		subject := body[0:i]
		body = body[i+2:]

		// lookup member
		members := MembersMatching(args)
		if len(members) == 0 {
			return fmt.Errorf("Failed to find member matching %v", args)
		}
		m := members[0]

		return Email(m.Emails(), subject, body)
	}
}
