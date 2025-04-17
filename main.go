package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	helloworld "github.com/jigyansunanda/Learn-Go-with-Tests/go-fundamentals/01-hello-world"
	adder "github.com/jigyansunanda/Learn-Go-with-Tests/go-fundamentals/02-integers"
)

func main() {
	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose your option between 1 to 2:\n")

	// Read the user input
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		return
	}

	// Trim any newline characters from the input
	input = strings.TrimSpace(input)

	// Convert the input string to an integer
	option, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		return
	}

	switch option {
	case 1:
		fmt.Println(helloworld.Hello(""))
		fmt.Println(helloworld.Hello("Jigyansu"))
	case 2:
		nums := make([]int, 0)
		for i := range 5 {
			nums = append(nums, i+1)
			fmt.Printf("Slice: %v, Sum: %d\n", nums, adder.Add(nums...))
		}
	default:
		fmt.Println("Looks like you are not ready to Go.😐")
	}
}
