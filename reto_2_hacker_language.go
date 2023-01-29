/*
 * Escribe un programa que reciba un texto y transforme lenguaje natural a
 * "lenguaje hacker" (conocido realmente como "leet" o "1337"). Este lenguaje
 *  se caracteriza por sustituir caracteres alfanuméricos.
 * - Utiliza esta tabla (https://www.gamehouse.com/blog/leet-speak-cheat-sheet/)
 *   con el alfabeto y los números en "leet".
 *   (Usa la primera opción de cada transformación. Por ejemplo "4" para la "a")
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertionLeet(texto string) {
	// Convert text to lower
	lowerText := strings.ToLower(texto)

	// Declare map with the content of the alphabet
	alphabetLeet := map[string]string{
		"a": "4",
		"b": "I3",
		"c": "[",
		"d": ")",
		"e": "3",
		"f": "|=",
		"g": "&",
		"h": "#",
		"i": "1",
		"j": ",_|",
		"k": ">|",
		"l": "1",
		"m": "/\\/\\",
		"n": "^/",
		"o": "0",
		"p": "|*",
		"q": "(_,)",
		"r": "I2",
		"s": "5",
		"t": "7",
		"u": "(_)",
		"v": "\\/",
		"w": "\\/\\/",
		"x": "><",
		"y": "j",
		"z": "2",
	}

	// The loop will convert the text in alphabet leet
	for _, character := range lowerText {
		if font, ok := alphabetLeet[string(character)]; ok {
			fmt.Print(font)
		} else {
			fmt.Print(string(character))
		}
	}
}

func main() {
	fmt.Println("Enter the text to be converted: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	ConvertionLeet(text)
}
