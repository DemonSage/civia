package main

import (
	"fmt"
	"os"
)

func main() {
	//a := app.New()
	//w := a.NewWindow("Hello")

	//w.SetContent(widget.NewLabel("Hello World!"))
	CallClear()
	startUp()
	loadingBar()

	fmt.Println(user)
	fmt.Println(database)

}

func startUp () {
	fmt.Print("Welcome to Civia!" +	"\nWhat would you like to do?")

	choice := optionsMenu("Log In", "Create Account", "See Results", "HOPE")
	switch choice {
	case "1":
		user = logIn()
		user.loggedIn = true
	case "2":
		user = createAccount()
		storeAccount(user)
	case "3":
		CallClear()
		fmt.Print("wip")
		os.Exit(3)
	case "4":
		CallClear()
		fmt.Println("**ADMIN ONLY**\nVersion 0.1 Pre-Alpha")
		os.Exit(-1)
	default:
		os.Exit(0)
	}
}

func seePublicResults () {

}

