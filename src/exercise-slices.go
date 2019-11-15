package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var ret [][]uint8
	for i := 0; i < dx; i++ {
		var xSlice []uint8
		for j := 0; j < dy; j++ {
			xSlice = append(xSlice, uint8(i*j))
		}
		ret = append(ret, xSlice)
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
