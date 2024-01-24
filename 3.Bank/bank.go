package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetBalanceFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())
	for {
		presentOptions()

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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
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
