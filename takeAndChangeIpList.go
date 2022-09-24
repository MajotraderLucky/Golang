package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	startContent, err := ioutil.ReadFile("/home/ryazanov/Downloads/tdlfFile.txt")

	if err != nil {
		log.Fatal(err)
	}

	str1 := string(startContent)

	pRe := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", pRe.String()) // print pattern
	fmt.Println(pRe.MatchString(str1))        // true

	submatchall := pRe.FindAllString(str1, -1)
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

	// Конец создания контента для файла onlyIP.txt
	// Далее идёт код из файла deleteFirstReadedIp.go

	content, err := ioutil.ReadFile("/home/ryazanov/MyCode/GoCode/MajorGo/Monitoring/onlyIP.txt")

	if err != nil {
		log.Fatal(err)
	}
	FirstIpInString := string(content)

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String())     // print pattern
	fmt.Println(re.MatchString(FirstIpInString)) // true
}