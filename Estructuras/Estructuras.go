package Estructuras

import (
	"fmt"
	"strconv"
)

/////////////////////////////////////////////////////////////////////                   ESTRUCTURAS ORDENADAS
type Data struct {
	Datos []*Principal `json:"Datos"`
}

type Principal struct {
	Indice        string `json:"Indice"`
	Departamentos []*Dep `json:"Departamentos"`
}

type Dep struct {
	Nombre  string    `json:"Nombre"`
	Tiendas []*Tienda `json:"Tiendas"`
}

//un sinonimo de nodo
type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Siguiente    *Tienda
	Anterior     *Tienda
}

//////////////////////////////////////////////////////////////////////////					prueba
type Principal1 struct {
	Indice        string
	Departamentos []*Dep1
}

type Dep1 struct {
	Nombre  string
	Tiendas []Tienda1
}

//un sinonimo de nodo
type Tienda1 struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
}

/////////////////////////////////////////////////////////////////////////                 LISTA

//lista doblemente enlazada
type Lista struct {
	Categoria    string
	Indice       string
	Calificacion int
	Tamanio      int
	Primero      *Tienda
	Ultimo       *Tienda
}

func (Lista *Lista) Insertar(Nombre string, Descripcion string, Contacto string, Calificacion int) {
	nuevo := new(Tienda)
	nuevo.Nombre = Nombre
	nuevo.Descripcion = Descripcion
	nuevo.Contacto = Contacto
	nuevo.Calificacion = Calificacion
	if Lista.Primero == nil {
		Lista.Primero = nuevo
		Lista.Ultimo = nuevo
	} else {
		Lista.Ultimo.Siguiente = nuevo
		Lista.Ultimo.Siguiente.Anterior = Lista.Ultimo
		Lista.Ultimo = nuevo
	}
	Lista.Tamanio++
}

func (Lista *Lista) Eliminar(Nombre string, Calificacion int) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Nombre == Nombre && aux.Calificacion == Calificacion {
			fmt.Println(Nombre)
			if Lista.Tamanio == 1 {
				Lista.Primero = nil
				Lista.Tamanio--
				return true
			}
			if aux == Lista.Primero {
				fmt.Println("Primero: " + Lista.Primero.Nombre)
				Lista.Primero = aux.Siguiente
				fmt.Println("Siguiente: " + aux.Siguiente.Nombre)
				aux.Siguiente = nil
				fmt.Println("Primero: " + Lista.Primero.Nombre)
				Lista.Primero.Anterior = nil
				fmt.Println("Primero: " + Lista.Primero.Nombre)
				Lista.Tamanio--
				return true
			} else if aux == Lista.Ultimo {
				Lista.Ultimo = aux.Anterior
				aux.Anterior = nil
				Lista.Ultimo.Siguiente = nil
				Lista.Tamanio--
				return true
			} else {
				aux.Anterior.Siguiente = aux.Siguiente
				aux.Siguiente.Anterior = aux.Anterior
				Lista.Tamanio--
				return true
			}
		}
		aux = aux.Siguiente
	}
	return false
}

/*
func (Lista *Lista) Eliminar(Nombre string, Calificacion int) bool {
	aux := Lista.Primero
	NuevaLista:=new(*Lista)
	for aux != nil {
		if aux.Nombre == Nombre && aux.Calificacion == Calificacion {
			fmt.Println("Encontrado")
			aux1 := Lista.Primero
			for aux1 != nil {
				if aux1 != aux {
					NuevaLista.Insertar(aux1.Nombre, aux1.Descripcion, aux1.Contacto, aux1.Calificacion)
				}
				aux1 = aux1.Siguiente
			}
			return *NuevaLista
		}
		aux = aux.Siguiente
	}
	fmt.Println("No se encontro")
	return *NuevaLista
}*/

func (Lista *Lista) Buscar(Nombre string, Calificacion int) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Nombre == Nombre && aux.Calificacion == Calificacion {
			fmt.Println("Nombre: " + aux.Nombre)
			fmt.Println("Descripcion: " + aux.Descripcion)
			fmt.Println("Contacto: " + aux.Contacto)
			fmt.Println("Calificacion: " + strconv.Itoa(aux.Calificacion))
			return true
		}
		aux = aux.Siguiente
	}
	fmt.Println("No existe en esta lista")
	return false
}
