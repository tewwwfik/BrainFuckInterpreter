package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var cells = make([]int, 30000)

type Command struct {
	name  int
	calc  func(int) int
	mvPtr func(int) int
}

type commandListT map[int]Command

func newCommandList() commandListT {
	return make(commandListT, 10)
}

func (cl commandListT) Add(cmd Command) (err error) {
	if _, ok := cl[cmd.name]; ok {
		return errors.New("Command Already Exist!")
	}
	if cmd.name > 127 {
		return errors.New("Command not exist in ASCII table!")
	}
	cl[cmd.name] = Command{
		cmd.name,
		cmd.calc,
		cmd.mvPtr,
	}
	return nil
}

//holds the loops as stack
type loopStack []int

func (s loopStack) push(v int) loopStack {
	return append(s, v)
}
func (s loopStack) pop() (loopStack, int, error) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1], nil
	} else {
		return nil, 0, errors.New("Loop brackets dont match! Could not found matching '['")
	}
}

func RunBf(input string, customCommands commandListT) (result string, err error) {
	loop := make(loopStack, 0)
	//loop map holds the ptr of loops.
	//open loops holds closed location and close loops holds open location
	loopMap := make(map[int]int)
	var ptr int = 0
	reader := bufio.NewReader(os.Stdin)
	//Creates loop indexes. Push and pop from the stack and stores open and close position to a map.
	//loopMap holds openLoopAddress->closeLoopAddress and closeLoopAddress->openLoopAddress
	for _, i := range input {
		if i == '[' {
			loop = loop.push(ptr)
		} else if i == ']' {
			var loopBeginningIndex int
			var err error
			loop, loopBeginningIndex, err = loop.pop()
			if err != nil {
				return "", err
			}
			loopMap[loopBeginningIndex] = ptr
			loopMap[ptr] = loopBeginningIndex
		}
		ptr++
	}
	//Creating loopMap from stack and all elements of stack should be popped!
	if len(loop) > 0 {
		return "", errors.New("Loop brackets dont match! Could not found matching ']'")
	}
	//Starting to execute program.
	ptr = 0
	var cellIndex int
	for ptr < len(input) {
		//instruction
		i := input[ptr]
		switch i {
		case '+':
			cells[cellIndex] += 1
			if cells[cellIndex] == 256 {
				cells[cellIndex] = 0
			}
		case '-':
			cells[cellIndex] -= 1
			if cells[cellIndex] == -1 {
				cells[cellIndex] = 256
			}
		case '<':
			cellIndex -= 1
		case '>':
			cellIndex += 1
			if cellIndex == len(cells) {
				cells = append(cells, 0)
			}
		case '.':
			//fmt.Printf("%c", cells[cellIndex]) //Prints to console optional!
			result += fmt.Sprintf("%c", cells[cellIndex])
		case ',':
			value, _ := reader.ReadByte()
			cells[cellIndex] = int(value)
		case '[':
			// if currentValue is zero it moves to close loop position, and it finds position from loopMap.
			if cells[cellIndex] == 0 {
				ptr = loopMap[ptr]
			}
		case ']':
			//if currentValue is not zero it moves to beginning of loop, and it finds beginning position from loopMap.
			if cells[cellIndex] != 0 {
				ptr = loopMap[ptr]
			}
		default:
			//Check if there is custom commands.
			if customCommands != nil {
				if v, ok := customCommands[int(i)]; ok {
					//Calls custom command implemented function.
					if v.calc != nil {
						cells[cellIndex] = v.calc(cells[cellIndex])
					} else if v.mvPtr != nil {
						cellIndex = v.mvPtr(cellIndex)
					}
				}
			}
		}
		ptr += 1
	}
	fmt.Printf("\n")
	return result, nil
}
