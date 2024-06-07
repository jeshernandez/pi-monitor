package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	_, faren := GetSensorTemperature(path + ds18b20 + "/w1_slave")

	fmt.Print("Temp: ", faren)	

}

func GetSensorTemperature(fileName string)(float64, float64) {
    fileData, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
breakLine := "\n"

secondLine := strings.Split(string(fileData[:]), breakLine)[1]
//fmt.Print("SecondLine: ", secondLine)
temperatureData := strings.Split(secondLine, " ")[9]
//fmt.Print("Temperature Data", temperatureData)
temperature, _ := strconv.ParseFloat(temperatureData[2:], 64)
fmt.Print("Temperature", temperature)

celcius1 := (temperature / 1000)

farenheit1 := (celcius * 1.8) + 32

return celcius1, farenheit1

}