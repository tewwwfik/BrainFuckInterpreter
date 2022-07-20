package brainFuckInt

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Brainfuck interpreter uses bf files as argument
//You can use test.bf for example.
func main() {
	//custom command test
	test()
	//bf file example as argument:
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
	result, err := RunBf(string(data), nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
}

//If you want to add your custom code, you can follow these:
func test() {
	//Creates a new command List
	customCommands := newCommandList()
	cmd := Command{
		name: '#',
		calc: square,
	}
	cmd2 := Command{
		name: '2',
		calc: makeDouble,
	}
	cmd3 := Command{
		name:  'R',
		mvPtr: moveDoubleRight,
	}
	//We are adding new commands
	err := customCommands.Add(cmd)
	if err != nil {
		panic(err.Error())
	}
	err = customCommands.Add(cmd2)
	if err != nil {
		panic(err.Error())
	}
	err = customCommands.Add(cmd3)
	str := "+++2#.R+++2#.<+++2#."
	result, err := RunBf(str, customCommands)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func square(x int) int {
	return x * x
}

func makeDouble(x int) int {
	return 2 * x
}

func moveDoubleRight(x int) int {
	return x + 2
}
