/*
* Escribe un programa que detecte si una palabra ingresada es un palíndromo.
* Devuelve TRUE o FALSE dependiendo del caso.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Palindrome(word string) bool {
	reverserdWord := ""

	for _, v := range word {
		reverserdWord = string(v) + reverserdWord
	}

	return word == reverserdWord
}

func main() {
	fmt.Print("Enter the text to be validated: ")
	reader := bufio.NewReader(os.Stdin)
	enter, _ := reader.ReadString('\n')
	text := strings.TrimRight(enter, "\n\r")

	fmt.Println(Palindrome(text))
}
