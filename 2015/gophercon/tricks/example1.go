package main

import "log"

//BEGIN OMIT
// example of an inline struct
var config struct { // HL
	RemoteHost string `env:"REMOTE_HOSTNAME"`
	Port       int    `env:"PORT"`
}

//END OMIT

// BEGIN1 OMIT
func main() {
	tests := []struct { // HL
		name       string // HL
		shouldFail bool   // HL
	}{{"blah", true}}

	for _, tc := range tests {
		if tc.shouldFail {
			log.Fatal("failed")
		}
	}
	// END1 OMIT
}
