package main

import (
	"example/secuencia/secuencia/model"
	"fmt"
	"os"
)

func main() {

	fmt.Println("------------- Bienvenido !! ----------------")

	exit := 'Y'
	puntaje := 0

	for exit != 'N' {
		secuencia := model.IniciarSecuencia()
		for i := 0; i < len(secuencia); i++ {
			if i != 2 {
				fmt.Print(secuencia[i])
				fmt.Print(" ")
			} else {
				fmt.Print("________  ")
			}
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Print(" Faltante :")

		var nro int

		n, err := fmt.Scanf("%d\n", &nro)
		if err != nil {
			fmt.Println(n, err)
			os.Exit(n)
		}

		if nro == secuencia[2] {
			fmt.Print(" Ganaste!! ")
			puntaje++
		} else {
			fmt.Print(" Perdiste!! ")
			puntaje--
		}
		fmt.Println(" Puntaje : ", puntaje)
		fmt.Print(" Desea continuar jugando ? (Y/N) ")
		fmt.Scanf("%c\n", &exit)
	}

}
