package main

import (
	"log"
	"strconv"
)

func main() {

	// Create an array that holds exactly four ints
	var a [2][2]string
	log.Println(a, len(a), cap(a), len(a[0]))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			a[i][j] = strconv.Itoa(i + j)
		}
	}
	log.Println(a, len(a), cap(a))
}
