package model

import (
	"math/rand"
	"time"
)

func generarSecuenciaPar() [4]int {
	var resultado [4]int
	rand.Seed(time.Now().Unix())
	semilla := rand.Intn(100)
	for i := 0; i < 4; i++ {
		resultado[i] = semilla + i*2
	}
	return resultado
}

func generarSecuenciaInpar() [4]int {
	var resultado [4]int
	rand.Seed(time.Now().Unix())
	semilla := rand.Intn(100)
	for i := 0; i < 4; i++ {
		resultado[i] = semilla + i*2 + 1
	}
	return resultado
}

func generarFibonacci() [4]int {
	var resultado [4]int
	rand.Seed(time.Now().Unix())
	semilla := rand.Intn(50)
	for i := 0; i < 4; i++ {
		if i < 2 {
			resultado[i] = semilla
		} else {
			resultado[i] = resultado[i-1] + resultado[i-2]
		}
	}
	return resultado
}

func IniciarSecuencia() [4]int {
	rand.Seed(time.Now().Unix())
	semilla := rand.Intn(3)
	switch semilla {
	case 0:
		return generarSecuenciaPar()
	case 1:
		return generarSecuenciaInpar()
	default:
		return generarFibonacci()
	}
}
