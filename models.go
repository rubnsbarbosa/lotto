package main

type Premio struct {
	DescricaoFaixa string  `json:"descricaoFaixa"`
	Faixa          int     `json:"faixa"`
	NumGanhadores  int     `json:"numeroDeGanhadores"`
	ValorPremio    float64 `json:"valorPremio"`
}

type Data struct {
	Acumulado    bool     `json:"acumulado"`
	DataApuracao string   `json:"dataApuracao"`
	DataProximoS string   `json:"dataProximoConcurso"`
	NumSorteados []string `json:"listaDezenas"`
	ListaPremios []Premio `json:"listaRateioPremio"`
	LocalSorteio string   `json:"localSorteio"`
	MunicSorteio string   `json:"nomeMunicipioUFSorteio"`
	ConcursoNume int      `json:"numero"`
	ProximoValor float64  `json:"valorEstimadoProximoConcurso"`
}
