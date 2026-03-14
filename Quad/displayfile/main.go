package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	data, _ := os.ReadFile(os.Args[1]) // تجاهلنا الـ error هنا
	fmt.Print(string(data))
}
