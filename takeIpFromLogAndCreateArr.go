// Эта программа извлекает ip-адреса из лог-файла и
// присваивает их массиву данных строкового типа.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	takeContent, err := ioutil.ReadFile("/home/sergey/Downloads/tdlfFile.txt") // Берем данные из лог-файла

	if err != nil {
		log.Fatal(err)
	}

	// Присваиваем данные переменной строкового типа str1
	str1 := string(takeContent)

	// Осуществляем поиск ip-адресов с помощью регулярного выражения
	// и складываем полученные данные в созданный файл

	pRe := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", pRe.String()) // print pattern
	fmt.Println(pRe.MatchString(str1))        // true

	createLogFile, err := os.Create("logTostringIp.log")
	if err != nil {
		panic(err)
	}
	defer createLogFile.Close()

	submatchall := pRe.FindAllString(str1, -1)
	for _, element := range submatchall {
		_, err = createLogFile.WriteString(" ")
		_, err = createLogFile.WriteString(element)
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println("Был создан файл logTostringIp.log")
	fmt.Println("В файл logTostringIp.log записана последовательность \nip-адресов, которые были получены из лог-файла")
	fmt.Println("------------------------------------------------------")

	// Читаем файл logTostringIp.log в строчную переменную listIpInString

	ipList, err := ioutil.ReadFile("logTostringIp.log")

	if err != nil {
		log.Fatal(err)
	}
	listIpInString := string(ipList)
	fmt.Println(pRe.MatchString(listIpInString)) // true

	// Проверяем, есть ли пробел в начале строки и удаляем его
	listIpInString = strings.TrimSpace(listIpInString)

	// Теперь создадим новый файл и запишем в него строчную переменную subListIpInString для проверки
	createNewLogFile, err := os.Create("newlogTostringIp.log")
	if err != nil {
		panic(err)
	}
	defer createNewLogFile.Close()

	createNewLogFile.Write([]byte(listIpInString))

	// Удаляем файл logTostringIp.log
	err = os.Remove("logTostringIp.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Был удалён файл logTostringIp.log")

	// Переименуем файл newlogTostringIp.log в logTostringIp.log
	err = os.Rename("newlogTostringIp.log", "logTostringIp.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Файл newlogTostringIp.log был переименован в logTostringIp.log")

}
