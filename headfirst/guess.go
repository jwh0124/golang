// guess program 은 플레이어가 난수를 맞히는 게임입니다.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	success := false

	for i := 0; i < 10; i++ {
		fmt.Println("You have", 10-i, "guesses left.")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry. You didn`t guess my number. It was: [", target, "]")
	}
}
