package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	intro()
	for {
		menu()

		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func getNameAndAge() (string, int) {
	name := "Amanda"
	age := 28
	return name, age
}

func intro() {

	name, age := getNameAndAge()

	version := 1.1
	fmt.Println("Olá, sra.", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)
	return command
}

func menu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func monitoring() {
	fmt.Println("Monitorando...")
	site := "https://alura.com.br/askjdbahsbciahsbca"
	response, _ := http.Get(site)
	
	if response.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
    }
}
