package main

import (
	"bufio"
	"fmt"
	"os"
)

type task struct {
	name string
	data string
	note string
}

var taskList = []task{}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem-vindo ao toDoList - aqui você realiza suas tarefas")

	for {

		if scanner.Scan() {
			typedText := scanner.Text()

			if typedText == "add" {
				fmt.Println("Insira um item na lista: ")
				scanner.Scan()
				NewTask := task{note: scanner.Text()}
				taskList = addTasK(taskList, NewTask)
			}

			if typedText == "list" {
				fmt.Println("")
				fmt.Println("A seguir, sua lista de afazeres: ")
				printList(taskList)
			}

			if typedText == "end" {
				break
			}
		}
	}
}

func addTasK(list []task, newTask task) []task {
	return append(list, newTask)

}

func printList(list []task) {
	for i, p := range list {
		fmt.Printf("%d.\t%s\n", i+1, p.note)
	}
}
