package main

import "fmt"

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

func compute(instructions []Instruction, t TuringTape) (TuringTape, int) {
	var symbol string
	var steps int
	for {
		if t.State == "F" {
			return t, steps
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
		steps++
	}
	return t, steps
}
