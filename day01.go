package main

import (
	"fmt"
	"strconv"
)

type instructions struct {
	direction string
	amount    int
}

func day01a(puzzleData []string) {
	dialStart := 50
	InstructionList := day1DataProcess(puzzleData)
	dial := dialStart
	passwordCount := 0
	for _, instruction := range InstructionList {

		if instruction.amount > 99 {

			instruction.amount = instruction.amount % 100

		}

		if instruction.direction == "L" {
			dial = dial - instruction.amount

			if dial < 0 {
				dial = (100 + dial)

			}
		}
		if instruction.direction == "R" {
			dial = dial + instruction.amount
			if dial > 99 {
				dial = 0 + (dial - 100)

			}
		}

		if dial == 0 {
			passwordCount++
		}

	}
	fmt.Println("Day 01a Answer: ", passwordCount)
}
func day01b(puzzleData []string) {
	dialStart := 50
	InstructionList := day1DataProcess(puzzleData)
	dial := dialStart
	passwordCount := 0
	for _, instruction := range InstructionList {

		if instruction.amount > 99 {

			instruction.amount = instruction.amount % 100

		}

		if instruction.direction == "L" {
			dial = dial - instruction.amount

			if dial < 0 {
				dial = (100 + dial)

			}
		}
		if instruction.direction == "R" {
			dial = dial + instruction.amount
			if dial > 99 {
				dial = 0 + (dial - 100)

			}
		}

		if dial == 0 {
			passwordCount++
		}

	}
	fmt.Println("Day 01a Answer: ", passwordCount)
}

func day1DataProcess(rawData []string) []instructions {
	var InstructionList []instructions
	for _, line := range rawData {
		var instruction instructions
		instruction.direction = string(line[0])
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println(err)
		}
		instruction.amount = amount
		InstructionList = append(InstructionList, instruction)
	}
	return InstructionList
}
