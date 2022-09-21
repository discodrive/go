package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name and checks for a valid result
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    // If what we want back does not match the expected message
    // OR the error is anything other than nil
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloName calls greetings.Hello with an empty string,
// checking for an error
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
