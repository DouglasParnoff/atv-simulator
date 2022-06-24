package main

import {
	"fmt"
	deliveryRoute "github.com/douglasparnoff/atv-simulator/app/route"
}

func main() {
	route := deliveryRoute.Route{ID: "123", ClientID: "456"}
	route.loadPositions()
	jsonPositions, _ := route.exportJsonPositions()
	fmt.Println(jsonPositions[0])
}