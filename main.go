package main

import (
	"fmt"
	newRoute "github.com/diegofranca92/simulador/application/route"
)

func main()  {
	route := newRoute.Route{
		ID: "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}