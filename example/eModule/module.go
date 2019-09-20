package main

import (
	"fmt"

	"github.com/fuwensun/goexample/mypkg/hello"
	"github.com/fuwensun/goexample/mypkg/world"
	// "github.com/fuwensun/goexample/hello"
	// "github.com/fuwensun/goexample/world"
)

func module() {
	fmt.Println("module()")
	var h hello.Hello
	h.Print()
	var w world.World
	w.Print()
}
