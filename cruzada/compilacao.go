package cruzada

import (
	"fmt"
	"runtime"
)

//GOOS=windows GOARCH=amd64 go build <arquivo>

func Welcome() {
	fmt.Println("Criado em um Ubuntu Version 22.04 e exutando em um:", runtime.GOOS, runtime.GOARCH)
}
