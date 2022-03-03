package main

import "fmt"

type Course struct { //CREAMOS LA ESTRUCTURA
	CourseName string
	Price      float64
	IsFree     bool
	UsersID    []uint
	Classes    map[int]string
}

func main() {
	Go := Course{ //DEFINIMOS LA ESTRUCTURA CON TODOS SUS DATOS
		CourseName: "Curso de Go",
		Price:      80.32,
		IsFree:     false,
		UsersID:    []uint{1, 7, 11},
		Classes: map[int]string{
			1: "Introduccion",
			2: "Explicacion",
			3: "Examen Final",
		},
	}
	fmt.Println(Go)

	Css := Course{CourseName: "Curso de CSS", IsFree: true} //DEFINIMOS LA ESTRUCTURA CON ALGUNOS DATOS
	fmt.Println(Css)

	Js := Course{} //DEFINIMOS LA ESTRUCTURA SOLA CON LOS DATOS SEPARADOS FUERA DE LA DEFINICION
	Js.CourseName = "Curso de JavaScript"
	fmt.Println(Js)

	printClasses(Go)
}

func printClasses(c Course) { //FUNCION PARA AGRUPAR LAS CLASES DE LOS CURSOS
	text := "Las clases que tenemos son: " //DEFINIMOS EL TEXTO
	for _, class := range c.Classes {
		text += class + ", " // LE SUMAMOS LAS CLASES QUE TENEMOS A LA VARIABLE TEXTO
	}
	fmt.Println(text[:len(text)-2]) //LE RESTAMOS DOS LUGARES AL TEXTO PARA QUE NO NOS QUEDE ASI "Examen Final, "
	// 				[DESDE : HASTA]
}
