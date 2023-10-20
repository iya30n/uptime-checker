package jobs

import (
	"fmt"
)

type Email struct {
	Data interface{}
}

func (e *Email) Handle() bool {
	// TODO: handle should return a bool to check for retry
	fmt.Printf("processing email job")
	return true
}
