package main

import "fmt"

func main() {
	// arrays y Slices

	array := [5]int{0, 1, 2, 3, 4}
	//nombreVar := [length] type {valors}
	array[3] = 8
	array[0] = 90
	fmt.Println(array, len(array), cap(array)) //len es el length y cap la capacidad maxima

	//slices son arrays que no tienen una capacidad definida, son como los arrays normales en js

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en los slices
	fmt.Println(slice[0])   //imprime el indice 0
	fmt.Println(slice[:3])  //imprime hasta el 3 [EL TRES NO SE INCLUYE]
	fmt.Println(slice[2:4]) //imprime entre el 2 y 4 [EL DOS SE INCLUYE Y EL CUATRO NO]
	fmt.Println(slice[4:])  // del cuatro en adelante

	slice = append(slice, 7) //esto coloca el valor 7 al final del slice {push()}

	newSlice := []int{1, 2, 3}
	slice = append(slice, newSlice...) //agregamos todos los valores del newSlice {OBSERVA QUE USA EL SPREAD OPERATOR}
	fmt.Println(slice)

}
