package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestSuperSeteResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/supersete"

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

	var supersete Data
	err = json.Unmarshal(body, &supersete)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("                                          _              ")
	fmt.Println("                                         | |             ")
	fmt.Println("    ___ _   _ _ __   ___ _ __    ___  ___| |_ ___        ")
	fmt.Println("   / __| | | | '_ \\ / _ \\ '__|  / __|/ _ \\ __/ _ \\   ")
	fmt.Println("   \\__ \\ |_| | |_) |  __/ |     \\__ \\  __/ ||  __/   ")
	fmt.Println("   |___/\\__,_| .__/ \\___|_|     |___/\\___|\\__\\___|  ")
	fmt.Println("             | |                                         ")
	fmt.Println("             |_|                                         ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", supersete.ConcursoNume, supersete.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", supersete.LocalSorteio, supersete.MunicSorteio)
	fmt.Printf("Numeros Sorteados: %s\n\n", strings.Join(supersete.NumSorteados, ", "))

	fmt.Printf("----------------------------------------------------------\n\n")
	for _, premio := range supersete.ListaPremios {

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

	if supersete.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Proximo Sorteio:", supersete.DataProximoS)
	fmt.Println("Valor Estimado do Proximo Concurso R$", FormatCurrency(supersete.ProximoValor))
}
