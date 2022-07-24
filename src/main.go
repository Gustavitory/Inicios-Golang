package main

import "fmt"

func main() {
	/************************************
	           VARIABLES
	*************************************/
	// Declaracion de constantes: valor que nunca va cambiar en el tiempo
	//no es necesario declarar el tipo siempre
	const pi float64 = 3.14
	const pi2 = 3.1416

	//delclaracion de variables enteras
	base := 12          // el := significa que es una variable que no he declarado pero que a partir de ahora la puedo usar y tiene ese valor
	var altura int = 14 //forma trandicional
	var area int        //se declara la variable (espacio en memoria) y luego se le puede asignar un valor
	area = 12 * 14

	//Zero values
	var a int     //0
	var b float64 //0
	var c string  //''
	var d bool    //false

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado //el solo determina el tipo que es

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)
	fmt.Println(a, c, b, d)
	fmt.Println("Area del cuadrado es:", areaCuadrado)

	/************************************
	        OPERADORES ARITMETICOS
	*************************************/

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("suma:", result)

	//resta

	result = y - x
	fmt.Println("resta:", result)

	//multiplicacion
	result = y * x
	fmt.Println("multiplicacion:", result)

	//division
	result = y / x
	fmt.Println("division:", result)

	// Modulo o residuo
	result = y % x
	fmt.Println("residuo:", result)

	// Incrementar
	x++
	fmt.Println("incremental:", x)

	//decremental
	x--
	fmt.Println("decremental:", x)

	/************************************
	    TIPOS DE DATOS PRIMITIVOS
	*************************************/

}
