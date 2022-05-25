package main

import "fmt"
func main() {
	var name string
	var age int

	fmt.Println("Digite seu nome")
	fmt.Scan(&name)

	fmt.Println("Digite sua idade")
	fmt.Scan(&age)

	version := 1.1
	fmt.Println("Olá, sra.", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var command int
	fmt.Scan(&command)

	fmt.Println("O comando escolhido foi", command)
}
