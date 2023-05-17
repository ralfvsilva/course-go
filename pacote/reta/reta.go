package main

import "math"

// Iniciando com letra maiuscula - Público (visível dentro e fora do pacote.)
// Iniciando com letra minuscula - Privado (visível apenas dentro do pacote.)

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - vai gerar algo privado.

// Ponto representa uma coordenada no plano cartesiano.
type Ponto struct { // Struct pública
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) { // privado no nível do pacote, não no do arquivo
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distância é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 { // método público
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
