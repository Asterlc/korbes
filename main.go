package main

import (
	"kobres/canais"
)

func main() {
	canal := make(chan int)
	canais.TesteCanalRange(10, canal)
}
