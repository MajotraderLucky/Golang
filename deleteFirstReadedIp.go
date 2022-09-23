package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
)

func main() {
	content, err := ioutil.ReadFile("/home/ryazanov/MyCode/GoCode/MajorGo/Monitoring/onlyIP.txt")

	if err != nil {
		log.Fatal(err)
	}
	FirstIpInString := string(content)

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String())     // print pattern
	fmt.Println(re.MatchString(FirstIpInString)) // true

	takeFirstIp := re.FindString(FirstIpInString)
	myFirstIp := takeFirstIp
	fmt.Println("Первый IP, с которым работаем - ", myFirstIp, " - присвоен переменной myFirstIp")
	typeMyFirstIp := reflect.TypeOf(myFirstIp)
	fmt.Println("Тип переменной myFirstIp - ", typeMyFirstIp)


	// Посчитаем количество символов в переменной myFirstIp
	howManyLetters := len(myFirstIp)
	fmt.Println("Символов в первом IP - ", howManyLetters, "присваивается переменной howManyLetters")

	// Проверим тип переменной howManyLetters
	fmt.Println("Переменная howManyLetters имеет тип - ", reflect.TypeOf(howManyLetters))
	howManyLetters = howManyLetters + 2

	// Удалим количество символов, которое соответствует переменной howManyLetters из файла onlyIP.txt
	subContent := content[howManyLetters:]
	//fmt.Print(subContent)

	// Создаём новый файл и записываем туда значение переменной subContent
	changedFile, err := os.Create("changedFile.txt")
	if err != nil {
		fmt.Println("Unable to create file", err)
		os.Exit(1)
	}

	defer changedFile.Close()
	changedFile.WriteString(string(subContent))

	fmt.Println("Создаём файл с удалённым первым обработанным IP changedFile.txt")
	fmt.Println("Создаём массив, в который будем заносить Ip-адреса из списка")

	var ipAdressArr [500]string
	var countOfArr int = 0 // Счётчик массива ipAdressArr
	ipAdressArr[countOfArr] = myFirstIp
	fmt.Println("В массиве ipAdressArr появился элемент -", ipAdressArr[countOfArr])
	countOfArr++
	fmt.Println("Счётчик countOfArr увеличился и соответствует значению -", countOfArr)
}
