package main

import (
	"os"
	"testing"
)

func TestLoadMember(t *testing.T) {
	m := LoadMember("testdata/members/regular/sl555555/member.yml")
	t.Log(m)
	//TODO do the comparison
}

func TestLoadMembers(t *testing.T) {
	os.Setenv("SKDIR", "testdata")
	m := LoadMembers()
	t.Log(len(m))
	t.Log(m)
	//TODO do the comparison
}

func TestLoadHours(t *testing.T) {
	os.Setenv("SKDIR", "testdata")
	h := LoadHours()
	t.Log(h)
	t.Logf("earliest: %v\n", h.Earliest())
	t.Logf("latest: %v\n", h.Latest())
	//TODO do the comparison
}
