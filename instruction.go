package main

import "fmt"

type Instructions struct {
	Instruction []Instruction `instructions`
}

type Instruction struct {
	CurrentState string `current-state`
	SymbolR      string `symbol-read`
	SymbolW      string `symbol-write`
	Head         string `head`
	NextState    string `next-state`
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %s -> %s %s %s", i.CurrentState, i.SymbolR, i.SymbolW, i.Head, i.NextState)
}
