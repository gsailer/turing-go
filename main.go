package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

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

type TuringTape struct {
	Tape     []string // ****11010101****
	Position int
	State    string
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %s -> %s %s %s", i.CurrentState, i.SymbolR, i.SymbolW, i.Head, i.NextState)
}

func (t TuringTape) String() string {
	return fmt.Sprintf("Position: %d State: %s\nTape: %s", t.Position, t.State, t.Tape)
}

func parseInstructions(path string) ([]Instruction, error) {
	var instructions Instructions
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Reading file %s: %v", path, err)
	}
	err = yaml.Unmarshal(source, &instructions)
	if err != nil {
		return nil, fmt.Errorf("Unmarshalling YAML %s: %v", path, err)
	}
	return instructions.Instruction, nil
}

func moveAndWrite(instructions []Instruction, t *TuringTape, symbol string) {
	for _, ins := range instructions {
		if ins.CurrentState == t.State && ins.SymbolR == symbol {
			fmt.Println("Symbol: ", symbol)
			fmt.Println("Instruction: ", ins)
			t.Tape[t.Position] = ins.SymbolW
			fmt.Println("Wrote: ", ins.SymbolW)
			t.State = ins.NextState
			switch ins.Head {
			case "R":
				t.Position++
				fmt.Println(t.Tape)
				break
			case "L":
				t.Position--
				fmt.Println(t.Tape)
				break
			case "N":
				fmt.Println(t.Tape)
				break
			}
		}
	}
}

func compute(instructions []Instruction, t TuringTape) TuringTape {
	var symbol string
	for {
		if t.State == "F" {
			return t
		}
		fmt.Println("-----------------------------------------------")
		fmt.Printf("State: %s\nPosition: %d\n", t.State, t.Position)
		symbol = t.Tape[t.Position]
		switch symbol {
		case "0":
			moveAndWrite(instructions, &t, symbol)
			break
		case "1":
			moveAndWrite(instructions, &t, symbol)
			break
		case "*":
			moveAndWrite(instructions, &t, symbol)
			break
		}
	}
	return t
}

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
	set := compute(instructions, tape)
	fmt.Println("-----------------------------------------------")
	fmt.Println("Finished: ", set)
}
