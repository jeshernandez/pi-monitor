package main

import (
	"bytes"
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
	path := "/sys/bus/w1/devices/"
	ds18b20, err := VerifyPathExists(path)
	if err != nil {
		log.Fatal(err) 
	}
	ReadLine(path + ds18b20 + "/w1_slave")

}

func ReadLine(fileName string) {
    fileData, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
word := []byte{}
breakLine := "\n"

for _, data := range fileData {
  if !bytes.Equal([]byte{data}, []byte(breakLine)) {
    word = append(word, data)
  } else {
    fmt.Printf("ReadLine: %q\n", word)
    word = word[2:]
  }
}	
}