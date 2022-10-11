package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/mutant", func(c *gin.Context) {
		var dato InfoADN

		if err := c.BindJSON(&dato); err != nil {
			fmt.Println(err)
		}
		fmt.Println(dato.Adn)
		esMutante := isMutant(dato.Adn)
		if !esMutante {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "No es mutante",
			})
			return
		}
		//fmt.Println(adn)
		c.JSON(http.StatusOK, gin.H{
			"message": "Es mutante",
		})
	})
	r.Run()

	/* var adn = [...]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"} */
	/* var adn = [...]string{"ATCCCC", "CAGTTC", "TTATGT", "AGAAGG", "CCCCGA", "TCACGG"} */
	/* int n := len(adn) */
	/* esMutante := isMutant(adn) */
	/* fmt.Println("es verdaderamente mutante", esMutante) */

}

type InfoADN struct {
	Adn [6]string `json:"adn"`
}

func isMutant(adn [6]string) bool {
	var secuenciaTotal, secuenciaParcialRow, secuenciaParcialCol, secuenciaParcialObl int
	secuenciaTotal = 0
	secuenciaParcialRow = int(secuenciaRow(adn))
	secuenciaTotal = secuenciaTotal + secuenciaParcialRow
	secuenciaParcialCol = int(secuenciaCol(adn))
	secuenciaTotal = secuenciaTotal + secuenciaParcialCol
	println("inicia la secuencia Oblicua")
	secuenciaParcialObl = int(secuenciaOblicua(adn))
	secuenciaTotal = secuenciaTotal + secuenciaParcialObl

	if secuenciaTotal >= 2 {
		return true
	}
	return false

	/* fmt.Println(strings.Contains(adn[4], "CCCC")) */
}

func secuenciaOblicua(adn [6]string) int {
	var i, j int
	var cadena, newCadena, antCadena string
	var repeticionesObl, secuencia int
	repeticionesObl = 0

	for i = 0; i < (len(adn)+1)/2; i++ {
		for j = 0; j < (len(adn)+1)/2; j++ {
			cadena = adn[i]

			if len(cadena) > j+3 {
				newCadena = adn[i+3]
				/* println("cadena " + cadena + "nueva cadena " + newCadena) */
				if string(cadena[j]) == string(newCadena[j+3]) {
					/* 					println("entra primer if oblicua")
					 */antCadena = cadena
					cadena = adn[i+1]
					newCadena = adn[i+2]

					if string(cadena[j+1]) == string(newCadena[j+2]) && string(cadena[j+1]) == string(antCadena[j]) {
						repeticionesObl = repeticionesObl + 2
						secuencia = secuencia + 1
						fmt.Println("hay una secuencia Oblicua")
					}
				}
			}
			repeticionesObl = 0

			/* if() */
		}
	}

	return secuencia
}

func secuenciaRow(adn [6]string) int {

	var i, j int
	var cadena string
	var repeticionesRow, secuencia int
	repeticionesRow = 0
	for i = 0; i < len(adn); i++ {
		println()
		for j = 0; j < len(adn); j = j + 3 {
			cadena = adn[i]

			if len(cadena) > j+3 {
				if string(cadena[j]) == string(cadena[j+3]) {
					repeticionesRow = repeticionesRow + 2
					if string(cadena[j+2]) == string(cadena[j+1]) && string(cadena[j]) == string(cadena[j+1]) {
						repeticionesRow = repeticionesRow + 2
						secuencia = secuencia + 1
						fmt.Println("hay una secuencia" + string(cadena[j+2]))
					}
				}
			}
			repeticionesRow = 0

			if j+2 < len(cadena) && j >= 3 {
				if string(cadena[j]) == string(cadena[j+1]) {
					repeticionesRow = repeticionesRow + 2
					if string(cadena[j-1]) == string(cadena[j+2]) && string(cadena[j]) == string(cadena[j+2]) ||
						string(cadena[j-2]) == string(cadena[j-1]) && string(cadena[j]) == string(cadena[j-1]) {
						repeticionesRow = repeticionesRow + 2
						secuencia = secuencia + 1
						fmt.Println("hay una secuencia" + string(cadena[j+2]))
					}
				}
			}
			repeticionesRow = 0

		}
		fmt.Println()
	}
	return secuencia

}

func secuenciaCol(adn [6]string) int {

	var i, j int
	var cadena, newCadena, ultCadena, antCadena string
	var repeticionesCol, secuencia int
	repeticionesCol = 0
	for i = 0; i < len(adn); i++ {
		println()
		/* fmt.Printf("iteración i %d ", i) */
		for j = 0; j < len(adn); j = j + 3 {
			cadena = adn[j]
			/* fmt.Printf("iteración j %d ", j) */
			/* fmt.Printf("iteración j %d ", j) */

			if len(adn) > j+3 {
				newCadena = adn[j+3]
				/* println("cadena primer if " + cadena + "nueva cadena primer if " + newCadena) */
				if string(cadena[i]) == string(newCadena[i]) {
					repeticionesCol = repeticionesCol + 2
					ultCadena = adn[j+2]
					newCadena = adn[j+1]
					/*  */
					if string(ultCadena[i]) == string(newCadena[i]) && string(newCadena[i]) == string(cadena[i]) {
						repeticionesCol = repeticionesCol + 2
						secuencia = secuencia + 1
						fmt.Println("hay una secuencia")
						/* fmt.Printf("En posición fila %d y posición columna %d", j, i) */
					}
				}
			}
			repeticionesCol = 0

			if j >= 3 && j+2 < len(cadena) {
				newCadena = adn[j+2]
				if string(cadena[i]) == string(newCadena[i]) {
					repeticionesCol = repeticionesCol + 2
					antCadena = cadena
					cadena = adn[j+1]
					newCadena = adn[j-1]
					ultCadena = adn[j-2]
					if string(cadena[i]) == string(newCadena[i]) && string(cadena[i]) == string(antCadena[i]) {
						repeticionesCol = repeticionesCol + 2
						secuencia = secuencia + 1
					}
				}
				cadena = adn[j-2]
				newCadena = adn[j+1]
				if string(cadena[i]) == string(newCadena[i]) {
					antCadena = cadena
					cadena = adn[j]
					newCadena = adn[j-1]
					if string(cadena[i]) == string(newCadena[i]) && string(antCadena[i]) == string(newCadena[i]) {
						repeticionesCol = repeticionesCol + 2
						secuencia = secuencia + 1
					}
				}
			}
			repeticionesCol = 0

		}
		fmt.Println()
	}
	return secuencia

}
