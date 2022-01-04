package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func loadingBar() {
	time.Sleep(450 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(450 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(450 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(700 * time.Millisecond)
	fmt.Println()
	CallClear()
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

//the format for using this should preclude/exclude any Println or "\n" before usage
func optionsMenu(options ...string) string {
	//display menu portion
	fmt.Println("\n————————————————————")
	for i, s := range options {
		num := i+1
		fmt.Print(num)
		fmt.Println(". | " + s)
	}
	fmt.Println("(-)  Quit" +
		"\n————————————————————")

	//user input portion
	userChoice := scan()
	valid := false
	for !valid {
		for i, _ := range options {
			validChoice := strconv.Itoa(i+1) //increments index by 1 to be consistent with ui, then converts to string to be compared
			if userChoice == validChoice {
				valid = true //breaks out of loop
			}
		}
		if !valid {
			fmt.Println("Invalid input. Please try again.") //default case if user input is not a valid option
			userChoice = scan()
		}
	}
	return userChoice
}