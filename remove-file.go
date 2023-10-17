package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Remove("example.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Berhasil menghapus File!")
}
