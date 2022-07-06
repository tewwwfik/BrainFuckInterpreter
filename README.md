# Brianfuck Interpreter
This is a Brainfuck interpreter written in Golang. It also allows you to add custom commands.
## What is *Brainfuck*
> Brainfuck is an esoteric programming language created in 1993 by Urban Müller, and notable for its extreme minimalism.  
> The language consists of only eight simple commands and an instruction pointer. While it is fully Turing complete, it is not intended for practical use, but to challenge and amuse programmers. Brainfuck simply requires one to break commands into microscopic steps.  
> The language's name is a reference to the slang term brainfuck, which refers to things so complicated or unusual that they exceed the limits of one's understanding.

--from [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck)

## What does *Brainfuck* look like
A HelloWorld program looks like this, consisting of only eight types of commands.
```brainfuck
>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.
```

## Features of this interpreter
- Syntax errors detect
- Native Brainfuck grammar interpretation
- You can add your custom commands for both cell and pointer

# Usage
Compile using go build and run:
```brainfuck
./brainfuck filename.bf
```

You can check example.go or brainfuck_test.go for see usage as package and custom commands.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
