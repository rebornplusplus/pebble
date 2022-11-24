package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 20 {
		fmt.Println("hello")
		time.Sleep(time.Second / 2)
		i += 1
	}
}
