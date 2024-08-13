package gocode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Address Addressing

func CallAddressAPI() {
	jsonFile, err := os.Open("geojson/kenyaMinerals.geojson")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Kenya Minerals.geojson")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Address)

}

var Zone Zoning

func CallZoningAPI() {

	jsonFile, err := os.Open("geojson/kenya.geojson")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened kenya.geojson")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Zone)

}