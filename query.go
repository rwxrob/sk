package main

import "strings"

func MembersMatching(args []string) []*Member {
	found := []*Member{}
	members := LoadMembers()
	config := LoadConfig()
	if len(args) < 1 {
		return found
	}
	first := strings.ToLower(args[0])
	for _, m := range members {
		switch {
		case strings.HasPrefix(first, config.MemberIDPrefix):
			if strings.HasPrefix(m.ID, first) {
				found = append(found, m)
			}
		case len(args) == 1:
			if strings.HasPrefix(strings.ToLower(m.FirstName), first) {
				found = append(found, m)
			}
		case len(args) == 2:
			last := strings.ToLower(args[1])
			if strings.HasPrefix(strings.ToLower(m.FirstName), first) &&
				strings.HasPrefix(strings.ToLower(m.LastName), last) {
				found = append(found, m)
			}
		}
	}
	return found
}
