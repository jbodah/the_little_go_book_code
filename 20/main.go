package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("a_file_to_read")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 256)
	file.Read(buffer)
	fmt.Println(buffer)
}
