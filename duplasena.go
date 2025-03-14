package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestDuplaSenaResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/duplasena"

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

	var duplasena Data
	err = json.Unmarshal(body, &duplasena)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("       _             _                                     ")
	fmt.Println("      | |           | |                                    ")
	fmt.Println("    __| |_   _ _ __ | | __ _   ___  ___ _ __   __ _        ")
	fmt.Println("   / _` | | | | '_ \\| |/ _` | / __|/ _ \\ '_ \\ / _` |    ")
	fmt.Println("  | (_| | |_| | |_) | | (_| | \\__ \\  __/ | | | (_| |     ")
	fmt.Println("   \\__,_|\\__,_| .__/|_|\\__,_| |___/\\___|_| |_|\\__,_|  ")
	fmt.Println("              | |                                          ")
	fmt.Println("              |_|                                          ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", duplasena.ConcursoNume, duplasena.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", duplasena.LocalSorteio, duplasena.MunicSorteio)
	fmt.Printf("Numeros do 1ª sorteio: %s\n", strings.Join(duplasena.NumSorteados, ", "))
	fmt.Printf("Numeros do 2ª sorteio: %s\n\n", strings.Join(duplasena.ND2Sorteados, ", "))

	fmt.Printf("-------------------------------------------------------\n\n")
	for _, premio := range duplasena.ListaPremios {

		if premio.Faixa == 1 {
			fmt.Println("1ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 2 {
			fmt.Println("2ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 3 {
			fmt.Println("3ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 4 {
			fmt.Println("4ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 5 {
			fmt.Println("5ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 6 {
			fmt.Println("6ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 7 {
			fmt.Println("7ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 8 {
			fmt.Println("8ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}
	}
	fmt.Printf("-------------------------------------------------------\n\n")

	if duplasena.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Proximo Sorteio:", duplasena.DataProximoS)
	fmt.Println("Valor Estimado do Proximo Concurso R$", FormatCurrency(duplasena.ProximoValor))
}
