package main

import "log"

// BEGIN OMIT
type MyInterface interface {
	WAT() string
}

type MyImplementation struct {
}

func (_ *MyImplementation) WAT() string {
	return "WAT"
}

var _ MyInterface = (*MyImplementation)(nil) // HL

func main() {
	m := &MyImplementation{}
	log.Println(m.WAT())
}

// END OMIT
