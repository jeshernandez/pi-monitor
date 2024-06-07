package main

import (
	"fmt"
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
//word := []byte{}
breakLine := "\n"

tempereatureData := strings.Split(string(fileData[:]), breakLine)[1]
fmt.Print("tempData: ", tempereatureData)
temperature := tempereatureData[2:]
fmt.Print("Temperature", temperature)


// for _, data := range fileData {
//   if !bytes.Equal([]byte{data}, []byte(breakLine)) {
//     word = append(word, data)
//   } else {
//     fmt.Printf("ReadLine: %q\n", word)
//     word = word[:0]
//   }
// }
	
}