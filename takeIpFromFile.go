package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	content, err := ioutil.ReadFile("/home/ryazanov/Downloads/tdlfFile.txt")

	if err != nil {
		log.Fatal(err)
	}

	str1 := string(content)

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}

	// Теперь необходимо записать полученные данные в новый файл
	f, err := os.Create("onlyIP.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, element2 := range submatchall {
		_, err = f.WriteString(" ") // Благодаря этой строке появляется пробел между IP-адресами в файле onlyIP.txt
		_, err = f.WriteString(element2)
	}
}	
