package main

import (
	"fmt"
)

func main() {
	type IpList struct {
		ipArr            [500]string
		numberOfRequests [500]int
	}
	a := new(IpList)

	a.ipArr[0] = "adf"
	a.numberOfRequests[0] = 1

	fmt.Println(a.ipArr[0], "---", a.numberOfRequests[0])
}
