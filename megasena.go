package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestMegaSenaResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/megasena"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Create a new GET request was not possible due to:", err)
		os.Exit(1)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Send the GET request was not possible due to:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Unexpected HTTP response status code:", res.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read the response body was not possible due to:", err)
		os.Exit(1)
	}

	var megasena Data
	err = json.Unmarshal(body, &megasena)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("                                                           ")
	fmt.Println("  _ __ ___   ___  __ _  __ _   ___  ___ _ __   __ _        ")
	fmt.Println(" | '_ ` _ \\ / _ \\/ _` |/ _` | / __|/ _ \\ '_ \\ / _` |   ")
	fmt.Println(" | | | | | |  __/ (_| | (_| | \\__ \\  __/ | | | (_| |     ")
	fmt.Println(" |_| |_| |_|\\___|\\__, |\\__,_| |___/\\___|_| |_|\\__,_|  ")
	fmt.Println("                  __/ |                                    ")
	fmt.Println("                 |___/                                     ")
	fmt.Println("                                                           ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", megasena.ConcursoNume, megasena.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", megasena.LocalSorteio, megasena.MunicSorteio)
	fmt.Printf("Números sorteados: %s\n\n", strings.Join(megasena.NumSorteados, ", "))

	fmt.Printf("-------------------------------------------------------\n\n")
	for _, premio := range megasena.ListaPremios {

		if premio.Faixa == 1 {
			fmt.Println("Mega Sena:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhador(es), R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 2 {
			fmt.Println("Quina:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 3 {
			fmt.Println("Quadra:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}
	}
	fmt.Printf("-------------------------------------------------------\n\n")

	if megasena.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Próximo sorteio:", megasena.DataProximoS)
	fmt.Println("Valor estimado do próximo concurso R$", FormatCurrency(megasena.ProximoValor))
}
