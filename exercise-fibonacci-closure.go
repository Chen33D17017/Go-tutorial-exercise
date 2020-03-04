package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	f1, f2 := 0, 1
	return func() int{
		if count == 0 {
			count += 1
			return 0
		} else {
			f1, f2 = f2, f1+f2
			count += 1
			return f1
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
