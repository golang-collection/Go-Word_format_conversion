package main

import (
	"Go-Word_format_conversion/cmd"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-14 20:09
* @Description:
**/

func main() {
	err := cmd.Execute()
	if err != nil{
		log.Fatalf("cmd.Execute err: %v", err)
	}
}