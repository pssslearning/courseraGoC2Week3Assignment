package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal an struct with its three main properties
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Associated Methods

// Eat prints out the Animal's kind of food it eats
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints out the Animal's kind of locomotion it uses
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints out the Animal's kind of noise it produces
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

// Main function: Entry Point
func main() {
	showInstructions()

	am := make(map[string]Animal)
	am["cow"] = Animal{"grass", "walk", "moo"}
	am["bird"] = Animal{"worms", "fly", "peep"}
	am["snake"] = Animal{"mice ", "slither", "hsss"}

	for {
		reqAnimal, reqAction := getUserRequest()

		animal, existing := am[reqAnimal]
		if !existing {
			fmt.Printf("Sorry, animal '%s' is not in my list. Please try with another animal.\n", reqAnimal)
			continue
		}

		switch reqAction {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("Sorry, action '%s' is not in my list of valid actions. Please try again.\n", reqAction)
		}
	}
}

func showInstructions() {
	fmt.Println("\n======== Purpose and Instructions ==============================================================")
	fmt.Println("This program presents the user with a prompt, “>”, to indicate that the user can type a request.")
	fmt.Println("Accepts one request at a time, prints out the answer to the request, and prints out a new prompt")
	fmt.Println("It should continue in this loop forever until 'CTRL + C' is pressed.")
	fmt.Println("Every request from the user must be a single line containing 2 strings:")
	fmt.Println(" 1. the name of an animal, either 'cow', 'bird', or 'snake'")
	fmt.Println(" 2. the name of the information requested about the animal, either 'eat', 'move', or 'speak'.")
	fmt.Println("Example: 'cow speak'")
	fmt.Println("==================================================================================================")

}

func getUserRequest() (string, string) {

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		in.Scan()
		scanned := in.Text()
		scanned = strings.TrimSpace(scanned)
		resulting := make([]string, 2, 2)
		cursor := 0
		splitted := strings.Split(scanned, " ")

		if len(splitted) < 2 {
			fmt.Println("Bad request format. Please try again")
			continue
		}

		// Nota: usando strings.Fields() no habría sido necesario este loop
		// ver AL/animals.go
		for _, y := range splitted {
			if len(y) > 0 {
				resulting[cursor] = y
				cursor++
				if cursor >= 2 {
					break
				}
			}
		}
		return strings.ToLower(resulting[0]), strings.ToLower(resulting[1])
	}
}
