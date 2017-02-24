package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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
