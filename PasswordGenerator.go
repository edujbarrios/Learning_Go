package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	var password strings.Builder
	for i := 0; i < passwordLength; i++ {
		password.WriteRune(rune(rand.Intn(94) + 33))
	}

	fmt.Println(password.String())
}
