package main

import (
	"fmt"
)

func deleteItem[T any](index int, slice []T) []T {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
	var (
		slice = []int{1,2,3,4,5,6}
	)
	
	for _, v := range slice {
		fmt.Println(v)
	}
	
	var result = deleteItem(1, slice)

	fmt.Println("after deleting the itme: ")
	
	for _, v := range result {
		fmt.Println(v)
	}
}