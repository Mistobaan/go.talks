package main

import "log"

// example of an inline struct
var config struct {
	RemoteHost string `env:"REMOTE_HOSTNAME"`
	Port       int    `env:"PORT"`
}

func main() {
	// Begin
	tests := []struct {
		name       string
		shouldFail bool
	}{"blah", true}
	// End
	for _, tc := range tests {
		if tc.shouldFail {
			log.Fatal("failed")
		}
	}
}
