package canais

import (
	"fmt"
)

func TesteCanal() {
	canal := make(chan string)
	go func() {
		canal <- "Hello"
	}()
	go func() {
		canal <- "World"
	}()

	x := <-canal
	fmt.Println(string(x))
}

func TesteCanal2(canal chan int) {
	go send(canal)
	receive(canal)
}

func send(s chan<- int) {
	s <- 1
}

func receive(r <-chan int) {
	x := <-r
	fmt.Println(x)
}

func TesteCanalRange(n int, c chan int) {
	go meuLoop(n, c)

	for v := range c {
		fmt.Println("Recebido do canal:", v)
	}
}

func meuLoop(t int, c chan<- int) {
	for i := 0; i < t; i++ {
		c <- i
	}
	close(c)
}
