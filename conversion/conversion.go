package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nilai32 int32 = 3294;

	var nilai64 int64 = int64(nilai32);
	var stringNumb string = strconv.Itoa(int(nilai64))


	fmt.Println("nilai 64", nilai64);
	fmt.Println("string numb", stringNumb);
}