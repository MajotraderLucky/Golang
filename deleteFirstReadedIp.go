package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	content, err := ioutil.ReadFile("/home/ryazanov/MyCode/GoCode/Monitoring/Golang/onlyIP.txt")

	if err != nil {
		log.Fatal(err)
	}
	FirstIpInString := string(content)

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String())     // print pattern
	fmt.Println(re.MatchString(FirstIpInString)) // true

	takeFirstIp := re.FindString(FirstIpInString)
	myFirstIp := takeFirstIp
	fmt.Println(myFirstIp)
}
