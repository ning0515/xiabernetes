package util

import (
	"fmt"
	"time"
)

func HandleCrash() {
	r := recover()
	if r != nil {
		fmt.Printf("Recovered from panic: %#v", r)
	}
}

// Loops forever running f every d.  Catches any panics, and keeps going.
func Forever(f func(), period time.Duration) {
	for {
		func() {
			defer HandleCrash()
			f()
		}()
		time.Sleep(period)
	}
}
