package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

//--------------------------------------------------------------------------------------------Cargar Datos
type Nodos struct {
	Nodos                []*Nodo `json:"Nodos"`
	PosicionInicialRobot string  `json:"PosicionInicialRobot"`
	Entrega              string  `json:"Entrega"`
}

var Data Nodos

type Nodo struct {
	Nombre  string    `json:"Nombre"`
	Enlaces []*Enlace `json:"Enlaces"`
}

type Enlace struct {
	Nombre    string `json:"Nombre"`
	Distancia int    `json:"Distancia"`
}

//------------------------------------------------------------------------------------------------
type ENLACE struct {
	Nodo1     string
	Nodo2     string
	Distancia int
}

var Enlaces []ENLACE

//------------------------------------------------------------------------------------------------
var Nodes []*Nodo

func GrafoInicial() {
	c := Nodos{}
	lector, err := ioutil.ReadFile("Grafo.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Data = c
	Nodes = c.Nodos
	//Graficar
	cadena := "digraph G\n{\nnode [shape=circle];\n"
	for i := 0; i < len(c.Nodos); i++ {
		cadena += c.Nodos[i].Nombre + " [label=\"" + c.Nodos[i].Nombre + "\"];\n"
		for j := 0; j < len(c.Nodos[i].Enlaces); j++ {
			Enlace1 := new(ENLACE)
			Enlace1.Nodo1 = c.Nodos[i].Enlaces[j].Nombre
			Enlace1.Nodo2 = c.Nodos[i].Nombre
			Enlace1.Distancia = c.Nodos[i].Enlaces[j].Distancia
			Enlace2 := new(ENLACE)
			Enlace2.Nodo1 = c.Nodos[i].Nombre
			Enlace2.Nodo2 = c.Nodos[i].Enlaces[j].Nombre
			Enlace2.Distancia = c.Nodos[i].Enlaces[j].Distancia
			Enlaces = append(Enlaces, *Enlace1)
			Enlaces = append(Enlaces, *Enlace2)
			cadena += c.Nodos[i].Enlaces[j].Nombre + " -> " +
				c.Nodos[i].Nombre +
				" [label=\"" +
				strconv.Itoa(c.Nodos[i].Enlaces[j].Distancia) +
				"\"];\n"
		}
	}
	cadena += "}"
	fmt.Println(cadena)
	var aux []*Nodo
	Despacho := new(Nodo)
	for i := 0; i < len(Nodes); i++ {
		if Nodes[i].Nombre == Data.Entrega {
			Despacho = Nodes[i]
		} else {
			aux = append(aux, Nodes[i])
		}
	}
	aux = append(aux, Despacho)
	Nodes = aux
}

var Rutas []Ruta

type Ruta struct {
	Actual     string
	Anteriores []string
	Recorrido  int
}

func GetRuta(Siguiente string) Ruta {
	Ruta := new(Ruta)
	if len(Rutas) != 0 {
		menor := Rutas[0].Recorrido
		for i := 0; i < len(Rutas); i++ {
			if Rutas[i].Actual == Siguiente && Rutas[i].Recorrido <= menor {
				menor = Rutas[i].Recorrido
				*Ruta = Rutas[i]
			}
		}
	}
	return *Ruta
}

func Camino(Actual string) {
	for i := 0; i < len(Enlaces); i++ {
		if Enlaces[i].Nodo1 == Actual {
			Route := GetRuta(Enlaces[i].Nodo1)
			var arr []string
			for k := 0; k < len(Route.Anteriores); k++ {
				arr = append(arr, Route.Anteriores[k])
			}
			arr = append(arr, Actual)
			Ruta := new(Ruta)
			Ruta.Anteriores = arr
			Ruta.Actual = Enlaces[i].Nodo2
			Ruta.Recorrido = Route.Recorrido + Enlaces[i].Distancia
			Rutas = append(Rutas, *Ruta)
			fmt.Println("Funciona")
		}
	}
}

func CaminoMasCorto() {
	Camino(Data.PosicionInicialRobot)
	for k := 0; k < 3; k++ {
		for i := 0; i < len(Nodes); i++ {
			if Nodes[i].Nombre != Data.PosicionInicialRobot &&
				Nodes[i].Nombre != Data.Entrega {
				Camino(Nodes[i].Nombre)
			}
		}
	}
	if len(Rutas) == 0 {
		return
	}
	Camino(Data.Entrega)
	var arr []Ruta
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Actual == Data.Entrega {
			arr = append(arr, Rutas[i])
		}
	}
	Rutas = arr
	menor := Rutas[0].Recorrido
	Correcto := new(Ruta)
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Anteriores[0] == Data.PosicionInicialRobot {
			if Rutas[i].Recorrido <= menor {
				menor = Rutas[i].Recorrido
				*Correcto = Rutas[i]
			}
		}
	}
	fmt.Println("El camino mas corto es: " + Correcto.Actual)
	fmt.Print(Correcto.Anteriores)
	fmt.Println(Correcto.Actual)
	fmt.Println("Y recorrido: " + strconv.Itoa(Correcto.Recorrido))
}

func main() {
	GrafoInicial()
	fmt.Println(len(Nodes))
	fmt.Println(Data.PosicionInicialRobot)
	fmt.Println(Data.Entrega)
	CaminoMasCorto()
}
