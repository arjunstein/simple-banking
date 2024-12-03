package main

import (
	"fmt"
	"simple-banking/fileops"
)

const accountBalanceFile = "balance.txt"

func main()  {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		// panic("Can't continue process")
	}
	for {		
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}