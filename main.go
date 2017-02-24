package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please give a file and tape content: ./turing-go [filepath] [tape]")
		os.Exit(1)
	}
	path := args[0]
	tape_raw := args[1]
	var tape_array []string
	for i := 0; i < len(tape_raw); i++ {
		tape_array = append(tape_array, string(tape_raw[i]))
	}
	tape := TuringTape{Tape: tape_array, Position: 4, State: "A"} // position 4, because of 4 leading *
	instructions, err := parseInstructions(path)
	if err != nil {
		panic(err)
	}
	set, steps := compute(instructions, tape)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Finished after %d steps: %v\n", steps, set)
}
