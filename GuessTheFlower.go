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

	// Create a slice of flowers
	flowers := []string{"rose", "tulip", "daisy", "lily", "sunflower"}

	// Choose a random flower
	flower := flowers[rand.Intn(len(flowers))]

	// Prompt the user to guess the flower
	fmt.Print("Guess the flower: ")

	// Read the user's guess from the command line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	guess := scanner.Text()

	// Check if the guess is correct
	if strings.ToLower(guess) == flower {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Incorrect. The flower was a %s.\n", flower)
	}
}
