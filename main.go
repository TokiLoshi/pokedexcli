package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting Pokedex")

	commandLine := configure()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Go Gopher shell")
	fmt.Println("******************")
	count := 0
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		if command, ok := commandLine[text]; ok {
			// at this point we get "help" or exit 
			fmt.Printf("Does this work: %v\n", command.description)
			commandLine[text].callback()
		count++
	}
}
	// Goal: when you run the program you should see a prompt 
	// It should wait for you to type in a command
	// then print the result of the command 
	// the prompt should look like pokedex >
	// you should have help - print a help message describing how to use REPL
	// exit command which exits the program 
	// Use an infinite for loop to keep REPL running
	// at the start of the loop block and wait for input
	// once input is received parse it 
	// then execute command 
	// when command is finished print output
	// go to next iteration and wait for more input 
	// get input using bufio.NewScanner 
	// .Scan 
	// .Text 


}