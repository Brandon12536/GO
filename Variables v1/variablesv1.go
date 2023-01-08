package main

import "fmt"

var (
	caricatura = "Dragon Ball Z"
	enemigo    = "Freezer"
)

func main() {
	//nombrar variables en GO
	numero := 25
	_nombre := "Brandon"
	numero2 := 23
	nombreUsuario := "admin"
	fmt.Println(numero, _nombre, numero2, nombreUsuario)

	fmt.Println(caricatura, enemigo)
	fmt.Println("La caricatura que sale todos los días es: "+caricatura, "su enemigo es: "+enemigo)
	imprimir()
}

func imprimir() {
	fmt.Println(caricatura, enemigo)
}
