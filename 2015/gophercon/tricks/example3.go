package main

import (
	"fmt"
	"log"
)

// BEGIN OMIT
func willFail() {
	panic(fmt.Errorf("told you"))
}

func cannotFail() (err error) { // HL
	defer func() {
		if errRecover := recover(); errRecover != nil {
			switch v := errRecover.(type) {
			case string:
				err = fmt.Errorf(v)
			case error:
				err = v // HL
			default:
				err = fmt.Errorf("received unknown value inside panic: %#v", v)
			}
		}
	}()

	willFail()
	return nil
}

// END OMIT

func main() {
	log.Println(cannotFail())
}
