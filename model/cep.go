package model

//DadosCep Recebe dados em formato json
type DadosCep struct {
	Bairro     string     `json:"bairro,omitempty" db:"bairro"`
	Cidade     string     `json:"cidade,omitempty" db:"cidade"`
	Logradouro string     `json:"logradouro,omitempty" db:"logradouro"`
	UFINFO     EstadoInfo `json:"estado_info,omitempty"`
	CEP        string     `json:"cep,omitempty" db:"cep"`
	CidadeInfo CidadeInfo `json:"cidade_info,omitempty"`
	Estado     string     `json:"estado,omitempty" db:"estado"`
}

//EstadoInfo teste teste teste
type EstadoInfo struct {
	AreaKm2    string `json:"area_km2,omitempty" db:"uf_area_km2"`
	CodigoIbge string `json:"codigo_ibge,omitempty" db:"codigo_ibge"`
	Nome       string `json:"nome,omitempty" db:"uf_nome"`
}

//CidadeInfo teste teste teste
type CidadeInfo struct {
	AreaKm2    string `json:"area_km2,omitempty" db:"cidade_area_km2"`
	CodigoIbge string `json:"codigo_ibge,omitempty" db:"cidade_codigo_ibge"`
}
