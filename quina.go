package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestQuinaResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/quina"

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

	var quina Data
	err = json.Unmarshal(body, &quina)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("                                    ")
	fmt.Println("    ____        _                   ")
	fmt.Println("   / __ \\      (_)                 ")
	fmt.Println("  | |  | |_   _ _ _ __   __ _       ")
	fmt.Println("  | |  | | | | | | '_ \\ / _` |     ")
	fmt.Println("  | |__| | |_| | | | | | (_| |      ")
	fmt.Println("   \\____\\_\\__,_|_|_| |_|\\__,_|  ")
	fmt.Println("                                    ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", quina.ConcursoNume, quina.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", quina.LocalSorteio, quina.MunicSorteio)
	fmt.Printf("Numeros Sorteados: %s\n\n", strings.Join(quina.NumSorteados, ", "))

	fmt.Printf("-------------------------------------------------------\n\n")
	for _, premio := range quina.ListaPremios {

		if premio.Faixa == 1 {
			fmt.Println("Quina:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 2 {
			fmt.Println("Quadra:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 3 {
			fmt.Println("Terno:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 4 {
			fmt.Println("Duque:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}
	}
	fmt.Printf("-------------------------------------------------------\n\n")

	if quina.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Proximo Sorteio:", quina.DataProximoS)
	fmt.Println("Valor Estimado do Proximo Concurso R$", FormatCurrency(quina.ProximoValor))
}
