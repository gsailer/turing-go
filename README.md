# Turing-go

turing-go is a simulation for Turing-machine code in the format
	
	state symbol/read -> symbol/write headmovement next-state

	A 1 -> * R A
	
the instructions are written in a YAML file with the structure as shown in example.yaml

You have to pass the executable the instruction-file and the tape content as arguments:
	
	turing-go [instruction-file] [tape]
	
The Tape has to be passed with four leading and trailing stars like:
	
	****1001****
