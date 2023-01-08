package main

import "fmt"

func main() {
	//Manejo de variables
	var numero, numero1, numero2 int
	var multiplicacion int
	multiplicacion = 25
	numero = 25
	//Imprimiendo número 1
	fmt.Println(numero)

	//Imprimiendo número 2
	numero = 5
	fmt.Println(numero)

	var nombre, nombre1, nombre2 string
	nombre = "Brandon"
	fmt.Println(nombre)

	nombre, numero = "Brandon Pérez Reyes", 23
	fmt.Println(nombre, numero)

	nombre, nombre1, nombre2, numero, numero1, numero2 = "Brandon", "Alejandro", "Luis", 23, 15, 19
	fmt.Println(nombre, nombre1, nombre2, numero, numero1, numero2)

	fmt.Println(numero * multiplicacion)
}
