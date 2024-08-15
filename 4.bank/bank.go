package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("can't continue, sorry.")
	}

	fmt.Println("Wecome to Go bank")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated ! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated ! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("GoodBye!")
			fmt.Println("Thanks for choosing our bank")
			return
			// 	break
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated ! New amount:", accountBalance)

		// } else if choice == 3 {
		// 	fmt.Print("Withdrawal amount: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)cc

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance updated ! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("GoodBye!")
		// 	break
		// }
	}

	// fmt.Println("Thanks for choosing our bank")
}
