package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLatestMaisMilionariaResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/maismilionaria"

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

	var milionaria Data
	err = json.Unmarshal(body, &milionaria)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("                  _ _ _               __       _           ")
	fmt.Println("    _            (_) (_)             /_/      (_)          ")
	fmt.Println("  _| |_ _ __ ___  _| |_  ___  _ __   __ _ _ __ _  __ _     ")
	fmt.Println(" |_   _| '_ ` _ \\| | | |/ _ \\| '_ \\ / _` | '__| |/ _` | ")
	fmt.Println("   |_| | | | | | | | | | (_) | | | | (_| | |  | | (_| |    ")
	fmt.Println("       |_| |_| |_|_|_|_|\\___/|_| |_|\\__,_|_|  |_|\\__,_| ")
	fmt.Println("                                                           ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", milionaria.ConcursoNume, milionaria.DataApuracao)

	fmt.Printf("Sorteio realizado no %s em %s\n", milionaria.LocalSorteio, milionaria.MunicSorteio)
	fmt.Printf("Números sorteados: %s\n", strings.Join(milionaria.NumSorteados, ", "))
	fmt.Printf("Trevos sorteados: %s\n\n", strings.Join(milionaria.TrvSorteados, ", "))

	fmt.Printf("-------------------------------------------------------\n\n")
	for _, premio := range milionaria.ListaPremios {

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

		if premio.Faixa == 9 {
			fmt.Println("9ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 10 {
			fmt.Println("10ª faixa:", premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}
	}
	fmt.Printf("-------------------------------------------------------\n\n")

	if milionaria.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Próximo sorteio:", milionaria.DataProximoS)
	fmt.Println("Valor estimado do próximo concurso R$", FormatCurrency(milionaria.ProximoValor))
}
