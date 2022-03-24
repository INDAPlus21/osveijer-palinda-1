package main

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	out := 0
	for _, n := range a {
		out += n
	}
	res <- out
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int, 2)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	s := 0
	s += <-ch
	s += <-ch

	return s
}
