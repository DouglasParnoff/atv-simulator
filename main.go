package main

import (
	"fmt"
	deliveryRoute "github.com/douglasparnoff/atv-simulator/app/route"
)

func main() {
	route := deliveryRoute.Route{
		ID: "123", 
		ClientID: "456",
	}
	
	route.LoadPositions()

	jsonPositions, _ := route.ExportJsonPositions()
	fmt.Println(jsonPositions[0])
}