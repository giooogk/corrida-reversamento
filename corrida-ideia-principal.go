package main

import (
	"fmt"
	"sync"
	"time"
)

var WaitGroup sync.WaitGroup
var reverzamentoTotal int = 4
var numPasasgem int

func corredor1(numCorredor chan int, corredor string) {
	posisaoBastao := <-numCorredor

	fmt.Printf("Corredor %d começou a correr\n", posisaoBastao)

	if posisaoBastao != reverzamentoTotal {
		numPasasgem = posisaoBastao + 1
		etapaCorredor()
	}

	numCorredor <- numPasasgem

}

func corredor2(numCorredor chan int, corredor string) {
	posisaoBastao := <-numCorredor

	if posisaoBastao != reverzamentoTotal {

		numPasasgem = posisaoBastao + 1
		fmt.Println(corredor, "recebeu o bastao")
		fmt.Println(corredor, " começou a correr")
		etapaCorredor()
	}

	numCorredor <- numPasasgem

}

func corredor3(numCorredor chan int, corredor string) {
	posisaoBastao := <-numCorredor

	if posisaoBastao != reverzamentoTotal {

		numPasasgem = posisaoBastao + 1
		fmt.Println(corredor, "recebeu o bastao")
		fmt.Println(corredor, " começou a correr")
		etapaCorredor()
	}

	numCorredor <- numPasasgem

}

func corredor4(numCorredor chan int, corredor string) {
	posisaoBastao := <-numCorredor

	if posisaoBastao == reverzamentoTotal {
		numPasasgem = posisaoBastao + 1
		fmt.Println(corredor, "recebeu o bastao")
		fmt.Println(corredor, " começou a correr")
		etapaCorredor()
		fmt.Printf("Corredor %d chegou ao final\n", posisaoBastao)

	}

	numCorredor <- 0

}

func etapaCorredor() {
	time.Sleep(2 * time.Second)
}

func main() {

	channel := make(chan int)

	tInicial := time.Now()
	fmt.Println("--- Corrida iniciada ---", tInicial.Format(time.Stamp))

	WaitGroup.Add(4)
	go corredor1(channel, "Corredor 1")
	channel <- 1
	go corredor2(channel, "Corredor 2")
	go corredor3(channel, "Corredor 3")
	go corredor4(channel, "Corredor 4")
	WaitGroup.Wait()

	tFinal := time.Now()
	fmt.Println("--- Corrida finalizada ---", tFinal.Format(time.Stamp))

}
