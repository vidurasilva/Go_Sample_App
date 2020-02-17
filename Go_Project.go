package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*Reade two file*/
	data, err := ioutil.ReadFile("D:/Project/GO/Go_apps/Data_File.txt")
	/*check errors*/
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
