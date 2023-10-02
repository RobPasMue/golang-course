package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Ensure that you have the proper number of arguments
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments. Expecting only 1, the file to be read.")
		os.Exit(1)
	}

	// Ensure that you can read the file
	//
	// OPTION 1
	//
	// bytes, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error encountered: ", err)
	// }
	// fmt.Println(string(bytes))
	//
	// OPTION 2
	//
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error encountered: ", err)
	}
	io.Copy(os.Stdout, file)
}
