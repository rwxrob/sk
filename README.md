# Private Mentored Community Management Tool

[![GoDoc](https://godoc.org/github.com/rwxrob/sk?status.svg)](https://godoc.org/github.com/rwxrob/sk)
[![License](https://img.shields.io/badge/license-Apache-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/sk)](https://goreportcard.com/report/github.com/rwxrob/sk)
[![Coverage](https://gocover.io/_badge/github.com/rwxrob/sk)](https://gocover.io/github.com/rwxrob/sk)

The `sk` command (and optional `.bashrc` function) make the life of a
private mentored learning community manager easier but only for those
who prefer a command-line interface.

The utility specifically focuses on the following needs:

* Navigate into the YAML member, session, and other data directories quickly.
* Quickly add invoice and session data during a text-edit session (vim-centric).
* Print terminal views of member and session data for quick reference.
* Automated email progress reporting to parents, sponsors, and members.

The command comes with built in tab completion using [Complete
Commander](https://github.com/rwxrob/cmdtab).

## Ongoing Project

This particular tool is one that will be continually ongoing but is in
production use now at [SKILSTAK](https://skilstak.io).

Admittedly the biggest thing missing right now is full documentation
including specific examples of the YAML files used in the database.

## Design Decisions

Most of the design decisions are based on the fundamental premise that
this utility should be fast and easy to fork and customize for those who
wish. Subcommands can be easily added or removed simply by adding and
removing the `<subcommand>.go` file thanks to the use of [Complete
Commander](https://github.com/rwxrob/cmdtab).

* *Just YAML.* No mentored community should ever grow bigger than 60
  members and even that is pushing it. (Remember, the community
  mentor/leader must have mentoring sessions with each member about once
  a week and still do most of the work to manage the business and
  content of the community. 25 (5 hours a day)  is a much healthier
  number. Based on this expectation there simply is no need for a
  complicated database system, just simple YAML files that are extremely
  easy for humans to read and write without errors. Loading every
  separate file for the community into memory --- even after several
  years of attendance records --- takes a few milliseconds. This also
  allows ultimate flexibility in what information is maintained for any
  specific community.

* *Go language.* Go is simply the best language for this kind of utility
  for multiple reasons including YAML and JSON marshalling, ridiculously
  easy date and time function, built-in cryptography, cross compilation,
  simply concurrency, full Unicode support, strict typing, and fast
  development, compile, and execution time. Indeed this was the reason
  it was created originally at Google. There is simply no better
  language on the planet for such things at the moment, which is no
  surprise given that it was created by those involved in creating C,
  UNIX, UNICODE, and many of the modern Java fixes.

* *No separate library.* Just no need. The model is simple and
  documented elsewhere. The view and controller are fused.

