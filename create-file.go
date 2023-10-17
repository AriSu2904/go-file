package main

import (
	"fmt"
	"os"
)

func main() {

	path := "example.txt"
	value := "Halo ini adalah example.txt"
	err := os.WriteFile(path, []byte(value), 0700)
	if err != nil {
		fmt.Println(err)
	}
}

func createFile(path string) {
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err == nil {
			fmt.Println("Berhasil Membuat File!")
		}
		defer file.Close()
	} else {
		fmt.Println("File sudah ada di direktori saat ini")
	}
}
