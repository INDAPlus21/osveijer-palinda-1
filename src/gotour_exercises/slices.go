package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		r := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			r[x] = uint8((x + y) / 2)
		}
		p[y] = r
	}
	return p
}

func main() {
	pic.Show(Pic)
}
