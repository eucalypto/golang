package main

import (
	"fmt"
)

func pointers() {
	var intx int = 1
	var inty int
	var intpointer *int // intpointer is a pointer for int

	intpointer = &intx
	inty = 43

	*intpointer = 42

	//println(intpointer)
	//println(intx)
	println(*intpointer)
	//print(*intpointer)
	println(inty)


	pointer2 := new(int)
	*pointer2 = 3
	println(*pointer2)
}

func ints() {
	var i int
	fmt.Printf("%T", i)
}

func runes() {
	//rune1 := "1"
	//println(unicode.IsDigit(rune(rune1)))
}

func iotas() {
	// iota is the equivalent of enums
	// they define distinct (but somewhat arbitrary) constants
	type Weekdays int
	const (
		Mon Weekdays = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)

	println(Mon, Tue, Wed)

	var day Weekdays
	//day := Mon
	day = Tue

	if day == Tue {
		println("Tuesday!")
	} else {
		println("Not tuesday!")
	}
}

func controlFlow() {
	condition := true
	if condition {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}


	for i := 1; i <= 3; i++ {
		fmt.Println(i, "Mississippi")
	}

	counter := 0
	for counter < 4 {
		fmt.Println(counter)
		counter++
	}

	counter2 := 99
	for {
		if counter2 < 96 {
			fmt.Println("hello")
			break
		}
		fmt.Println(counter2, "bottles of beer on the wall.")
		counter2--
	}


	x := 3
	switch x {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}


	// Tagless switch with boolean expressions
	// the first expression that evaluates to true will be chosen
	weight := 101
	age := 53
	switch {
	case weight > 100:
		fmt.Println("You should eat less sugar.")
	case weight < 80:
		fmt.Println("You look good; have you lost weight?")
	case age > 30:
		fmt.Println("hello old man")
	case age < 20:
		fmt.Println("Welcome, youngling.")
	default:
		fmt.Println("Nothing to say.")
	}
}

func userInput() {
	var slicesOfBacon int
	fmt.Print("How many slices of bacon do you want?: ")
	fmt.Scan(&slicesOfBacon)

	fmt.Println("you asked for", slicesOfBacon, "slices of bacon.")
	
}

func main() {
	//pointers()
	//ints()
	//iotas()
	//controlFlow()
	userInput()
}
