package main

import (
	"fmt"
	"log"
	"regexp"
	"io/ioutil"
	"os"
)

func main(){
	takeContent, err := ioutil.ReadFile("/home/ryazanov/Downloads/tdlfFile.txt")

	if err != nil {
		log.Fatal(err)
	}

	str1 := string(takeContent)

	pRe := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", pRe.String()) // print pattern
	fmt.Println(pRe.MatchString(str1))        // true

	createLogFile, err := os.Create("testLog.log")
	if err != nil {
		panic(err)
	}
	defer createLogFile.Close()

	submatchall := pRe.FindAllString(str1, -1)
	for _, element := range submatchall {
		//fmt.Println(element, "\n")
		_, err = createLogFile.WriteString(" ")
		_, err = createLogFile.WriteString(element)
	}

	// На этом этапе всё работает правильно.
	// Мы фильтруем данные, которые получаем из файла tdlfFile.txt,
	// и помещаем отфильтрованные IP-адреса в созданный файл testLog.log
	// Всё протестировано и всё работает нормально.
	// Теперь необходимо присвоить первый IP из файла testLog.log переменной

	takeFirstIp := pRe.FindString(str1)
	
}