package main

import (
	"fmt"
	"strconv"
)

type nodo struct {
	nombre       string
	descripcion  string
	contacto     string
	calificacion int
	siguiente    *nodo
	anterior     *nodo
}

type lista struct {
	categoria    string
	indice       string
	calificacion int
	tamanio      int
	primero      *nodo
	ultimo       *nodo
}

func (Lista *lista) insertar(nombre string, descripcion string, contacto string, calificacion int) {
	nuevo := new(nodo)
	nuevo.nombre = nombre
	nuevo.descripcion = descripcion
	nuevo.contacto = contacto
	nuevo.calificacion = calificacion
	if Lista.primero == nil {
		Lista.primero = nuevo
		Lista.ultimo = nuevo
	} else {
		Lista.ultimo.siguiente = nuevo
		Lista.ultimo.siguiente.anterior = Lista.ultimo
		Lista.ultimo = nuevo
	}
	Lista.tamanio++
}

//saber si funciona bien todavia
func (Lista *lista) eliminar(nombre string, calificacion int) {
	aux := Lista.primero
	for aux != nil {
		if aux.nombre == nombre && aux.calificacion == calificacion {
			if Lista.tamanio == 1 {
				Lista.primero = nil
				break
			}
			if aux == Lista.primero {
				Lista.primero = aux.siguiente
				aux.siguiente = nil
				Lista.primero.anterior = nil
				Lista.tamanio--
				break
			} else if aux == Lista.ultimo {
				Lista.ultimo = aux.anterior
				aux.anterior = nil
				Lista.ultimo.siguiente = nil
				Lista.tamanio--
				break
			} else {
				aux.anterior.siguiente = aux.siguiente
				aux.siguiente.anterior = aux.anterior
				Lista.tamanio--
				break
			}
		}
		aux = aux.siguiente
	}
}

//hacia adelante
func (Lista lista) mostrar() {
	aux := Lista.primero
	for aux != nil {
		fmt.Println(aux.nombre + "_" + strconv.Itoa(aux.calificacion))
		aux = aux.siguiente
	}
	fmt.Println(strconv.Itoa(Lista.tamanio))
}

//funcion para ordenar vector
func vector() {

}

func main() {
	a := 22
	b := 5
	c := a % b
	fmt.Println(c)
}

/*
//para ordenar una lista segun calificaciones de tienda
func (Lista *lista) ordenar() *lista {
	nuevalista := new(lista)
	cont := 1
	for cont <= Lista.tamanio {
		aux := Lista.primero
		for aux != nil {
			if aux.calificacion == cont {
				nuevalista.insertar(aux.nombre, aux.descripcion, aux.contacto, aux.calificacion)
			}
			aux = aux.siguiente
		}
		cont++
	}
	return nuevalista
}
*/

/*
//abecedario y comparar letras
	abecedario := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 26; i++ {
		fmt.Println(string((abecedario)[i]))
	}

	palabra := "ABCDE"
	fmt.Println(string((palabra)[1]))
*/

/*
//pruebas casi inutiles
var arreglo []lista
Lista := new(lista)
Lista.insertar("tienda1", "desc", "6666", 5)
Lista.insertar("tienda2", "desc", "6666", 4)
Lista.insertar("tienda3", "desc", "6666", 3)
Lista.insertar("tienda4", "desc", "6666", 1)
Lista.insertar("tienda5", "desc", "6666", 1)
fmt.Println("antes")
Lista.mostrar()
fmt.Println("despues")
Lista1 := Lista.ordenar()
Lista1.mostrar()
Listaex := new(lista)
arreglo = append(arreglo, *Lista)
arreglo = append(arreglo, *Lista1)
arreglo = append(arreglo, *Listaex)
fmt.Println(arreglo[0].primero)
fmt.Println(arreglo[1].ultimo)
fmt.Println(arreglo[2].primero)
*/
