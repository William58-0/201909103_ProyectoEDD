package Grafos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"../AVL"
	"../MatrizDispersa"
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

//------------------------------------------------------------------------------------------------GRAFICAR
var Nodes []*Nodo

type ENLACE struct {
	Nodo1     string
	Nodo2     string
	Distancia int
}

var Enlaces []ENLACE

func GrafoInicial(w http.ResponseWriter, r *http.Request) {
	Enlaces = nil
	c := Nodos{}
	lector, err := ioutil.ReadAll(r.Body)
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
	cadena := "digraph G\n{\nnode [shape=circle style=filled];\n"
	for i := 0; i < len(c.Nodos); i++ {
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
		}
	}
	//fmt.Println(cadena)
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
	for i := 0; i < len(Nodes); i++ {
		if Nodes[i].Nombre != Data.PosicionInicialRobot && Nodes[i].Nombre != Data.Entrega {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
				" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"blue\"];\n"
		} else if Nodes[i].Nombre == Data.PosicionInicialRobot {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
				" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"red\"];\n"
		} else if Nodes[i].Nombre == Data.Entrega {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
				" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"green\"];\n"
		}
		for j := 0; j < len(Nodes[i].Enlaces); j++ {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
				strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
				" [label=\"" +
				strconv.Itoa(Nodes[i].Enlaces[j].Distancia) +
				"\" dir=both];\n"
		}
	}
	cadena += "}"
	GenerarImagen(cadena, "GrafoInicial")
}

func GenerarImagen(cadena, Nombre string) {
	//se escribe el archivo dot
	b := []byte(cadena)
	err := ioutil.WriteFile("../frontend/src/assets/img/"+Nombre+".dot", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "../frontend/src/assets/img/"+Nombre+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile("../frontend/src/assets/img/"+Nombre+".png", cmd, os.FileMode(mode))
	fmt.Println("Se creó un grafo")
}

//------------------------------------------------------------------------------------------------CAMINO

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
		}
	}
}

func CaminoMasCorto(Inicial, Final string) Ruta {
	Route := new(Ruta)
	if Inicial == "" || Final == "" {
		return *Route
	}
	if Inicial == Final {
		return *Route
	}
	Rutas = nil
	Camino(Inicial)
	for k := 0; k < 3; k++ {
		for i := 0; i < len(Nodes); i++ {
			if Nodes[i].Nombre != Inicial &&
				Nodes[i].Nombre != Final {
				Camino(Nodes[i].Nombre)
			}
		}
	}
	if len(Rutas) == 0 {
		return *Route
	}
	Camino(Final)
	var arr []Ruta
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Actual == Final {
			arr = append(arr, Rutas[i])
		}
	}
	Rutas = arr
	menor := Rutas[0].Recorrido
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Anteriores[0] == Inicial {
			if Rutas[i].Recorrido <= menor {
				menor = Rutas[i].Recorrido
				*Route = Rutas[i]
			}
		}
	}
	fmt.Println("El camino mas corto es: " + Route.Actual)
	fmt.Print(Route.Anteriores)
	fmt.Println(Route.Actual)
	fmt.Println("Y recorrido: " + strconv.Itoa(Route.Recorrido))
	return *Route
}

func Trayectoria(Ruta Ruta) {
	var arr []string
	arr = Ruta.Anteriores
	arr = append(arr, Ruta.Actual)
	Ruta.Anteriores = arr
	for k := 0; k < len(Ruta.Anteriores); k++ {
		cadena := "digraph G\n{\nnode [shape=circle style=filled];\n"
		for i := 0; i < len(Nodes); i++ {
			if Nodes[i].Nombre != Data.PosicionInicialRobot &&
				Nodes[i].Nombre != Data.Entrega &&
				Nodes[i].Nombre != Ruta.Anteriores[k] {
				cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
					" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"blue\"];\n"
			} else if Nodes[i].Nombre == Data.PosicionInicialRobot &&
				Nodes[i].Nombre != Ruta.Anteriores[k] {
				cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
					" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"red\"];\n"
			} else if Nodes[i].Nombre == Data.Entrega &&
				Nodes[i].Nombre != Ruta.Anteriores[k] {
				cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
					" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"green\"];\n"
			} else if Nodes[i].Nombre == Ruta.Anteriores[k] {
				cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
					" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"yellow\"];\n"
			}
			for j := 0; j < len(Nodes[i].Enlaces); j++ {
				if k < len(Ruta.Anteriores)-1 &&
					((Nodes[i].Enlaces[j].Nombre == Ruta.Anteriores[k] &&
						Nodes[i].Nombre == Ruta.Anteriores[k+1]) ||
						(Nodes[i].Nombre == Ruta.Anteriores[k] &&
							Nodes[i].Enlaces[j].Nombre == Ruta.Anteriores[k+1])) {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
						strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
						" [label=\"" +
						strconv.Itoa(Nodes[i].Enlaces[j].Distancia) +
						"\" color=red];\n"
				} else {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
						strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
						" [label=\"" +
						strconv.Itoa(Nodes[i].Enlaces[j].Distancia) +
						"\" dir=both];\n"
				}
			}
		}
		cadena += "}"
		Llevar(Ruta.Anteriores[k])
		GenerarImagen(cadena, "Paso"+strconv.Itoa(NPaso))
		NPaso++
	}
	Actual = Ruta.Anteriores[len(Ruta.Anteriores)]
	EvaluarCaminos(Actual)
}

