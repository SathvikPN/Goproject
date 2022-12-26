package greetings

import "testing"

func TestHelloName(t *testing.T) {
	name := "sathvik"
	want := "Hello sathvik"
	msg, err := Hello(name)
	if msg != want || err != nil {
		t.Fatalf("Hello(%s) = %#q, %#v    want %#q, nil", name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fail()
	}
}

func TestHelloSpace(t *testing.T) {
	name := "    "
	want := ""
	msg, err := Hello(name)
	if msg != want || err != nil {
		t.Fatalf("\n[call] Hello(%s) \n[got] %#q, %#v    \n[want] %#q, %#v", name, msg, err, want, nil)
	}
}
