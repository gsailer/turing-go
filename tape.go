package main

import "fmt"

type TuringTape struct {
	Tape     []string // ****11010101****
	Position int
	State    string
}

func (t TuringTape) String() string {
	return fmt.Sprintf("Position: %d State: %s\nTape: %s", t.Position, t.State, t.Tape)
}