//var Pedido []AVL.Producto1
var Actual string
var NPaso int
var DepPendientes []string
var Camion []AVL.Producto1
var PorRecoger []AVL.Producto1

type Paso struct {
	Numero     int
	Recorrido  string
	Distancia  int
	Recogidos  []AVL.Producto1
	Pendientes []AVL.Producto1
}

var Pasos []Paso

func Llevar(Departamento string) {
	Paso := new(Paso)
	//Carga los Productos al Camion
	for i := 0; i < len(PorRecoger); i++ {
		if PorRecoger[i].Departamento == Departamento {
			Camion = append(Camion, PorRecoger[i])
		}
	}
	//Actualiza Productos Pendientes
	var nuevo []AVL.Producto1
	for i := 0; i < len(PorRecoger); i++ {
		if PorRecoger[i].Departamento != Departamento {
			nuevo = append(nuevo, PorRecoger[i])
		}
	}
	PorRecoger = nuevo
	//Se elimina de DepPendientes
	var new []string
	for i := 0; i < len(DepPendientes); i++ {
		if DepPendientes[i] != Departamento {
			new = append(new, DepPendientes[i])
		}
	}
	Paso.Numero = NPaso
	Paso.Recorrido = ""
	Paso.Distancia = 0
	Paso.Recogidos = Camion
	Paso.Pendientes = PorRecoger
	Pasos = append(Pasos, *Paso)
	DepPendientes = new
}

func EvaluarCaminos(Actual string) {
	if len(DepPendientes) == 0 {
		return
	}
	var Longitudes []int
	for i := 0; i < len(DepPendientes); i++ {
		a := CaminoMasCorto(Actual, DepPendientes[i]).Recorrido
		Longitudes = append(Longitudes, a)
	}
	//hallar el mas corto
	Destino := ""
	if len(Longitudes) != 0 {
		menor := Longitudes[0]
		for i := 0; i < len(Longitudes); i++ {
			if Longitudes[i] <= menor {
				menor = Longitudes[i]
				Destino = DepPendientes[i]
			}
		}
	}
	if Destino == "" {
		Destino = Data.Entrega
	}
	Trayectoria(CaminoMasCorto(Actual, Destino))
}

func SiguienteDep() {

}

func AgregarDepPendientes(Pedido []AVL.Producto1) {
	for i := 0; i < len(Pedido); i++ {
		if Pedido[i].Departamento != Data.Entrega {
			DepPendientes = append(DepPendientes, Pedido[i].Departamento)
		}
	}
	PorRecoger = Pedido
}

func IniciarRecorrido(Pedido []AVL.Producto1) {
	NPaso = 0
	DepPendientes = nil
	Camion = nil
	PorRecoger = nil
	Actual = Data.PosicionInicialRobot
	AgregarDepPendientes(Pedido)
	Llevar(Data.PosicionInicialRobot)
	EvaluarCaminos(Data.PosicionInicialRobot)
}

func EnviarPedido(w http.ResponseWriter, r *http.Request) {
	var productos []*MatrizDispersa.Producto
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &productos)
	fmt.Println("El nuevo pedido: ")
	for i := 0; i < len(productos); i++ {
		fmt.Println(productos[i].Nombre)
		fmt.Println(productos[i].Codigo)
	}
}
