package concorrencia

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ExeConcorrencia() {
	wg.Add(1)
	go um()
	dois()
}

func um() {
	for i := 0; i < 10; i++ {
		fmt.Println("Função um", i)
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}

func dois() {
	for i := 0; i < 10; i++ {
		fmt.Println("Função dois", i)
		time.Sleep(time.Second * 1)
	}
}
