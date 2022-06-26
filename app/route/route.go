package route

import (
	"fmt"
	"errors"
	"os"
	"bufio"
	"strings"
	"strconv"
	"encoding/json"
)

type Route struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool `json:"finished"`
}

func(pRoute *Route) LoadPositions() error {
	if pRoute.ID == "" {
		return errors.New("Route ID is empty")
	}

	var filePath = "destinations/" pRoute.ID + ".txt"

	file, error := os.Open(filePath)

	if error != nil {
		fmt.Println("File '" + filePath + "' not found.")
		return error
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		
		lat, error := strconv.ParseFloat(data[0], 64)
		if error != nil {
			return nil
		}
		long, error := strconv.ParseFloat(data[1], 64)
		if error != nil {
			return nil
		}

		pRoute.Positions = append(pRoute.Positions, Position{Lat: lat, Long: long})

	}
	return nil
}

func(pRoute *Route) ExportJsonPositions() ([]string, error) {
	var partialRoute PartialRoutePosition
	var result []string
	totalPositions := len(pRoute.Positions)

	for loopIndex, coordinates := range pRoute.Positions {
		partialRoute.ID = pRoute.ID
		partialRoute.ClientID = pRoute.ClientID
		partialRoute.Position = []float64{coordinates.Lat, coordinates.Long}
		partialRoute.Finished = false
		if totalPositions -1 == loopIndex {
			partialRoute.Finished = true
		}
		jsonRoute, error := json.Marshal(partialRoute)
		if error != nil {
			return nil, error
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}