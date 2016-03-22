package main

import (
    "fmt"
)

func main() {
    greet("circle")
}

func greet(name string) (string, error) {
    return fmt.Sprintf("Hello %s, this is just a test", name), nil
}