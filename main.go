package main

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/factory"
)

func main() {
	ca := factory.NewControllerAdapter()
	ca.Start()
}
