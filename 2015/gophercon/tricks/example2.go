package main

import "log"

// BEGIN3 OMIT
type MyStruct struct {
	secret string
}

func (m *MyStruct) GetSecret() string {
	return m.secret
}

// END3 OMIT

func main() {

	// BEGIN1 OMIT
	m := &MyStruct{"secret discovered"}
	if canGetSecret, ok := (interface{})(m).(interface { // HL
		GetSecret() string // HL
	}); ok { // HL
		log.Println(canGetSecret.GetSecret())
	}
	// END1 OMIT

	// BEGIN2 OMIT
	var i interface{} = &MyStruct{"secret discovered"}
	if canGetSecret, ok := i.(interface { // HL
		GetSecret() string // HL
	}); ok { // HL
		log.Println(canGetSecret.GetSecret())
	}
	// END2 OMIT
}
