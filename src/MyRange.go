package main

func main() {
	println("Hello, this is dog.")

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)  // == 2**i
	}

	for i, value := range pow {
		println(i, value)
	}

}
