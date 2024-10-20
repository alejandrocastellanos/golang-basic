package main

import (
	"fmt"
	"golang_basic/basic_api"
	"golang_basic/basic_tutorial"
)

func basics() {
	fmt.Println("Hello, World!")
	basic_tutorial.BasicTypes()

	var addResult int = basic_tutorial.Add(2, 4)
	fmt.Println("The ADD result is", addResult)

	ifControlResult := basic_tutorial.IfControl(10)
	fmt.Println(ifControlResult)

	forControlResult := basic_tutorial.ForControl(10)
	fmt.Println("El resultado de sumar cada numero del 1 al 10 es", forControlResult)
}

func asyncMethods() {
	basic_tutorial.GorutinesBasic()
	basic_tutorial.CallWorker()
	basic_tutorial.CallChannels()
	basic_tutorial.Library()
}

func main() {
	basic_api.RunApi()
}
