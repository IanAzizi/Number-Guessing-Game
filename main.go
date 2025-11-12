package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("welcome to Random Guess number")
	PlayGame()

}

func PlayGame() {
	// Game logic goes here
	rand.Seed(time.Now().UnixNano())
	SecretNumber := rand.Intn(100)
	var Chance int
	var Guess int
	var level int
	fmt.Println("choose your level")
	fmt.Scanln(&level)
	switch level {
	case 1:
		defer fmt.Printf("the secret number was %v\n", SecretNumber)
		Chance = 10
		fmt.Println("you chosse easy level")
		fmt.Println("you have 10 chances to guess the number : ")
		for i := 0; i < Chance; i++ {

			fmt.Scanln(&Guess)
			if Guess == SecretNumber {
				fmt.Println("your guess was right the secret number was ", SecretNumber)
				break
			} else {
				fmt.Printf("wrong guess you only have %v chances left\n", Chance-i-1)
				if Guess < SecretNumber {
					fmt.Printf("the Secret Number is higher then %v\n", Guess)
				} else {
					fmt.Printf("the Secret Number is lower then %v\n", Guess)
				}

			}

		}
	case 2:
		defer fmt.Printf("the secret number was %v\n", SecretNumber)
		Chance = 5
		fmt.Println("you chosse medium level")
		fmt.Println("you have 5 chances to guess the number : ")
		for i := 0; i < Chance; i++ {
			fmt.Scanln(&Guess)
			if Guess == SecretNumber {
				fmt.Println("your guess was right the secret number was ", SecretNumber)
				break
			} else {

				fmt.Printf("wrong guess you only have %v chances left\n", Chance-i-1)
				if Guess < SecretNumber {
					fmt.Printf("the Secret Number is higher then %v\n", Guess)
				} else {
					fmt.Printf("the Secret Number is lower then %v\n", Guess)
				}
			}

		}
	case 3:
		defer fmt.Printf("the secret number was %v\n", SecretNumber)
		Chance = 3
		fmt.Println("you chosse hard level")
		fmt.Println("you have 3 chances to guess the number : ")
		for i := 0; i < Chance; i++ {
			fmt.Scanln(&Guess)
			if Guess == SecretNumber {
				fmt.Println("your guess was right the secret number was ", SecretNumber)
				break
			} else {

				fmt.Printf("wrong guess you only have %v chances left\n", Chance-i-1)
				if Guess < SecretNumber {
					fmt.Printf("the Secret Number is higher then %v\n", Guess)
				} else {
					fmt.Printf("the Secret Number is lower then %v\n", Guess)
				}

			}

		}
	}
}
