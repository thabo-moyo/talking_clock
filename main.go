package main

import (
	"fmt"
	"github.com/talking_clock/pkg/utils"
	"log"
	"os"
)

func main() {
	//Argument without program
	timeInput := os.Args[1:]

	//Check if terminal argument is empty
	if len(timeInput) < 0 {
		fmt.Println(utils.GetCurrentTime())
	} else {
		if len(timeInput[0]) > 0 {
			fmt.Println(utils.GetTimeByString(timeInput[0]))
		} else {
			log.Fatal("Error: Missing time argument.")
		}
	}

	/*	cmd.StartServer()*/
}
