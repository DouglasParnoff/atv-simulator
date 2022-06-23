package route

import {
	"errors"
	"os"
	"bufio"
	"strings"
	"strconv"
}

type Route struct {
	ID string
	ClientID string
	Positions []Position
}

type Position struct {
	Lat float64
	Long float64
}

func(route *Route) loadPositions() error {
	if route.ID == "" {
		return errors.New("Route ID is empty")
	}

	file, error := os.Open("destinations/" + route.ID + ".txt")

	if error != nil {
		return error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		
		lat, error := strconv.parseFloat(data[0], 64)
		if error != nil {
			return nil
		}
		long, error := strconv.parseFloat(data[1], 64)
		if error != nil {
			return nil
		}

		route.Positions = append(route.Positions, Position{Lat: lat, Long: long})

		return nil
	}
}