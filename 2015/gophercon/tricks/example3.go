package main

import (
	"fmt"
	"log"
)

func willFail() {
	panic(fmt.Errorf("told you"))
}

func cannotFail() (err error) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			switch v := errRecover.(type) {
			case string:
				err = fmt.Errorf(v)
			case error:
				err = v
			default:
				err = fmt.Errorf("received unknown value inside panic: %#v", v)
			}
		}
	}()

	willFail()
	return nil
}

func main() {
	log.Println(cannotFail())
}
