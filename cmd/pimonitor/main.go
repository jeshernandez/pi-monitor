package main

import (
	"fmt"
	"log"
	"os"
)

func VerifyPathExists(path string) (string, error) {
    entries, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range entries {
			return e.Name(), nil
 		}
	return "error", nil 
}

func main() {
	ds18b20, err := VerifyPathExists("/sys/bus/w1/devices")
	if err != nil {
		log.Fatal("Could not capture")
  }

	fmt.Println(ds18b20)
	//if err != nil { fmt.Print("Impossible, sensor is not configured correctly!") }


	
}