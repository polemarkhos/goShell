package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Started.")
	// Initialise

	//Interpret
	interpret()
	//Exit after interpretation loop terminates
	defer os.Exit(1)
}

func interpret() {
	// this loop will run continuously to capture keyboard I/O until exit
	for {
		reader := bufio.NewReader(os.Stdin)
		// writer := bufio.NewWriter(os.Stdout)

		input, _ := reader.ReadString('\n')

		executeInput(input)

	}
}

func findCommand(args string) []string {

	cleanupArgs := func(args string) []string {
		argSlice := strings.Split(args, " ")
		lastElem := len(argSlice) - 1
		argSlice[lastElem] = strings.TrimRight(argSlice[lastElem], "\n")
		return argSlice
	}
	argSlice := cleanupArgs(args)

	return argSlice
}

func executeInput(args string) {
	argSlice := findCommand(args)
	fmt.Println(argSlice)
	if argSlice[0] == "exit" {
		fmt.Println("Shutting down.")
		os.Exit(0)
	}

	cmd := exec.Command(argSlice[0], argSlice[1:]...)

	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	slurp, _ := io.ReadAll(stdout)
	fmt.Printf("%s", slurp)

	cmd.Wait()
}
