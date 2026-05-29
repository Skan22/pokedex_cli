package main

import (
	"bufio"
	"fmt"
	
	"os"
	"strings"
)

type cliCommand struct{
	name string
	description string
	callback func() error
}


func startRepl(){
	scanner :=bufio.NewScanner(os.Stdin)
	for {

	fmt.Print("> ")
	if !scanner.Scan(){
		break
	}

	text_slice:= cleanInput(scanner.Text())
	if len(text_slice) == 0 {
            continue
        }
	commands :=getCommands()
	command,exists:=commands[text_slice[0]]
	if exists {
		command.callback()
	}else {
		fmt.Println("Command not found")
	}
}
}





func getCommands()map[string]cliCommand{
	return map[string]cliCommand{
		"exit" : cliCommand{"exit",
						"Exit the Pokedex",
						commandExit,	},
		"help" : cliCommand{"help",
						"Displays a help message",
						commandHelp,

	},
}	
}




func commandExit()error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp()error{
	fmt.Println("Welcome to the Pokedex!","\nUsage:")
	for commandname,details:=range getCommands() {
		fmt.Println(commandname,":",details.description)
	}
	return nil
}
func cleanInput(text string) []string{
	text_slice := strings.Fields(strings.ToLower(text)) 
	return  text_slice
}