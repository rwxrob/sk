package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// LoadMembers loads all members with MemberStatus "Regular" and "Periodic".
// Each call to LoadMembers loads and parses them directly from the YAML files
// with no caching. Since the amount of data in these files is minimal having
// the freshest data is more important than the unnecessary optimization that
// could come from caching.
func LoadMembers() []*Member {
	members := []*Member{}
	regulars := LoadMembersByStatus("regular")
	periodic := LoadMembersByStatus("periodic")
	if regulars != nil {
		members = append(members, regulars...)
	}
	if periodic != nil {
		members = append(members, periodic...)
	}
	return members
}

// LoadAllMembers loads all members with any status. See LoadMembers for more.
func LoadAllMembers() []*Member {
	members := []*Member{}
	regulars := LoadMembersByStatus("regular")
	periodic := LoadMembersByStatus("periodic")
	former := LoadMembersByStatus("former")
	banned := LoadMembersByStatus("banned")
	if regulars != nil {
		members = append(members, regulars...)
	}
	if periodic != nil {
		members = append(members, periodic...)
	}
	if former != nil {
		members = append(members, former...)
	}
	if banned != nil {
		members = append(members, banned...)
	}
	return members
}

func LoadMembersByStatus(status string) []*Member {
	if !(status == "regular" || status == "periodic" || status == "former" || status == "banned") {
		log.Printf("no such status: \"%v\"", status)
		return nil
	}
	members := []*Member{}
	mdir := filepath.Join(Dir(), "members", status)
	dir, err := os.Open(mdir)
	if err != nil {
		log.Printf("failed to load members from %v: %v", mdir, err)
		return nil
	}
	names, _ := dir.Readdirnames(0)
	for _, name := range names {
		path := filepath.Join(mdir, name, "member.yml")
		m := LoadMember(path)
		m.Status = status
		if m != nil {
			members = append(members, m)
		}
	}
	return members
}

func LoadMember(path string) *Member {
	m := new(Member)
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("failed to read member at %v: %v", path, err)
		return nil
	}
	err = yaml.Unmarshal(byt, m)
	if err != nil {
		log.Printf("failed to parse YAML for member at %v: %v", path, err)
		return nil
	}
	return m
}

// TODO consolidate Load*

func LoadHours() Hours {
	h := Hours{}
	path := filepath.Join(Dir(), "schedule", "hours.yml")
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("failed to read Hours at %v: %v", path, err)
		return nil
	}
	err = yaml.Unmarshal(byt, &h)
	if err != nil {
		log.Printf("failed to parse YAML for Hours at %v: %v", path, err)
		return nil
	}
	return h
}

func LoadConfig() Config {
	conf := Config{}
	path := filepath.Join(Dir(), "config.yml")
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("failed to read Config at %v: %v", path, err)
		return conf
	}
	err = yaml.Unmarshal(byt, &conf)
	if err != nil {
		log.Printf("failed to parse YAML for Config at %v: %v", path, err)
		return conf
	}
	return conf
}
