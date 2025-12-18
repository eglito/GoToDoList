package main

import (
	"bufio"
	"fmt"
	"os"
)

type task struct {
	name  string
	data  string
	texto string
}

var lista = make([]string, 0)

func main() {

	lista = append(lista, "Comprar cigarro")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bem-vindo ao toDoList - aqui você realiza suas tarefas")

	for {

		if scanner.Scan() {
			textoDigitado := scanner.Text()

			if textoDigitado == "add" {
				fmt.Println("Insira um item na lista: ")
				scanner.Scan()
				lista = addTasK(lista, scanner.Text())
			}

			if textoDigitado == "list" {
				fmt.Println("")
				fmt.Println("A seguir, sua lista de afazeres: ")
				returnList(lista)
			}

			if textoDigitado == "end" {
				break
			}
		}
	}
}

func addTasK(lista []string, stg string) []string {
	lista = append(lista, stg)
	return lista
}

func returnList(lista []string) {
	for i, p := range lista {
		fmt.Printf("%d: %s\n", i+1, p)
	}
}
