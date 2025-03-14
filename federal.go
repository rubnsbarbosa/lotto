package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestFederalResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/federal"

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

	var federal Data
	err = json.Unmarshal(body, &federal)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("    __         _                _      ")
	fmt.Println("   / _|       | |              | |     ")
	fmt.Println("  | |_ ___  __| | ___ _ __ __ _| |     ")
	fmt.Println("  |  _/ _ \\/ _` |/ _ \\ '__/ _` | |   ")
	fmt.Println("  | ||  __/ (_| |  __/ | | (_| | |     ")
	fmt.Println("  |_| \\___|\\__,_|\\___|_|  \\__,_|_| ")
	fmt.Println("                                       ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", federal.ConcursoNume, federal.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", federal.LocalSorteio, federal.MunicSorteio)
	fmt.Printf("Bilhetes sorteados: %s\n\n", strings.Join(federal.NumSorteados, ", "))

	fmt.Printf("-------------------------------------------------------\n\n")
	for _, premio := range federal.ListaPremios {

		if premio.Faixa == 1 {
			fmt.Println("1ª bilhete sorteado:", federal.NumSorteados[0])
			fmt.Println("Unidade Lotérica:", federal.UFGanhadores[4].NomeFatansia)
			fmt.Printf("Cidade/UF: %s/%s\n", federal.UFGanhadores[4].Municipio, federal.UFGanhadores[4].UF)
			fmt.Printf("Valor do Prêmio (R$) %s\n\n", FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 2 {
			fmt.Println("2ª bilhete sorteado:", federal.NumSorteados[1])
			fmt.Println("Unidade Lotérica:", federal.UFGanhadores[2].NomeFatansia)
			fmt.Printf("Cidade/UF: %s/%s\n", federal.UFGanhadores[2].Municipio, federal.UFGanhadores[2].UF)
			fmt.Printf("Valor do Prêmio (R$) %s\n\n", FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 3 {
			fmt.Println("3ª bilhete sorteado:", federal.NumSorteados[2])
			fmt.Println("Unidade Lotérica:", federal.UFGanhadores[3].NomeFatansia)
			fmt.Printf("Cidade/UF: %s/%s\n", federal.UFGanhadores[3].Municipio, federal.UFGanhadores[3].UF)
			fmt.Printf("Valor do Prêmio (R$) %s\n\n", FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 4 {
			fmt.Println("4ª bilhete sorteado:", federal.NumSorteados[3])
			fmt.Println("Unidade Lotérica:", federal.UFGanhadores[0].NomeFatansia)
			fmt.Printf("Cidade/UF: %s/%s\n", federal.UFGanhadores[0].Municipio, federal.UFGanhadores[0].UF)
			fmt.Printf("Valor do Prêmio (R$) %s\n\n", FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 5 {
			fmt.Println("5ª bilhete sorteado:", federal.NumSorteados[4])
			fmt.Println("Unidade Lotérica:", federal.UFGanhadores[1].NomeFatansia)
			fmt.Printf("Cidade/UF: %s/%s\n", federal.UFGanhadores[1].Municipio, federal.UFGanhadores[1].UF)
			fmt.Printf("Valor do Prêmio (R$) %s\n\n", FormatCurrency(premio.ValorPremio))
		}
	}
	fmt.Println("-------------------------------------------------------")

}
