package rflags

import "testing"

// source="./data" debug output=out

type OutFlags struct {
	Source string `rflag:"source,s,src"`
	Debug  bool   `rflag:"debug,d"`
	Number int    `rflag:"a"`
	Output string
}

func TestParseFlags(t *testing.T) {
	f := OutFlags{}
	args := []string{`source="./data"`, `debug`, `output=out`, `a=1200`}
	if err := ParseFlags(&f, args); err != nil {
		t.Error(err)
	}
	if f.Source != "./data" {
		t.Errorf("Source should be: %s, got: %s", "./data", f.Source)
	}
	if f.Debug != true {
		t.Errorf("Debug should be: %v, got: %v", true, f.Debug)
	}
	if f.Output != "out" {
		t.Errorf("Output should be: %s, got: %s", "out", f.Output)
	}
	if f.Number != 1200 {
		t.Errorf("Number should be: %d, got: %d", 1200, f.Number)
	}
}
