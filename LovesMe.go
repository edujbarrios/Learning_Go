package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random sequence of "Loves Me" and "Loves Me Not" phrases
	lovesMe := []string{"Loves Me", "Loves Me Not"}
	sequence := []string{}
	for i := 0; i < rand.Intn(10)+5; i++ {
		sequence = append(sequence, lovesMe[rand.Intn(2)])
	}

	// Print the sequence to the console
	fmt.Println(strings.Join(sequence, " "))

	// Prompt the user to guess the final phrase
	fmt.Print("Will the final phrase be 'Loves Me' or 'Loves Me Not'? ")

	// Read the user's guess from the command line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	guess := scanner.Text()

	// Check if the guess is correct
	if strings.ToLower(guess) == sequence[len(sequence)-1] {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Incorrect. The final phrase was '%s'.\n", sequence[len(sequence)-1])
	}
}
