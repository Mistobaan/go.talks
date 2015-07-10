package main

import "log"

// BEGIN OMIT
func Func1(a string) int {
	log.Println(a + "1")
	return 1
}

func Func2(a string) int {
	log.Println(a + "2")
	return 2
}

func Func3(a string) int {
	log.Println(a + "3")
	return 3
}

func main() {
	cmds := [](func(string) int){Func1, Func2, Func3} // HL
	total := 0
	for _, c := range cmds {
		total += c("->")
	}
	log.Println(total)
}

// END OMIT
