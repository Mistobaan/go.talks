package main

import "log"

type MyStruct struct {
	secret string
}

func (m *MyStruct) GetSecret() string {
	return m.secret
}

func main() {

	m := &MyStruct{"secret discovered"}

	if canGetSecret, ok := (interface{})(m).(interface {
		GetSecret() string
	}); ok {
		log.Println(canGetSecret.GetSecret())
	}

	// cleaner ?
	var i interface{} = &MyStruct{"secret discovered"}
	if canGetSecret, ok := i.(interface {
		GetSecret() string
	}); ok {
		log.Println(canGetSecret.GetSecret())
	}

}
