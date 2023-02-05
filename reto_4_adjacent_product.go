/*
Dado un array de elementos numericos enteros, encontrar el mayor producto
entre el número adyacente del arrary dado
*/

package main

import "fmt"

func solution(inputArray []int) int {
	var product []int

	for i, v := range inputArray {
		product = append(product, v*(inputArray[i+1]))

		if len(product) == len(inputArray)-1 {
			break
		}
	}

	major := product[0]

	for _, v := range product {
		if major > v {
			major = major
		} else {
			major = v
		}
	}

	return major
}

func main() {
	fmt.Println(solution([]int{3, 6, -2, -5, 7, 3}))
	fmt.Println(solution([]int{-1, -2}))
	fmt.Println(solution([]int{5, 1, 2, 3, 1, 4}))
	fmt.Println(solution([]int{1, 2, 3, 0}))
	fmt.Println(solution([]int{9, 5, 10, 2, 24, -1, -48}))
	fmt.Println(solution([]int{5, 6, -4, 2, 3, 2, -23}))
	fmt.Println(solution([]int{4, 1, 2, 3, 1, 5}))
	fmt.Println(solution([]int{-23, 4, -3, 8, -12}))
}
