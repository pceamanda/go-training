package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
	"io/ioutil"
)

const monitoramentos = 3
const delay = 5

func main() {

	intro()
	for {
		menu()

		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			printLogs()
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

func getSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

	reader := bufio.NewReader(file)

	for {
		line, err :=reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	fmt.Println("O meu slice tem", len(sites), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(sites), "itens")

	return sites
}

func monitoring() {
	fmt.Println("Monitorando...")
	sites := getSites()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		log(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		log(site, false)
	}
}

func log(site string, status bool) {

    file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + 
		" - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

    file, err := ioutil.ReadFile("log.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(file))
}