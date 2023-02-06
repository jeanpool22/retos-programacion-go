/*
Dada una matriz de n * n:
1. Encontrar la suma de la primera diagonal
2. Encontrar la suma de la segunda diagonal
3. Restar el resultado de la primera diagonal a la segunda diagonal
4. Devolver el valor absoluto de la resta anterior
*/
package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int) int {
	var sumA int
	var sumB int
	var result int

	for i, _ := range arr {
		for j, _ := range arr[i] {
			if i == j {
				sumA += arr[i][j]
			}
			if i+j == len(arr)-1 {
				sumB += arr[i][j]
			}
		}
	}

	result = int(math.Abs(float64(sumA) - float64(sumB)))
	return result
}

func main() {
	fmt.Println(diagonalDifference([][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}
