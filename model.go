// keep all the data model stuff in this one file for easy
// changes as it evolves over time

package main

import (
	"fmt"
	"strconv"
	"time"
)

type Member struct {
	ID         string
	FirstName  string
	LastName   string
	Gender     string
	Born       string
	Email      string
	Phone      string
	Distance   string
	LastSchool string
	Contacts   []Contact
	Address    Address
	Accounts   map[string]string
	Reserved   []TimeSlot
	Enrollment []Enrollment
	Web        string
	Status     string // regular, periodic, former, banned
}

func (m *Member) ShortName() string {
	return fmt.Sprintf("%.10s %.1s.", m.FirstName, m.LastName)
}

func (m *Member) Emails() []string {
	emails := []string{}
	if len(m.Email) > 0 {
		emails = append(emails, m.Email)
	}
	for _, c := range m.Contacts {
		if len(c.Email) > 0 {
			emails = append(emails, c.Email)
		}
	}
	return emails
}

type Address struct {
	Local string
	City  string
	State string
	Zip   string
}

type Contact struct {
	FirstName string
	LastName  string
	Gender    string
	Email     string
	Phone     string
	Relation  string
	Note      string
}

// TimeSlot is an int between 100 and 723 indicating the day (hundreds) and
// hour (24 hour) during a given week for that slot. (Also see DateTime.)
type TimeSlot int

type Enrollment struct {
	Invoice  Invoice
	Sessions []Session
}

type Invoice struct {
	Number string
	Amount int
	Paid   bool
}

type Session struct {
	Date   string
	Hour   int
	Status int // 0 MISS, 1 HERE, 2 PUSH, 3 HOLI
	Note   string
	X      string
}

func (s Session) Past() bool {
	then, _ := time.Parse("2006-01-02 15", fmt.Sprintf("%v %01v", s.Date, s.Hour))
	return time.Now().After(then)
}

type Hours map[string]interface{}

// Hours returns the earliest hour from the set.
func (hours Hours) Earliest() int {
	earliest := 1000
	for hour, _ := range hours {
		h, _ := strconv.Atoi(hour[1:])
		if h < earliest {
			earliest = h
		}
	}
	return earliest
}

// Latest returns the latest hour from the set.
func (hours Hours) Latest() int {
	latest := 0
	for hour, _ := range hours {
		h, _ := strconv.Atoi(hour[1:])
		if h > latest {
			latest = h
		}
	}
	return latest
}

func (o Member) String() string     { return ConvertToJSON(o) }
func (o Address) String() string    { return ConvertToJSON(o) }
func (o Contact) String() string    { return ConvertToJSON(o) }
func (o TimeSlot) String() string   { return ConvertToJSON(o) }
func (o Enrollment) String() string { return ConvertToJSON(o) }
func (o Invoice) String() string    { return ConvertToJSON(o) }
func (o Session) String() string    { return ConvertToJSON(o) }
func (o Hours) String() string      { return ConvertToJSON(o) }
