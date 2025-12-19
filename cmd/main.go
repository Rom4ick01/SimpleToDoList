package main

import (
	"bufio"
	"fmt"
	"os"
	"simpletodolist/internal/event"
	"simpletodolist/internal/todo"
	"strings"
	"time"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")

		if ok := scan.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scan.Text()

		command := strings.Fields(text)

		if len(command) == 0 {
			fmt.Println("\nВы ничего не ввели\n")
			event.NewEvent(text, "Вы ничего не ввели", time.Now().Format("2006-01-02 15:04:05"))

		} else {
			switch command[0] {
			case "help":
				todo.Help(text)

			case "exit":
				todo.ExitDo()
				return

			case "events":
				event.ShowEvent()

			case "add":
				todo.Add(command, text)

			case "list":
				todo.List(text)

			case "del":
				todo.Del(command, text)

			case "done":
				todo.Done(command, text)

			default:
				fmt.Println("\nНекорректный ввод\n")
				event.NewEvent(text, "Некорректный ввод", time.Now().Format("2006-01-02 15:04:05"))
			}
		}
	}
}
