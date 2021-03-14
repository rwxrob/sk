package main

import (
	"os"
	"testing"
)

func TestDir_home(t *testing.T) {
	d := Dir()
	t.Log(d)
	if len(d) == 0 {
		t.Error("Dir is empty")
	}
}

func TestDir_env(t *testing.T) {
	os.Setenv("SKDIR", "/some")
	d := Dir()
	t.Log(d)
	if d != "/some" {
		t.Error("Failed to see SKDIR")
	}
}
