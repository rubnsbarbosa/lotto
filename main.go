package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nConfira os últimos resultados")

		fmt.Println("01: +milionária")
		fmt.Println("02: mega sena")
		fmt.Println("03: lotofácil")
		fmt.Println("04: quina")
		fmt.Println("05: lotomania")
		fmt.Println("06: timemania")
		fmt.Println("07: dupla sena")
		fmt.Println("08: federal")
		fmt.Println("09: loteca")
		fmt.Println("10: dia de sorte")
		fmt.Println("11: super sete")
		fmt.Println("00: sair")
		fmt.Print("Digite o número da opção desejada: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Número inválido. Digite um número válido.")
			continue
		}

		switch option {
		case 1:
			fmt.Println("Buscando último resultado da +milionária")
			GetLatestMaisMilionariaResult()
		case 2:
			fmt.Println("Buscando último resultado da mega sena")
			GetLatestMegaSenaResult()
		case 3:
			fmt.Println("Buscando último resultado da lotofácil")
			GetLatestLotoFacilResult()
		case 4:
			fmt.Println("Buscando último resultado da quina")
			GetLatestQuinaResult()
		case 5:
			fmt.Println("Buscando último resultado da lotomania")
			GetLatestLotoManiaResult()
		case 6:
			fmt.Println("Buscando último resultado da timemania")
			GetLatestTimeManiaResult()
		case 7:
			fmt.Println("Buscando último resultado da dupla sena")
			GetLatestDuplaSenaResult()
		case 8:
			fmt.Println("Buscando último resultado da federal")
			GetLatestFederalResult()
		case 9:
			fmt.Println("Buscando último resultado da loteca")
			GetLatestLotecaResult()
		case 10:
			fmt.Println("Buscando último resultado dia de sorte")
			GetLatestDiaDeSorteResult()
		case 11:
			fmt.Println("Buscando último resultado da super sete")
			GetLatestSuperSeteResult()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Selecione uma opção válida.")
		}
	}
}
