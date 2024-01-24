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
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644) // file permission 0644 (read and write permission)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				println("Invalid amount.Must be greater than zero!")
				// return
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				println("Invalid amount.Must be greater than zero!")
				// return
				continue
			}
			if withdrawAmount > accountBalance {
				println("Invalid amount. You can't withdraw more than account balance")
				// return
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing us!")
			// cannot use break here because break is reserve work in the switch statement
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		println("Invalid amount.Must be greater than zero!")
		// 		// return
		// 		continue
		// 	}
		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdraw amount: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 {
		// 		println("Invalid amount.Must be greater than zero!")
		// 		// return
		// 		continue
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		println("Invalid amount. You can't withdraw more than account balance")
		// 		// return
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	// return
		// 	break
		// }
	}

	// fmt.Println("Thanks for choosing us!")
}
