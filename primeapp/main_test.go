package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{name: "prime", testNum: 7, expected: true, msg: "7 is a prime number!"},
		{name: "not prime", testNum: 8, expected: false, msg: "8 is not prime, because it is divisible by 2!"},
		{name: "zero", testNum: 0, expected: false, msg: "0 is not prime, by definition!"},
		{name: "one", testNum: 1, expected: false, msg: "1 is not prime, by definition!"},
		{name: "negative", testNum: -5, expected: false, msg: "Negative numbers are not prime, by definition!"},
	}

	for _, test := range primeTests {
		result, msg := isPrime(test.testNum)
		if test.expected && !result {
			t.Errorf("%s: expected true but false", test.name)
		}

		if !test.expected && result {
			t.Errorf("%s: expected false but true", test.name)
		}

		if test.msg != msg {
			t.Errorf("%s: expected %s, but got %s", test.name, test.msg, msg)
		}
	}
}

func TestPrompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // set os.Stdout to write pipe
	prompt()
	_ = w.Close()
	os.Stdout = oldOut // Reset os.Stdout
	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("Incorrect prompt: expected -> but got %s", string(out))
	}
}

func TestIntro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // set os.Stdout to write pipe
	intro()
	_ = w.Close()
	os.Stdout = oldOut // Reset os.Stdout
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.") {
		t.Errorf("intro test not corrext: got %s", string(out))
	}
}
