package main

import "testing"

func TestCallSayHi(t *testing.T) {
	name := "Henry Hai"
	expected := "Hi Henry Hai from public"
	greeting := callSayHi(name)
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}
