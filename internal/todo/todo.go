package todo

import (
	"fmt"
	"slices"
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

func Help() {
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
}

func ExitDo() {
	fmt.Println("")
	fmt.Println("Вы завершили работу программы, всего доброго!")
}

func Add(c []string) {
	switch len(c) {
	case 1, 2:
		fmt.Println("")
		fmt.Println("Некорректный ввод")
		fmt.Println("")

	default:
		txt := ""
		for i := 2; i < len(c); i++ {
			txt += c[i]
		}

		allTask = append(allTask, &Task{c[1], txt, time.Now().Format("2006-01-02 15:04:05"), false, ""})

		pp.Println("Вы добавили задачу:", allTask)
	}
}

func List() {
	if len(allTask) == 0 {
		fmt.Println("")
		fmt.Println("У вас нет задач")
		fmt.Println("")
	} else {
		fmt.Println("")
		pp.Println(allTask)
		fmt.Println("")
	}
}

func Del(c []string) {
	if len(c) > 2 || len(c) == 1 {
		fmt.Println("")
		fmt.Println("Некорректный ввод")
		fmt.Println("")
	} else {
		switch c[1] {
		case "":
			fmt.Println("")
			fmt.Println("Некорректный ввод")
			fmt.Println("")

		default:
			for k, v := range allTask {
				if c[1] == v.title {
					fmt.Println("")
					fmt.Println("Задача", v.title, " удалена")
					fmt.Println("")
					poiter := &allTask
					*poiter = slices.Delete(allTask, k, k+1)
					return
				}
			}

			fmt.Println("")
			fmt.Println("Не найдено такой задачи")
			fmt.Println("")
		}
	}
}

func Done(c []string) {
	if len(c) > 2 || len(c) == 1 {
		fmt.Println("")
		fmt.Println("Некорректный ввод")
		fmt.Println("")
	} else {
		switch c[1] {
		case "":
			fmt.Println("")
			fmt.Println("Некорректный ввод")
			fmt.Println("")

		default:
			for _, v := range allTask {
				if c[1] == v.title {
					fmt.Println("")
					fmt.Println("Задача", v.title, " отмечена как выполненная")
					fmt.Println("")
					v.done = true
					v.timedone = time.Now().Format("2006-01-02 15:04:05")

					return
				}
			}

			fmt.Println("")
			fmt.Println("Не найдено такой задачи")
			fmt.Println("")
		}
	}
}
