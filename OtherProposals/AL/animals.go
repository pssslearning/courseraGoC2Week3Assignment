package main

import (
	"fmt"
	// "sort"
	"os"
	"bufio"
	// "strconv"
	"strings"
)

	type Animal struct {
		food string
		move string
		noise string
	}

	func (obj Animal) Eat() {
		fmt.Println(obj.food)
	}

	func (obj Animal) Move() {
		fmt.Println(obj.move)
	}

	func (obj Animal) Speak() {
		fmt.Println(obj.noise)
	}

func main() {

	bird := Animal{food:"worms", noise:"peep", move:"fly"}
	cow := Animal{food:"grass", noise:"moo", move:"walk"}
	snake := Animal{food:"mice", noise:"hsss", move:"slither"}


	for true {
		fmt.Printf("> ")
		in := bufio.NewReader(os.Stdin)
		user_input, err := in.ReadString('\n')
		user_input = strings.Trim(user_input, "\n")
		if err != nil {
			fmt.Println(err)
			return
		}

		slice := strings.Fields(user_input)
		object := slice[0]
		action := slice[1]
		switch {
		case object == "cow" && action == "eat":
			cow.Eat()
		case object == "cow" && action == "move":
			cow.Move()
		case object == "cow" && action == "speak":
			cow.Speak()
		case object == "bird" && action == "eat":
			bird.Eat()
		case object == "bird" && action == "move":
			bird.Move()
		case object == "bird" && action == "speak":
			bird.Speak()
		case object == "snake" && action == "eat":
			snake.Eat()
		case object == "snake" && action == "move":
			snake.Move()
		case object == "snake" && action == "speak":
			snake.Speak()

		}
	}

}