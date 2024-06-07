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

		for _, e := range entries {
			fmt.Println(e.Name)
		}

		fmt.Print(entries)
}