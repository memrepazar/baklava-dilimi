package main

import "fmt"

func main() {
	var eus int
	a := 0
	k := 0
	fmt.Print("en uzun satır : ")
	_,err := fmt.Scanln(&eus)
	if err != nil {
		fmt.Println("hatalı değer girişi")
		return
	}
	if eus%2==0 {
		fmt.Println("değer tek sayı olmalı")
		return
	}

	for i := 0; i < eus/2+1; i++ {
		for b := 0; b < eus/2-i; b++ {
			fmt.Print(" ")
		}
		for y := 0; y < 1+2*a; y++ {
			fmt.Print("*")
		}
		a++
		fmt.Println()

	}
	for i := 0; i < eus/2; i++ {
		for b := 0; b < 1+i; b++ {
			fmt.Print(" ")
		}
		for y := 0; y < eus-2-2*k; y++ {
			fmt.Print("*")
		}
		k++
		fmt.Println()
	}
}
