package jobs

import (
	"fmt"
)

type Email struct {
	Data interface{}
}

func (e *Email) Handle() {
	fmt.Printf("processing email job")
}
