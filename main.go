package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)

func balance_check(balance float64) {
	fmt.Print("Your current balance is $")
	fmt.Print(balance, "\n")
}

//Uses pointer to access balance in memory
func withdraw(balance *float64) {
	reader := bufio.NewReader(os.Stdin)
	var difference float64
	fmt.Println("How much money would you like to withdraw")
	input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
	//Converts string to float
	inputNum, _ := strconv.ParseFloat(input, 64)

	if (*balance == 0.0) {
		fmt.Println("You currently don't have a balance, please deposit money if you wish to make a withdraw")
	} else if (inputNum > *balance) {
		difference = inputNum - *balance
		failure := "Unable to complete transaction because the requested withdraw amount is $" + strconv.FormatFloat(difference, 'E', -1, 64) + " higher than your available balance"
		fmt.Println(failure)
		fmt.Println("If you wish to make a withdraw please select an amount lower than $")
		fmt.Print(balance)
	} else {
		*balance -= inputNum
		fmt.Print("$")
		fmt.Print(input)
		fmt.Println(" has been withdrawn from your account!")
	}

}

func deposit(balance *float64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How much would you like to deposit")
	input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
	//Fix input problem removing input spaces
	input = strings.TrimSpace(input)
	//Convert string to float
	inputNum, _ := strconv.ParseFloat(input, 32)
	*balance += inputNum
	fmt.Print("$", inputNum)
	fmt.Println(" has been added to your account!")
}

//Does the user want to make another transaction?
func ending_check(end bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Thank you for your transaction!\nWould you like to make another? Please type 'y' or 'n'\n")
	input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
	//Weird input fix
	runes := []rune(input)
	input = string(runes[0])

	if (input == "y" || input == "Y") {
		end = false
	} else if (input == "n" || input == "N") {
		end = true
	} else {
		fmt.Println("Invalid value, please try again")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	end := false
	var balance float64 = 0.00

	for !end  {
		//Check prompt
		fmt.Println("What would you like to do")
		fmt.Print("Type 'b' for a balance check\nType 'w' for a withdraw\nType 'd' for deposit\nOr press 'q' to quit the program\n")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		//Weird input fix
		runes := []rune(input)
		input = string(runes[0])

		switch input {
		case "b":
			balance_check(balance)
			ending_check(end)
			break
		case "w":
			withdraw(&balance)
			ending_check(end)
			break
		case "d":
			deposit(&balance)
			ending_check(end)
			break
		case "q":
			end = true
			break
		default:
			fmt.Println("Please insert a valid option listed above")
			break
		}
	}

	fmt.Print("Thanks for using this program!")
}