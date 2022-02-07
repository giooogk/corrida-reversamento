package main

import (
	"fmt"
	"sync"
	"time"
)

var WaitGroup sync.WaitGroup
var reverzamentoTotal int = 4
var proximoCorredor int

func corredor(numCorredor chan int) {

	posisaoBastao := <-numCorredor

	fmt.Printf("Corredor %d começou a correr\n", posisaoBastao)

	if posisaoBastao < reverzamentoTotal {
		etapaCorredor()
		proximoCorredor = posisaoBastao + 1
		fmt.Printf("PEGA!!!, corredor %d \n", proximoCorredor)
		go corredor(numCorredor)
	} else {
		etapaCorredor()
		fmt.Printf("Corredor %d chegou ao final\n", posisaoBastao)
		WaitGroup.Done()
		return
	}

	fmt.Printf("Corredor %d recebeu o bastão\n", proximoCorredor)

	numCorredor <- proximoCorredor

}

func etapaCorredor() {
	time.Sleep(3 * time.Second)
}

func main() {

	channel := make(chan int)
	tInicial := time.Now()
	fmt.Println("--- Corrida iniciada ---", tInicial.Format(time.Stamp))

	WaitGroup.Add(1)
	go corredor(channel)
	channel <- 1
	WaitGroup.Wait()

	tFinal := time.Now()
	fmt.Println("--- Corrida finalizada ---", tFinal.Format(time.Stamp))

}
