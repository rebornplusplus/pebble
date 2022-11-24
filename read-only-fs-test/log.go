package main

import (
	"log"
	"os"
	"time"
	"strconv"
	"fmt"
)

func main() {
	fmt.Println("Opening/Creating file ..")
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0
	for counter < 10 {
		counter += 1
		fmt.Println("Writing", counter, "..")
		_, err = f.WriteString(strconv.FormatInt(int64(counter), 10) + "\n")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}

	fmt.Println("Closing file ..")
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
