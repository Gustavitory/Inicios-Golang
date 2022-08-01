package main

import "fmt"

//clase 1
// func main() {
// 	/************************************
// 	           VARIABLES
// 	*************************************/
// 	// Declaracion de constantes: valor que nunca va cambiar en el tiempo
// 	//no es necesario declarar el tipo siempre
// 	const pi float64 = 3.14
// 	const pi2 = 3.1416

// 	//delclaracion de variables enteras
// 	base := 12          // el := significa que es una variable que no he declarado pero que a partir de ahora la puedo usar y tiene ese valor
// 	var altura int = 14 //forma trandicional
// 	var area int        //se declara la variable (espacio en memoria) y luego se le puede asignar un valor
// 	area = 12 * 14

// 	//Zero values
// 	var a int     //0
// 	var b float64 //0
// 	var c string  //''
// 	var d bool    //false

// 	//Area de un cuadrado
// 	const baseCuadrado = 10
// 	areaCuadrado := baseCuadrado * baseCuadrado //el solo determina el tipo que es

// 	fmt.Println("pi:", pi)
// 	fmt.Println("pi2:", pi2)
// 	fmt.Println("base:", base)
// 	fmt.Println("altura:", altura)
// 	fmt.Println("area:", area)
// 	fmt.Println(a, c, b, d)
// 	fmt.Println("Area del cuadrado es:", areaCuadrado)

// 	/************************************
// 	        OPERADORES ARITMETICOS
// 	*************************************/

// 	x := 10
// 	y := 50

// 	//suma
// 	result := x + y
// 	fmt.Println("suma:", result)

// 	//resta

// 	result = y - x
// 	fmt.Println("resta:", result)

// 	//multiplicacion
// 	result = y * x
// 	fmt.Println("multiplicacion:", result)

// 	//division
// 	result = y / x
// 	fmt.Println("division:", result)

// 	// Modulo o residuo
// 	result = y % x
// 	fmt.Println("residuo:", result)

// 	// Incrementar
// 	x++
// 	fmt.Println("incremental:", x)

// 	//decremental
// 	x--
// 	fmt.Println("decremental:", x)

// 	/************************************
// 	    TIPOS DE DATOS PRIMITIVOS
// 	*************************************/
// 	//Numeros enteros
// 	//int = Depende del OS (32 o 64 bits)
// 	//int8 = 8bits = -128 a 127
// 	//int16 = 16bits = -2^15 a 2^15-1
// 	//int32 = 32bits = -2^31 a 2^31-1
// 	//int64 = 64bits = -2^63 a 2^63-1

// 	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
// 	//uint = Depende del OS (32 o 64 bits)
// 	//uint8 = 8bits = 0 a 127
// 	//uint16 = 16bits = 0 a 2^15-1
// 	//uint32 = 32bits = 0 a 2^31-1
// 	//uint64 = 64bits = 0 a 2^63-1

// 	//numeros decimales
// 	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
// 	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308  por lo general este

// 	//textos y booleanos
// 	//string = ""
// 	//bool = true or false

// 	//numeros complejos
// 	//Complex64 = Real e Imaginario float32
// 	//Complex128 = Real e Imaginario float64
// 	//Ejemplo : c:=10 + 8i

// 	/*
// 		funciones del paquete fmt
// 	*/

// 	helloMessage := "Hello"
// 	worldMessage := "world"

// 	//PrintLn
// 	fmt.Println(helloMessage, worldMessage) //agrega un salto de linea de forma automatica

// 	//Printf
// 	nombre := "platzi"
// 	cursos := 500

// 	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
// 	//%v cualquier tipo NO RECOMENDADO
// 	//%s string
// 	//%d entero

// 	//Sprintf
// 	//para guardar strings que consideren variables en una variable
// 	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)

// 	fmt.Println(message)

// 	//tipo de datos:
// 	fmt.Printf("El Nombre:%T",nombre)
// 	fmt.Printf("Los Cursos:%T",cursos)

// }

/************************************************************************************************************************
*************************************************************************************************************************/

//clase 2

func normalFunc(message string) {
	fmt.Println(message)
}
func complexFunc(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 3
}

func main() {
	normalFunc("Hola pequeñin")

	complexFunc(1, 2, "Compleja")

	returnValue(2)

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(3)

	_, value3 := doubleReturn(2)

	fmt.Println(value1, value2)
	fmt.Println(value3)
}
