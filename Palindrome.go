package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main()  {
	var char string
	fmt.Println("Character:")
	fmt.Scanln(&char)
	if isPalindrome(char){
		fmt.Println(char, "Adalah palindrome")
	} else {
		fmt.Println(char, "Bukan palindrome")
	}
}