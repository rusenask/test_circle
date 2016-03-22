package main

import (
    "testing"
)

func TestGreet(t *testing.T) {
    resp, err := greet("circleci")
    if err != nil {
        t.Errorf("error was supposed to be nil")
    }
    if len(resp) < 1 {
        t.Errorf("response expected to be longer than this..")
    }
    
}