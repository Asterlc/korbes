package concorrencia

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ExeConcorrencia() {
	wg.Add(runtime.NumGoroutine())
	go um()
	dois()
	fmt.Println(runtime.NumGoroutine())
}

func um() {
	for i := 0; i < 10; i++ {
		fmt.Println("Função um", i)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

func dois() {
	for i := 0; i < 10; i++ {
		fmt.Println("Função dois", i)
		time.Sleep(time.Second * 1)
	}
}
