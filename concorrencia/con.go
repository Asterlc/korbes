package concorrencia

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func ExeConcorrencia() {

	wg.Add(1)
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

func Corrida() {
	contador := 0
	totalRoutines := 12000

	wg.Add(totalRoutines)
	for i := 0; i < totalRoutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Rotinas", runtime.NumGoroutine())
	fmt.Println("contador:", contador)
}

func CorridaMutex() {
	contador := 0
	totalRoutines := 12000

	wg.Add(totalRoutines)
	for i := 0; i < totalRoutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Rotinas", runtime.NumGoroutine())
	fmt.Println("contador:", contador)
}

func CorridaAtomic() {
	var contador int64 = 0
	totalRoutines := 12000

	wg.Add(totalRoutines)
	for i := 0; i < totalRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			atomic.LoadInt64(&contador)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Rotinas", runtime.NumGoroutine())
	fmt.Println("contador:", contador)
}

func contaDevagar(index int) {
	time.Sleep(time.Second * 2)
	fmt.Println("DONE:", index)

}

func ContaRapida(numero int) {
	rotinas := 10000000
	wg.Add(rotinas)
	for i := 0; i < rotinas; i++ {
		go func() {
			contaDevagar(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x", runtime.NumGoroutine())
}
