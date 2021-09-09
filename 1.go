package main

import (
	"fmt"
	"math"
	"strconv"
)

func Print(input string) {
	fmt.Println(input);
}

func Trunc(number float64) int {
	return int(math.Trunc(number));
}

func main() {
	var first string;
	fmt.Scanln(&first);

	var second float64;
	fmt.Scanln(&second);
	
    Print(first);
	Print(strconv.Itoa(Trunc(second)));
}