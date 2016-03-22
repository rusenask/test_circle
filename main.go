package main

import (
    "fmt"
)

func main() {
    msg, _ := greet("circle")
    fmt.Println(msg)
    
}

func greet(name string) (string, error) {
    return fmt.Sprintf("Hello %s, this is just a test, first feature", name), nil
}