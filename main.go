package main

import (
	"log"
	"math/rand"
)

var totalFolhas int

// Ramo eh um elemento da arvore binaria
type Ramo struct {
	Chave int
	Esquerda *Ramo
	Direita *Ramo
}

func (r *Ramo) inserir(valor int) {
	if r.Chave < valor {
		if r.Direita == nil {
			r.Direita = &Ramo{Chave: valor}
		} else {
			r.Direita.inserir(valor)
		}
	} else if r.Chave > valor {
		if r.Esquerda == nil {
			r.Esquerda = &Ramo{Chave: valor}
		} else {
			r.Esquerda.inserir(valor)
		}
	}
}

func (r *Ramo) busca(valor, profundidade int) (bool, int) {
	if r == nil {
		return false, 0
	}
	profundidade++
	if r.Chave < valor {
		return r.Direita.busca(valor, profundidade)
	} else if r.Chave > valor {
		return r.Esquerda.busca(valor, profundidade)
	}
	return true, profundidade
}

func (r *Ramo) imprimir() {
	if r == nil {
		return 
	}
	log.Printf("%d\n", r.Chave)
	totalFolhas++
	if r.Direita != nil {
		r.Direita.imprimir()
	} 
	if r.Esquerda != nil {
		r.Esquerda.imprimir()
	}
}

func main()  {
	//Parametros
	valorChaveInicial := 250
	valorLimiteAleatorio := 500
	valorMax := 52
	valorMin := 380

	arvore := &Ramo{Chave: valorChaveInicial}
	log.Printf("Estado atual da arvore é: \n")
	arvore.imprimir()
	for i:=0; i<100; i++ {
		arvore.inserir(rand.Intn(valorLimiteAleatorio))
	}
	arvore.inserir(valorMax)
	arvore.inserir(valorMin)
	log.Printf("Estado atual da arvore é: \n")
	arvore.imprimir()
	log.Printf("Total de folhas é de %d\n", totalFolhas)
	resultado, profundidade := arvore.busca(valorMin, 0)
	log.Printf("Resultado da busca por %d é: %t com profundidade de %d\n", valorMin, resultado, profundidade)
	resultado, profundidade = arvore.busca(valorMax, 0)
	log.Printf("Resultado da busca por %d é: %t com profundidade de %d\n", valorMax, resultado, profundidade)
}