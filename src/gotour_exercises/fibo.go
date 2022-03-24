package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i1 := 0
	i2 := 1
	return func() int {
		temp1 := i1
		temp2 := i2
		i2 = i1 + i2
		i1 = temp2
		return temp1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
