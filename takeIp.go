// Берём ip в лог-файле и считаем количество повторений
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Функция для сравнения строки с элементами массива
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	takeContent, err := ioutil.ReadFile("/home/ryazanov/Downloads/tdlfFile.txt") // Берем данные из лог-файла

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

	// Создаём файл, в который будем записывать строковую переменую
	// со списком ip-адресов

	createLogFile, err := os.Create("logTostringIp.log")
	if err != nil {
		panic(err)
	}
	defer createLogFile.Close()

	// Осуществляем запись ip-адресов в файл

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

	//-------------Блок считает пробелы в переменной-------------------
	// Это позволяет определить количество ip-адресов в файле
	ipList, err = ioutil.ReadFile("logTostringIp.log")

	if err != nil {
		log.Fatal(err)
	}

	listIpInString = string(ipList)
	spacesInString := strings.Count(listIpInString, " ")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("--------------------------------------------------------")
	//---------------------------------------------------------------------

	// Получаем первый ip-адрес из переменной listIpInString
	takeFirstIp := pRe.FindString(listIpInString)
	// Удалим пробелы, если есть, в строчной переменной takeFirstIp
	takeFirstIp = strings.ReplaceAll(takeFirstIp, " ", "")

	// Создаём структуру с массивом ip-адресов и
	// массив с индексом массива
	type IpList struct {
		ipInStringSlice []string // Сюда будем помещать ip-адреса
		indexForSlice   int      // Индекс массива ipInStringArr
	}

	a := new(IpList)    // Инициализация структуры IpList
	a.indexForSlice = 0 // Индекс среза начинается с 0
	// Положим ip из takeFirstIp в срез
	a.ipInStringSlice = append(a.ipInStringSlice, takeFirstIp)

	// Считаем количество символов ip-адреса в переменной takeFirstIp
	howManyLetters := (len(takeFirstIp)) + 1
	// Теперь удалим количество знаков переменной howManyLetters
	// из строчной переменной listIpInString
	subListIpInString := listIpInString[howManyLetters:]
	halfSpacesInString := spacesInString / 2

	for i := 0; i <= halfSpacesInString-2; i++ {
		//--------------------Повторяющийся блок-----------------------------
		// Получаем первый ip теперь из переменной subListIpInString
		takeFirstIp = pRe.FindString(subListIpInString)
		// Удалим пробелы, если есть, в строчной переменной takeFirstIp
		takeFirstIp = strings.ReplaceAll(takeFirstIp, " ", "")
		// Положим ip из takeFirstIp в срез
		a.ipInStringSlice = append(a.ipInStringSlice, takeFirstIp)
		// Считаем количество символов ip-адреса в переменной takeFirstIp
		howManyLetters = (len(takeFirstIp)) + 1
		// Теперь удалим количество знаков переменной howManyLetters
		// из строчной переменной subListIpInString
		listIpInString = subListIpInString[howManyLetters:]
		takeFirstIp = pRe.FindString(listIpInString)
		// Удалим пробелы, если есть, в строчной переменной takeFirstIp
		takeFirstIp = strings.ReplaceAll(takeFirstIp, " ", "")
		// Положим ip из takeFirstIp в срез
		a.ipInStringSlice = append(a.ipInStringSlice, takeFirstIp)
		// Считаем количество символов ip-адреса в переменной takeFirstIp
		howManyLetters = (len(takeFirstIp)) + 1
		subListIpInString = listIpInString[howManyLetters:]
		if i%100 == 0 {
			fmt.Print("*")
		}
		//--------------------Конец повторяющегося блока-----------------------
	}
	fmt.Println("")
	fmt.Println("Длинна среза ip-адресов =", len(a.ipInStringSlice))

	//-----------Блок подсчёта запросов с одного ip-адреса-----------------------
	type CountingRequests struct {
		ipTable              []string // Срез для ip, которые были посчитаны
		countsIpTables       []int    // Количество повторений ip ipTable
		indexIpTable         int
		counter              int
		indexIpInStringSlice int
	}

	c := new(CountingRequests)
	c.indexIpTable = 0
	c.ipTable = append(c.ipTable, a.ipInStringSlice[0])
	c.countsIpTables = append(c.countsIpTables, 1)
	// Проверим, сколько первый элемент списка повторяется в срезе a.ipInStringSlice
	c.indexIpInStringSlice = 1 // индекс среза, в котором хранятся ip
	c.counter = 1
	c.indexIpTable = 0

	for y := len(a.ipInStringSlice); y >= 0; y-- {
		if c.indexIpInStringSlice < (spacesInString - 3) {
			//---Новый блок поиска уникальных ip-адресов для добавления их в срез ipTable------
			for {
				f := c.ipTable
				g := a.ipInStringSlice[c.indexIpInStringSlice]

				result := Contains(f, g)
				// fmt.Println(result)
				if result == true {
					c.indexIpInStringSlice += 1
					continue
				} else {
					c.ipTable = append(c.ipTable, a.ipInStringSlice[c.indexIpInStringSlice])
					break
				}
				// --------------Конец нового блока поиска уникальных ip---------------------------
			}

			//-----------Блок подсчёта повторений ip-адреса--------------------------------

			c.indexIpTable += 1
			c.counter = 1

			for i := spacesInString - 3; i > 0; i-- {
				if c.ipTable[c.indexIpTable] == a.ipInStringSlice[i] {
					c.counter += 1
				}
			}
			c.countsIpTables = append(c.countsIpTables, c.counter)
			//--------------------Конец блока повторений---------------------------------------
		}
	}
	fmt.Println(c.ipTable)
	fmt.Println(c.countsIpTables)
	fmt.Println(len(c.ipTable))

	// Далее выводим результат, - сколько запросов с каждого уникального ip
	lenUnicIpTable := len(c.ipTable)
	c.indexIpTable = 0
	for p := lenUnicIpTable - 1; p >= 0; p-- {
		fmt.Print(c.ipTable[p])
		for l := 20 - len(c.ipTable[p]); l >= 0; l-- {

			fmt.Print(" ")
		}
		fmt.Print(" - (", c.countsIpTables[p], ")           - ")
		for i := c.countsIpTables[p]; i > 0; i-- {
			fmt.Print("|")
		}
		fmt.Println("")
	}
}
