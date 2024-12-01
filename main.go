package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read file.")	
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64)  {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main()  {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
	}
	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scanln(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount // same with accountBalance = accountBalance + depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}

	// 	if choice == 1 {
	// 		fmt.Println("Your balance is: ", accountBalance)
	// 	} else if choice == 2 {
	// 		fmt.Print("Your deposit: ")
	// 		var depositAmount float64
	// 		fmt.Scanln(&depositAmount)
	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid deposit amount. Must be greater than 0.")
	// 			// return
	// 			continue
	// 		}
	// 		accountBalance += depositAmount // same with accountBalance = accountBalance + depositAmount
	// 		fmt.Println("Your new balance is: ", accountBalance)
	// 	} else if choice == 3 {
	// 		fmt.Print("Your withdraw: ")
	// 		var withdrawAmount float64
	// 		fmt.Scanln(&withdrawAmount)
	// 		if withdrawAmount <= 0 {
	// 			fmt.Println("Invalid withdraw amount. Must be greater than 0.")
	// 			return
	// 		}
	// 		if withdrawAmount > accountBalance {
	// 			fmt.Println("Insufficient balance.")
	// 			return
	// 		}

	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("Your new balance is: ", accountBalance)
	// 	} else {
	// 		fmt.Println("Goodbye!")
	// 		// return
	// 		break
	// 	}		
	// }
	// 	fmt.Println("Thanks for choosing our bank")

}