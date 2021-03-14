package main

import (
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	MemberIDPrefix  string
	SourceDir       string
	Email           SMTPConfig
	BlockCount      float64
	SessionsPerYear float64
	SessionRate     float64
}

type SMTPConfig struct {
	Host      string
	Port      string
	From      string
	User      string
	Pass      string
	VerifyTLS bool
}

func (o Config) String() string { return ConvertToJSON(o) }

// Dir returns the SKDIR environment variable if found or the .main
// directory within the current user's home directory if not.
func Dir() string {
	d := os.Getenv("SKDIR")
	if d == "" {
		u, err := user.Current()
		if err == nil {
			d = filepath.Join(u.HomeDir, ".main")
		}
	}
	return d
}
