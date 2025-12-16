package todo

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	Title      string
	text       string
	createtime string
	done       bool
	timedone   string
}

var allTask []Task

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
		now := time.Now()
		formatted := now.Format("2006-01-02 15:04:05")

		allTask = append(allTask, Task{c[1], txt, formatted, false, ""})

		pp.Println("Вы добавили задачу:", allTask)
	}
}
