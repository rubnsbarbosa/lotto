package main

type Premio struct {
	DescricaoFaixa string  `json:"descricaoFaixa"`
	Faixa          int     `json:"faixa"`
	NumGanhadores  int     `json:"numeroDeGanhadores"`
	ValorPremio    float64 `json:"valorPremio"`
}

type UF struct {
	Municipio    string `json:"municipio"`
	NomeFatansia string `json:"nomeFatansiaUL"`
	Posicao      int    `json:"posicao"`
	UF           string `json:"uf"`
}

type Equipe struct {
	DiaJogo      string `json:"diaSemana"`
	DataJogo     string `json:"dtJogo"`
	EquipeUm     string `json:"nomeEquipeUm"`
	EquipeDois   string `json:"nomeEquipeDois"`
	NumGolEqUm   int    `json:"nuGolEquipeUm"`
	NumGolEqDois int    `json:"nuGolEquipeDois"`
}

type Data struct {
	Acumulado    bool     `json:"acumulado"`
	DataApuracao string   `json:"dataApuracao"`
	DataProximoS string   `json:"dataProximoConcurso"`
	NumSorteados []string `json:"listaDezenas"`
	ND2Sorteados []string `json:"listaDezenasSegundoSorteio"`
	TrvSorteados []string `json:"trevosSorteados"`
	ListaPremios []Premio `json:"listaRateioPremio"`
	UFGanhadores []UF     `json:"listaMunicipioUFGanhadores"`
	EqEsportivas []Equipe `json:"listaResultadoEquipeEsportiva"`
	LocalSorteio string   `json:"localSorteio"`
	MunicSorteio string   `json:"nomeMunicipioUFSorteio"`
	ConcursoNume int      `json:"numero"`
	ProximoValor float64  `json:"valorEstimadoProximoConcurso"`
	ValAcumulado float64  `json:"valorAcumuladoProximoConcurso"`
	ValAcumula05 float64  `json:"valorAcumuladoConcurso_0_5"`
}
