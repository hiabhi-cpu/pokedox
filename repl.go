package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	con := getConfig()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		input = strings.ToLower(strings.Fields(strings.Trim(input, " "))[0])
		comm, d := getCommands(&con)[input]
		if d {
			err := comm.callback(&con)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(con)
		} else {
			fmt.Println("Unkown Command")
		}

	}
}

func cleanInputText(text string) []string {
	text = strings.ToLower(text)
	temp := strings.Split(text, " ")
	res := []string{}
	for _, str := range temp {
		if len(str) != 0 {
			res = append(res, str)
		}
	}
	return res
}
