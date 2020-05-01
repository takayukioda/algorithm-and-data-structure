package main

import (
	"errors"
	"fmt"
)

const SIZE = 5

// insert function will always remove last array element
// Array in go has no way to check "number of elements" and so we cannot return error because of "no space"
func insert(arr *[SIZE]string, idx int, val string) error {
	if idx >= SIZE || idx < 0 {
		return errors.New("Invalid index specified")
	}
	// Skip last index to avoid having range error
	for i := len(arr) - 2; i > idx; i-- {
		arr[i+1] = arr[i]
	}
	arr[idx] = val
	return nil
}

func del(arr *[SIZE]string, idx int) error {
	if idx >= SIZE || idx < 0 {
		return errors.New("Invalid index specified")
	}
	for i := idx; i < SIZE-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[SIZE-1] = ""
	return nil
}

func main() {
	var strarr [SIZE]string = [SIZE]string{"one", "two"}
	fmt.Println(strarr)
	if err := insert(&strarr, 4, "three"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(strarr)
	if err := del(&strarr, 3); err != nil {
		fmt.Println(err)
	}
	fmt.Println(strarr)
}
