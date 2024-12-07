package greetings

import (
	"regexp"
	"testing"
)

func TestGreetingName(t *testing.T) {
	name := "Tom"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting("Tom")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greeting("Tom") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetingEmpty(t *testing.T) {
	msg, err := Greeting("")
	if msg != "" || err == nil {
		t.Fatalf(`Greeting("") = %q, %v, want "", error`, msg, err)
	}
}
