package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetLatestLotecaResult() {
	url := "https://servicebus2.caixa.gov.br/portaldeloterias/api/loteca"

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

	var loteca Data
	err = json.Unmarshal(body, &loteca)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("   _       _                        ")
	fmt.Println("  | |     | |                       ")
	fmt.Println("  | | ___ | |_ ___  ___ __ _        ")
	fmt.Println("  | |/ _ \\| __/ _ \\/ __/ _` |     ")
	fmt.Println("  | | (_) | ||  __/ (_| (_| |       ")
	fmt.Println("  |_|\\___/ \\__\\___|\\___\\__,_|  ")
	fmt.Println("                                    ")

	fmt.Println("Resultado")
	fmt.Printf("Concurso: %d (%s)\n\n", loteca.ConcursoNume, loteca.DataApuracao)

	fmt.Printf("-------------------------------------------------------\n\n")
	for _, premio := range loteca.ListaPremios {

		if premio.Faixa == 1 {
			fmt.Println(premio.DescricaoFaixa)
			fmt.Printf("%d ganhador(es), R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}

		if premio.Faixa == 2 {
			fmt.Println(premio.DescricaoFaixa)
			fmt.Printf("%d ganhadores, R$ %s\n\n", premio.NumGanhadores, FormatCurrency(premio.ValorPremio))
		}
	}
	fmt.Printf("-------------------------------------------------------\n\n")

	for _, equipe := range loteca.EqEsportivas {
		fmt.Printf("%s - %s\n", equipe.DiaJogo, equipe.DataJogo)
		fmt.Printf("%s %d x %d %s\n\n", equipe.EquipeUm, equipe.NumGolEqUm, equipe.NumGolEqDois, equipe.EquipeDois)
	}
	fmt.Printf("-------------------------------------------------------\n\n")

	if loteca.Acumulado {
		fmt.Println("Acumulou!!!")
	}
	fmt.Println("Próximo Sorteio:", loteca.DataProximoS)
	fmt.Println("Valor estimado do próximo concurso R$", FormatCurrency(loteca.ProximoValor))
	fmt.Println("Acumulado próximo concurso R$", FormatCurrency(loteca.ValAcumulado))
	fmt.Println("Acumulado próximo concurso final zero/cinco R$", FormatCurrency(loteca.ValAcumula05))
}
