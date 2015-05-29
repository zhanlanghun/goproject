package main

import (
	"common/webcommon"
	"config"
	"fmt"
)

type person struct {
}

func main() {
	result := config.Config("fdsfdsa")
	result2 := webcommon.Commonconfig("dddd")
	fmt.Println(result)
	fmt.Println(result2)
	//per := make(person)
	//fmt.Println(per)

	px := [10]int{2, 4, 1, 9, 2, 5, 1, 6}
	n := len(&px)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if px[i] > px[j] {
				tem := px[j]
				px[j] = px[i]
				px[i] = tem
			}
		}
	}
	fmt.Println(px)
}
