package main

import (
	"fmt"
	"regexp"
)

func DetectNumber(st string) bool {
	if match, err := regexp.MatchString("^-*[0-9]*[.e]?[0-9]+$", st); err != nil || !match {
		return false
	}

	return true
}

func main() {
	var st string

	fmt.Print("Enter your name: ")
	if _, err := fmt.Scanf("%s", &st); err == nil {
		if DetectNumber(st) {
			fmt.Println("You entered a number")
		} else {
			fmt.Println("You did not enter a number")
		}
	} else {
		fmt.Printf("Error happened reading input: %s\n", err.Error())
	}
}
