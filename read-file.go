package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	//konversi ke string
	convert := string(data)

	fmt.Println(convert)

}
