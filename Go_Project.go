package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// /*Reade two file*/
	// data, err := ioutil.ReadFile("D:/Project/GO/Go_app/Data_File.txt")
	// /*check errors*/
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	// fmt.Println("Contents of file:", string(data))
	readFile, err := os.Open("D:/Project/GO/Go_app/Data_File.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	// for _, eachline := range fileTextLines {
	// 	if eachline == "### classes text file" {
	// 		fmt.Println(eachline)
	// 	}
	// }
	/* for loop execution */

	for i, x := range fileTextLines {
		fmt.Println(x, i)
	}
}
