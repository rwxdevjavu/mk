package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func createFile(filename string) error {
	filename = strings.Replace(filename, "\\", "/", -1)
	dir := strings.Split(filename, "/")
	filepath := strings.Join(dir[:len(dir)-1], "/")
	if len(dir) > 1 {
		if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
			return err
		}
	}
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if _, err := os.Create(filename); err != nil {
			return err
		}

	} else if !fileInfo.IsDir() {
		return errors.New("File Already Exists")
	}
	return nil
}

func help() {
	fmt.Println(`Usage: mk filepath
	
Example:
	mk Untitled.txt
	`)
}

func main() {
	switch {
	case len(os.Args) == 1:
		fmt.Println("Error: File Name missing")
		fmt.Println("Run 'mk help' for usage")
	case os.Args[1] == "":
		fmt.Println("Error: File Name missing")
		fmt.Println("Run 'mk help' for usage")
	case os.Args[1] == "help":
		help()
	default:
		if err := createFile(os.Args[1]); err != nil {
			fmt.Println(err)
			fmt.Println("Run 'mk help' for usage")
		}
	}
}
