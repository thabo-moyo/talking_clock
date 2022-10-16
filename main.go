package main

import (
	"fmt"
	"github.com/talking_clock/cmd"
	"github.com/talking_clock/pkg/utils"
	"os"
)

func main() {
	//Argument without program
	timeInput := os.Args[1:]

	//Check if terminal argument is empty
	if len(os.Args) < 2 {
		fmt.Println(utils.GetCurrentTime())
	} else {
		byString, err := utils.GetTimeByString(timeInput[0])
		if err != nil {
			fmt.Println("Incorrect time format")
		}

		fmt.Println(byString)
	}
	//Starting server for rest service
	cmd.StartServer()
}
