package main

import (
	"fmt"
	"os"
)

var (
	user Account
	defaultAccount = Account{"default", "123", "america","john doe", false}
	testAccount1 = Account{"acc1", "vanilla", "narnia", "Prince Caspian", false}
	testAccount2 = Account{"acc2", "chocolate", "Texas", "Nuha", false}
	database = []Account{defaultAccount, testAccount1, testAccount2}
)

type Account struct {
	username, password, location, name string
	loggedIn bool
}

type Citizen struct {
	ssn int
	Account
}

type Politician struct {
	position, key string
	Account
}

type Judge struct {
	court string
	Account
}

// collects user input to log them in
// using the retrieveAccount function
func logIn () Account {
	loggedIn := false
	var theirAccount Account

	for !loggedIn {
		fmt.Println("————————————————————")
		fmt.Print("Enter your username: ")
		un := scan()
		fmt.Print("Enter your password: ")
		pw := scan()

		attempt, found := retrieveAccount(un, pw)

		if !found {
			fmt.Println("That didn't seem to work. Please try again.")
		} else {
			loggedIn = true
			theirAccount = attempt
		}
	}
	fmt.Println("Welcome!")
	return theirAccount
}

func createAccount () Account {
	CallClear()
	fmt.Print("Creating an account is fast & easy!" +
		"\nFirst up, what type of account would you like to create?")
	optionsMenu("Citizen Account",
		"Legislative Account",
		"Judicial Account",
		"(coming soon) Executive Account",
		"(coming soon) Integrative Account")
	choice := scan()

	for choice != "-" && choice != "1" && choice != "2" && choice != "3" {
		fmt.Println("Invalid input. Please try again.")
		choice = scan()
	}

	CallClear()
	var un, pw, loc, n string
	if choice == "1"{
		fmt.Println("Cool! I just need to get some information & then we should be set!")
		fmt.Print("Please enter a memorable username: ")
		un = scan()
		fmt.Print("Please enter a strong password: ")
		pw = scan()
		fmt.Print("Please enter your legal residence: ")
		loc = scan()
		fmt.Print("Please enter your legal name: ")
		n = scan()
	} else if choice == "2" {
		fmt.Println("wip")
	} else if choice == "3" {
		fmt.Println("wip")
	} else {
		os.Exit(0)
	}

	return Account{un, pw, loc, n, true}
}

func storeAccount (account Account) {
	// sort by alphabetical order
	// then later by frequency of login
	database = append(database, account)
}

// loops through database, checking for a match using given un & pw
// if found returns account, else returns default account
func retrieveAccount (username, password string) (Account, bool) {
	for i, s := range database {
		if s.username == username && s.password == password {
			return database[i], true
		}
	}
	return database[0], false
}

//combine with retrieveAccount
func authenticatePW () {

}

func (user Account) accountActions () {
	fmt.Println("Hi, " + user.name +
		"\nWhat would you like to do today?" +
		"\n1. | Let's Vote!" +
		"\n2. | See Tally History" +
		"\n3. | Account Settings")

}

func changePW () {

}

func logOut () {

}




















