package todo

import (
	"fmt"
	"simpletodolist/internal/event"
	"slices"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	title      string
	text       string
	createtime string
	done       bool
	timedone   string
}

var allTask []*Task

func Help(t string) {
	fmt.Println("")
	fmt.Println("Список команд:")
	fmt.Println("help — эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("add {заголовок задачи из одного слова} {текст задачи} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("list — эта команда позволяет получить полный список всех задач")
	fmt.Println("del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("events — эта команда позволяет получить список всех событий")
	fmt.Println("exit — эта команда позволяет завершить выполнение программы")
	fmt.Println("")

	event.NewEvent(t, "", time.Now().Format("2006-01-02 15:04:05"))
}

func ExitDo() {
	fmt.Println("")
	fmt.Println("Вы завершили работу программы, всего доброго!")
}

func Add(c []string, t string) {
	switch len(c) {
	case 1, 2:
		fmt.Println("")
		fmt.Println("Некорректный ввод")
		fmt.Println("")

		event.NewEvent(t, "Некорректный ввод", time.Now().Format("2006-01-02 15:04:05"))

	default:
		txt := strings.Join(c[2:], " ")

		allTask = append(allTask, &Task{c[1], txt, time.Now().Format("2006-01-02 15:04:05"), false, ""})

		pp.Println("Вы добавили задачу:", allTask)

		event.NewEvent(t, "", time.Now().Format("2006-01-02 15:04:05"))
	}
}

func List(t string) {
	if len(allTask) == 0 {
		fmt.Println("")
		fmt.Println("У вас нет задач")
		fmt.Println("")

		event.NewEvent(t, "У вас нет задач", time.Now().Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println("")
		pp.Println(allTask)
		fmt.Println("")

		event.NewEvent(t, "", time.Now().Format("2006-01-02 15:04:05"))
	}
}

func Del(c []string, t string) {
	if len(c) > 2 || len(c) == 1 {
		fmt.Println("")
		fmt.Println("Некорректный ввод")
		fmt.Println("")

		event.NewEvent(t, "Некорректный ввод", time.Now().Format("2006-01-02 15:04:05"))
	} else {
		switch c[1] {
		case "":
			fmt.Println("")
			fmt.Println("Некорректный ввод")
			fmt.Println("")

			event.NewEvent(t, "Некорректный ввод", time.Now().Format("2006-01-02 15:04:05"))

		default:
			for k, v := range allTask {
				if c[1] == v.title {
					fmt.Println("")
					fmt.Println("Задача", v.title, " удалена")
					fmt.Println("")
					allTask = slices.Delete(allTask, k, k+1)
					event.NewEvent(t, "", time.Now().Format("2006-01-02 15:04:05"))
					return
				}
			}

			fmt.Println("")
			fmt.Println("Не найдено такой задачи")
			fmt.Println("")

			event.NewEvent(t, "Не найдено такой задачи", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

func Done(c []string, t string) {
	if len(c) > 2 || len(c) == 1 {
		fmt.Println("")
		fmt.Println("Некорректный ввод")
		fmt.Println("")

		event.NewEvent(t, "Некорректный ввод", time.Now().Format("2006-01-02 15:04:05"))
	} else {
		switch c[1] {
		case "":
			fmt.Println("")
			fmt.Println("Некорректный ввод")
			fmt.Println("")

			event.NewEvent(t, "Некорректный ввод", time.Now().Format("2006-01-02 15:04:05"))

		default:
			for _, v := range allTask {
				if c[1] == v.title {
					fmt.Println("")
					fmt.Println("Задача", v.title, " отмечена как выполненная")
					fmt.Println("")
					v.done = true
					v.timedone = time.Now().Format("2006-01-02 15:04:05")

					event.NewEvent(t, "", time.Now().Format("2006-01-02 15:04:05"))
					return
				}
			}

			fmt.Println("")
			fmt.Println("Не найдено такой задачи")
			fmt.Println("")

			event.NewEvent(t, "Не найдено такой задачи", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
