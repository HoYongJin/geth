package main

import (
	"fmt"
)

func checOne(num int, c chan bool) {
	if num == 1 {
		c <- true
	} else {
		c <- false
	}
}

func main() {
	num := 0
	ch := make(chan bool)

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	go checOne(num, ch)

	res := <-ch
	fmt.Println(res)

}
