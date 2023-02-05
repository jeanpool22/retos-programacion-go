/*
	Dado un array de número enteros:
	1. Ordenar el array de menor a mayor
	2. Contar los números que faltan entre la primera y útima posición del array
*/

package main

import (
	"fmt"
	"sort"
)

func Consecutive(statues []int) int {
	sort.Ints(statues)

	var count int
	firstPosition := statues[0]
	lastPosition := statues[len(statues)-1]

	for i := firstPosition; i <= lastPosition; i++ {
		result := Search(i, statues)
		if result == false {
			count++
		}
	}

	return count
}

func Search(num int, statues []int) bool {
	for _, v := range statues {
		if num == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Consecutive([]int{6, 2, 3, 8}))
	fmt.Println(Consecutive([]int{0, 3}))
	fmt.Println(Consecutive([]int{5, 4, 6}))
	fmt.Println(Consecutive([]int{6, 3}))
	fmt.Println(Consecutive([]int{1}))
}
