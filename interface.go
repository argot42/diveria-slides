package main

import "fmt"

// START OMIT
type World struct {}

func (w *World) String() string {
    return "世界"
}

func main() {
    fmt.Printf("Hello, %s\n", new(World))
}
// END OMIT
