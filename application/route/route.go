package route

import (
	"errors"
	"os"
	"bufio"
	"strings"
	"strconv"
	"encoding/json"
)

type Route struct {
	ID string		`json: "routeId"`
	ClientID string	`json: "clientId"`
	Positions []Position	`json: "position"`
}

type Position struct {
	Lat float64		`json: "lat"`
	Long float64	`json: "long"`
}

type PartialRoutePosition struct {
	// Go permite ja definir como ficara a chave no JSON
	ID string		`json: "routeId"`
	ClientID string		`json: "clientId"`
	Position []float64		`json: "position"`
	Finished bool		`json: "finished"`
}

// Carregar as posicoes das rodas - Nesse exemplo para esses arquivos
func(r *Route) LoadPositions() error {
	// idenficar o Id do arquivo de rotas a ser lidas
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	// Abrir o arquivo
	file, err := os.Open("destinations/" + r.ID + ".txt")
	// retornar erro caso nao ler o arquivo
	if err != nil {
		return err
	}
	// fechar arquivo
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// ler linha a linha do arquivo
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		// conversao para Float os indices do Split
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil { return nil }
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil { return nil }

		r.Positions = append(r.Positions, Position{
			Lat: lat,
			Long: long,
		})

	}

	return nil
}

// Percorrer as posicoes criadas e montar um JSON
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for key, val := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{val.Lat, val.Long}
		route.Finished = false

		if total-1 == key {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)

		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}