package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestDiaDeSorteResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/diadesorte"

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

	var diadesorte Data
	err = json.Unmarshal(body, &diadesorte)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("        _ _              _                        _                ")
	fmt.Println("       | (_)            | |                      | |               ")
	fmt.Println("     __| |_  __ _     __| | ___    ___  ___  _ __| |_ ___          ")
	fmt.Println("    / _` | |/ _` |   / _` |/ _ \\  / __|/ _ \\| '__| __/ _ \\      ")
	fmt.Println("   | (_| | | (_| |  | (_| |  __/  \\__ \\ (_) | |  | ||  __/       ")
	fmt.Println("    \\__,_|_|\\__,_|   \\__,_|\\___|  |___/\\___/|_|   \\__\\___|  ")
	fmt.Println("                                                                   ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", diadesorte.ConcursoNume, diadesorte.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", diadesorte.LocalSorteio, diadesorte.MunicSorteio)
	fmt.Printf("Numeros Sorteados: %s\n\n", strings.Join(diadesorte.NumSorteados, ", "))

	fmt.Printf("----------------------------------------------------------\n\n")
	for _, premio := range diadesorte.ListaPremios {

		if premio.Faixa == 1 {
			fmt.Println("1ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhador(es), R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
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
	}
	fmt.Printf("----------------------------------------------------------\n\n")

	if diadesorte.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Proximo Sorteio:", diadesorte.DataProximoS)
	fmt.Println("Valor Estimado do Proximo Concurso R$", FormatCurrency(diadesorte.ProximoValor))
}
