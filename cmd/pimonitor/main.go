package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Attemping to capture temperature...")
	  entries, err := os.ReadDir("/sys/bus/w1/devices")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Print(entries)
}