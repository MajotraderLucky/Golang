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
	takeContent, err := ioutil.ReadFile("/home/sergey/Downloads/auth.log") // Берем данные из лог-файла

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
	fmt.Println("--------------------------------------------------------")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Эти данные выводятся до применения цикла for:")
	spacesInString := strings.Count(listIpInString, " ")
	fmt.Println("Количество пробелов в строке listIpInString =", spacesInString)
	//fmt.Println("Количество пробелов в строке listIpInString =", strings.Count(listIpInString, " "))
	fmt.Println("--------------------------------------------------------")
	fmt.Println("--------------------------------------------------------")
	//---------------------------------------------------------------------

	// Получаем первый ip-адрес из переменной listIpInString
	takeFirstIp := pRe.FindString(listIpInString)
	// Удалим пробелы, если есть, в строчной переменной takeFirstIp
	takeFirstIp = strings.ReplaceAll(takeFirstIp, " ", "")
	fmt.Println(takeFirstIp, "- Выводим переменную takeFirstIp")

	// Создаём структуру с массивом ip-адресов и
	// массив с индексом массива
	type IpList struct {
		ipInStringArr [100000]string // Сюда будем помещать ip-адреса
		indexForArr   int            // Индекс массива ipInStringArr
	}

	a := new(IpList)  // Инициализация структуры IpList
	a.indexForArr = 0 // Индекс массива, который будет увеличиваться инкрементом
	// после того как строка с ip-адресом будет положена в массив.
	a.ipInStringArr[a.indexForArr] = takeFirstIp
	fmt.Print("Массиву с индексом [", a.indexForArr, "] было присвоено значение ", a.ipInStringArr[0], "\n")
	// a.indexForArr++
	// fmt.Println("Индекс массива indexForArr увеличился =", a.indexForArr)
	fmt.Println("--------------------------------------------------------")

	// Считаем количество символов ip-адреса в переменной takeFirstIp
	howManyLetters := (len(takeFirstIp)) + 1
	fmt.Println("В строке ", takeFirstIp, "-", howManyLetters, "символов.")

	// Теперь удалим количество знаков переменной howManyLetters
	// из строчной переменной listIpInString
	subListIpInString := listIpInString[howManyLetters:]
	fmt.Println("Из строки listIpInString было удалено ", howManyLetters, "символов")

	// Теперь создадим новый файл и запишем в него строчную переменную subListIpInString для проверки
	createNewLogFile, err = os.Create("newlogTostringIp.log")
	if err != nil {
		panic(err)
	}
	defer createNewLogFile.Close()

	createNewLogFile.Write([]byte(subListIpInString))
	fmt.Println("Был создан временный файл newlogTostringIp.log,")
	fmt.Println("в который была записана изменённая переменная \nlistIpInString")
	fmt.Println("--------------------------------------------------------")

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

	// Далее нужно добавить цикл for, который перенесёт ip-адреса в
	// массив ipInStringArr

	// В цикле for должны повторятся действия по извлечению, удалению
	// и добавлению ip в масиив ipInStringArr

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Запускаем цикл for")
	fmt.Println("--------------------------------------------------------")

	for i := 1000; i >= 0; i-- {
		ipList, err = ioutil.ReadFile("logTostringIp.log")

		if i%10 == 0 {
			fmt.Print(".")
		}

		if err != nil {
			log.Fatal(err)
		}
		listIpInString = string(ipList)
		takeFirstIp = pRe.FindString(listIpInString)

		// Удалим пробелы, если есть, в строчной переменной takeFirstIp
		takeFirstIp = strings.ReplaceAll(takeFirstIp, " ", "")

		a.indexForArr++
		a.ipInStringArr[a.indexForArr] = takeFirstIp
		// Дальше посчитаем, количество символов в строчной переменной takeFirstIp
		howManyLetters = (len(takeFirstIp)) + 1
		// Теперь удалим количество знаков переменной howManyLetters
		// из строчной переменной listIpInString
		subListIpInString = listIpInString[howManyLetters:]
		// Создаём новый файл newlogTostringIp.log и записываем в него subListIpInString
		createNewLogFile, err = os.Create("newlogTostringIp.log")
		if err != nil {
			panic(err)
		}
		defer createNewLogFile.Close()

		createNewLogFile.Write([]byte(subListIpInString))

		// Удаляем файл logTostringIp.log и переименовываем файл newlogTostringIp.log
		// в logTostringIp.log
		err = os.Remove("logTostringIp.log")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Переименуем файл newlogTostringIp.log в logTostringIp.log
		err = os.Rename("newlogTostringIp.log", "logTostringIp.log")
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	fmt.Println(" ")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Выводим массив ipInStringArr ", a.ipInStringArr[0:50])
	fmt.Println("--------------------------------------------------------")

	// Далее необходимо определить, сколько раз повторяется ip-адрес в массиве
	// ipInStringArr

	//-------------Блок считает пробелы в переменной-------------------
	// Это позволяет определить количество ip-адресов в файле
	ipList, err = ioutil.ReadFile("logTostringIp.log")

	if err != nil {
		log.Fatal(err)
	}

	listIpInString = string(ipList)
	fmt.Println("--------------------------------------------------------")
	fmt.Println("--------------------------------------------------------")
	spacesInString = strings.Count(listIpInString, " ")
	fmt.Println("Количество пробелов в строке listIpInString =", spacesInString)
	//fmt.Println("Количество пробелов в строке listIpInString =", strings.Count(listIpInString, " "))
	fmt.Println("--------------------------------------------------------")
	fmt.Println("--------------------------------------------------------")
	//---------------------------------------------------------------------
	// Далее отсортируем ip-адреса и добавим в срез, откуда
	// сделаем вывод с результатом подсчёта, сколько раз повторяется
	// ip в массиве.

}
