package main

import "fmt"

func main() {
	var string_to_reverse string
	fmt.Print("Enter a string to reverse: ")
	fmt.Scan(&string_to_reverse)
	string_to_reverse_slice := []rune(string_to_reverse)
	reversed_string := ""
	for i := len(string_to_reverse_slice) - 1; i >= 0; i-- {
		reversed_string += string(string_to_reverse_slice[i])
	}
	fmt.Printf("String %v reversed = %v\n", string_to_reverse, reversed_string)
}
