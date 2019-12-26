package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"./engine"
)

// printCommand - our command for priniting messages to console
type printCommand struct {
	arg string
}

// Execute - run command
func (pc *printCommand) Execute(loop engine.Handler) {
	fmt.Println(pc.arg)
}

// splitCommand - our command for splitting string with separator and print each
type splitCommand struct {
	//main string
	arg1 string
	//separator
	arg2 string
}

// Execute - run command
func (sp *splitCommand) Execute(loop engine.Handler) {
	stringSlice := strings.Split(sp.arg1, sp.arg2)
	for _,el := range stringSlice{
		loop.Post(&printCommand{arg: el})
	}
}

// parseInput - parsing messages to commands structs
func parseInput(commandline string) engine.Command {
	parts := strings.Fields(commandline)
	if parts[0]=="print" {
		return &printCommand{arg: parts[1]}
	}	else if parts[0] == "split" {
		return &splitCommand{arg1:parts[1],arg2:parts[2] }
	} else {
		return &printCommand{arg: "Syntax Error Unexpected command"}
	}
}

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open("./commands.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			eventLoop.Post(parseInput(scanner.Text()))
		}
	}
	eventLoop.AwaitFinish()
}