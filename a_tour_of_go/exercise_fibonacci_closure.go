package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var slice []int
	return func() int {
		count := len(slice)
		if count < 2{
			slice = append(slice, count)
		} else{
			slice = append(slice, slice[count-2] + slice[count - 1])
		}
		return slice[count]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
