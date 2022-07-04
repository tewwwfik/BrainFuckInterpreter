package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

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

//Brainfuck interpreter uses bf files as argument
//You can use test.bf for example.
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("You should give the file as parameter : %s {filename.bf}\n", args[0])
		return
	}

	file := args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading %s\n", file)
		return
	}
	err, _ = RunBf(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished!")
}

func RunBf(input string) (err error, result string) {
	cells := make([]int, 2)
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
				return err, ""
			}
			loopMap[loopBeginningIndex] = ptr
			loopMap[ptr] = loopBeginningIndex
		}
		ptr++
	}
	//Creating loopMap from stack and all elements of stack should be popped!
	if len(loop) > 0 {
		return errors.New("Loop brackets dont match! Could not found matching ']'"), ""
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
			fmt.Printf("%c", cells[cellIndex])
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
		}
		ptr += 1
	}
	fmt.Printf("\n")
	return nil, result
}
