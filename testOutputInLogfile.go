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

	// Создаём структуру, в которой будут храниться:
	// строковый массив ip-адресов и количество запросов по каждому ip
	// данные будем получать из файла testLog.log

	type IpList struct {
		ipInStringArr    [500]string // Сюда будем помещать Ip
		numberOfRequests [500]int    // Здесь будем считать количество запросов
		indexForArr      int
	}

	a := new(IpList)

	//var ipInStringArr [500]string
	a.indexForArr = 0 // Индекс массива, который будет увеличиваться инкрементом
	// после того, как строка с ip-адресом будет положена в массив.
	a.ipInStringArr[a.indexForArr] = takeFirstIp
	fmt.Print("Массиву с индексом [", a.indexForArr, "] было присвоено значение ", a.ipInStringArr[0], "\n")

	a.numberOfRequests[a.indexForArr] = 1 // ip-адрес массива a.ipInStringArr[a.indexForArr] был упомянут a.indexForArr 1 раз.

	a.indexForArr++
	fmt.Println("Индекс массива indexForArr увеличился и принял значение -", a.indexForArr)

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
