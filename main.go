package main

import (
	"bufio"
	"fmt"
	"os"
	"simpletodolist/internal/todo"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")

		if ok := scan.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		command := strings.Fields(scan.Text())

		if len(command) == 0 {
			fmt.Println("")
			fmt.Println("Вы ничего не ввели")
			fmt.Println("")

		} else {
			switch command[0] {
			case "help":
				todo.Help()

			case "exit":
				todo.ExitDo()
				return

			case "add":
				todo.Add(command)

			default:
				fmt.Println("")
				fmt.Println("Некорректный ввод")
				fmt.Println("")
			}
		}
	}
}
