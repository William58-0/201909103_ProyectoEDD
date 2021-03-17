package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hola hola"
	b := strings.ReplaceAll(a, " ", "_")
	fmt.Println(b)
}
