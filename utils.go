package main

import (
	"fmt"
	"os"
)

func checkIsDir(path string) bool {
	if src, err := os.Stat(path); err == nil && src.IsDir() {
		return true
	}

	return false
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Printf("%s [-p <port>] <folder1> <folder2> ... --fileserve all the folders on given port\n", os.Args[0])
}
