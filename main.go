package main

import (
	"kobres/canais"
)

func main() {
	canalQualquer := make(chan int)
	canais.TesteCanal2(canalQualquer)
	canais.TesteCanal()
}
