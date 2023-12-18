package main

import (
	"github.com/rogeriofbrito/go-mvc/src/factory"
)

func main() {
	ca := factory.NewControllerAdapter()
	ca.Start()
}
