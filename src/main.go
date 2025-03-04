package main

import (
	"fmt"
	"os"
)

func printArguments(args []string) {
	fmt.Println(args)
}

func main() {
	args := os.Args[1:]
	printArguments(args)
}
