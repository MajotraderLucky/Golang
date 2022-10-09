package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func main() {
	takeContent, err := ioutil.ReadFile("/home/sergey/Downloads/tdlfFile.txt")

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
	// Теперь надо создать переменную и присвоить ей строку IP-адресов из
	// файла testLog.log

	ipList, err := ioutil.ReadFile("/home/sergey/MyCode/GoCode/Golang/testLog.log")

	if err != nil {
		log.Fatal(err)
	}
	listIpInString := string(ipList)
	fmt.Println(pRe.MatchString(listIpInString)) // true

	takeFirstIp := pRe.FindString(listIpInString)
	fmt.Println(takeFirstIp, " - Выводим переменную takeFirstIp")

	// Сначала необходимо создать массив данных и присвоить ему первый IP
	// Далее нам необходимо удалить это IP-адрес из файла testLog.log

	// Проверим тип переменной takeFirstIp
	fmt.Println("Тип переменной takeFirstIp -", reflect.TypeOf(takeFirstIp))
	// Удалим пробелы, если есть, в строчной переменной takeFirstIp
	takeFirstIp = strings.ReplaceAll(takeFirstIp, " ", "")

	// Создаём массив, где будут храниться Ip-адреса, которые мы будем извлекать по
	// одному из файла testLog.log

	var ipInStringArr [500]string
	var indexForArr int = 0 // Индекс массива, который будет увеличиваться инкрементом
	// после того, как строка с ip-адресом будет положена в массив.
	ipInStringArr[indexForArr] = takeFirstIp
	fmt.Print("Массиву с индексом [", indexForArr, "] было присвоено значение ", ipInStringArr[0], "\n")
	indexForArr++
	fmt.Println("Индекс массива indexForArr увеличился и принял значение -", indexForArr)

	// Дальше посчитаем, количество символов в строчной переменной
	// takeFirstIp

	howManyLetters := (len(takeFirstIp)) + 2
	fmt.Println("В строке ", takeFirstIp, "-", howManyLetters, "символов.")

	// Теперь удалим количество знаков переменной howManyLetters
	// из строчной переменной listIpInString
	subListIpInString := listIpInString[howManyLetters:]
	//fmt.Print(subListIpInString)

	// Теперь создадим новый файл и запишем в него строчную переменную subListIpInString для проверки
	createNewLogFile, err := os.Create("newTestLog.log")
	if err != nil {
		panic(err)
	}
	defer createNewLogFile.Close()

	createNewLogFile.Write([]byte(subListIpInString))

}
