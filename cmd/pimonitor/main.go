package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	path := "/sys/bus/w1/devices"
	ds18b20, err := VerifyPathExists(path)
	location := path + ds18b20 + "/w1_slave"
	fmt.Println("Location", location)

	fileIO, err := os.OpenFile(location, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		if i == 2 {
			fmt.Println(line)
		}
	}

	
}