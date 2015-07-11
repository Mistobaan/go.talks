package main

import "log"

func info2(namespace string, args ...string) { // HL
	log.Printf("%s:%#v\n", namespace, args)
}

func info1(info string, args ...string) {
	info2("[database]", append([]string{"more info", "database name"}, args...)...) // HL
}

func main() {
	info1("test", "a", "b", "c", "d")
}
