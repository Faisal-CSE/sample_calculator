package sample_calculator

// I'm creating a simple calculator that
// performs one calculator operation as per the
// user's choice. For readability of code,
// I named the package as "calculator"
// And remember, the first executable line
// must always be as mentioned above:
// the keyword package followed by a name
// that you wish to give to your package*
//* indicates very very important

import (
	"fmt"
	"log"
)

// importing fmt package for basic
// printing & scan operations

func displayOptionList() {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("******************************************************")
}

func Calc() {

	// a simple Calc function that contains
	// all code within and has no return
	// type mentioned
	// Println prints the input string in new line
	fmt.Println("Welcome to calculator app")
	fmt.Println("********************MAIN MENU*************************")
	displayOptionList()

	var choice int

	// choice will store the user's
	// input as per the menu shown above
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Println("Scan error: ", err)
		return
	}
	
	if choice == 5 {
		fmt.Println("You clicked on exit.")
		fmt.Println("******************************************************")
		fmt.Println("Thank you for using calculator! Have a nice day.")
		fmt.Println("******************************************************")
		return
	}
	var a, b int

	// After the choice of operation, user
	// will be asked to enter 2 int
	// values one by one to perform
	// the operation on
	fmt.Println("Enter value of a: ")
	_, err1 := fmt.Scan(&a)
	if err1 != nil {
		log.Println("Scan error1: ", err)
		return
	}
	fmt.Println("Enter value of b: ")
	_, err2 := fmt.Scan(&b)
	if err2 != nil {
		log.Println("Scan error2: ", err)
		return
	}
	if choice == 1 {
		// choice 1 activates this part --> addition
		ans := a + b
		fmt.Println("Answer = ", ans)
	} else if choice == 2 {
		// choice 2 activates this part --> subtraction
		ans := a - b
		fmt.Println("Answer = ", ans)
	} else if choice == 3 {
		// choice 3 activates this part --> multiplication
		ans := a * b
		fmt.Println("Answer = ", ans)
	} else if choice == 4 {
		// choice 4 activates this part --> division
		// remember not to enter second value as 0
		// as that would raise a DivideByZero error
		// or may display infinity
		ans := a / b
		fmt.Println("Answer = ", ans)
	} else {
		fmt.Println("You clicked Invalid key.")
		fmt.Println("******************************************************")
		fmt.Println("Thank you for using calculator! Have a nice day.")
		fmt.Println("******************************************************")
		return
	}

	fmt.Println("******************************************************")
	fmt.Println("Thank you for using calculator! Have a nice day.")
	fmt.Println("******************************************************")
}
