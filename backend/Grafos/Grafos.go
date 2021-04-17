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
	Nombre    string  `json:"Nombre"`
	Distancia float64 `json:"Distancia"`
}

//------------------------------------------------------------------------------------------------GRAFICAR
var Nodes []*Nodo

type ENLACE struct {
	Nodo1     string
	Nodo2     string
	Distancia float64
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
				fmt.Sprintf("%f", Nodes[i].Enlaces[j].Distancia) +
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
var sepudo bool

type Ruta struct {
	Actual     string
	Anteriores []string
	Recorrido  float64
}

var arr []Ruta

func GetAllRutas(Siguiente string) []Ruta {
	arr = nil
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Actual == Siguiente {
			arr = append(arr, Rutas[i])
		}
	}
	return arr
}

func GetRuta(Siguiente string) Ruta {
	arr = GetAllRutas(Siguiente)
	Ruta := new(Ruta)
	if len(arr) != 0 {
		menor := arr[0].Recorrido
		for i := 0; i < len(arr); i++ {
			if arr[i].Recorrido <= menor {
				menor = arr[i].Recorrido
				*Ruta = arr[i]
			}
		}
	}
	return *Ruta
}

func Camino(Actual string) {
	if sepudo {
		for i := 0; i < len(Enlaces); i++ {
			if Enlaces[i].Nodo1 == Actual {
				Route := GetRuta(Enlaces[i].Nodo1)
				var arr = Route.Anteriores
				arr = append(arr, Actual)
				Ruta := new(Ruta)
				Ruta.Anteriores = arr
				Ruta.Actual = Enlaces[i].Nodo2
				Ruta.Recorrido = Route.Recorrido + Enlaces[i].Distancia
				if Ruta.Recorrido != 0 {
					Rutas = append(Rutas, *Ruta)
				}
			}
		}
	}
}

func CaminoMasCorto(Inicial, Final string) Ruta {
	fmt.Println("------------------------------------Camino mas corto-----------------------------")
	fmt.Println("De: " + Inicial + " a " + Final)
	Route := new(Ruta)
	Rutas = nil
	if Final == Data.Entrega {
		fmt.Println("finnn")
	}
	if Inicial == "" || Final == "" {
		fmt.Println("finnn")
		return *Route
	}
	if Inicial == Final {
		fmt.Println("finnn")
		return *Route
	}
	Camino(Inicial)
	for k := 0; k < 5; k++ {
		for i := 0; i < len(Nodes); i++ {
			if Nodes[i].Nombre != Inicial &&
				Nodes[i].Nombre != Final {
				Camino(Nodes[i].Nombre)
			}
		}
	}
	Camino(Final)
	var Repuesto = Rutas
	if len(Rutas) == 0 {
		fmt.Println("finnn")
		return *Route
	}
	fmt.Println("------------------------------------------Rutas Disponibles----------------------------")
	for i := 0; i < len(Rutas); i++ {
		fmt.Println(Rutas[i])
	}
	var arr []Ruta
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Actual == Final {
			arr = append(arr, Rutas[i])
		}
	}
	Rutas = arr
	fmt.Println("------------------------------------------Primer Filtro--------------------------------")
	for i := 0; i < len(Rutas); i++ {
		fmt.Println(Rutas[i])
	}
	var arr1 []Ruta
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Anteriores[0] == Inicial {
			fmt.Println("Son iguales")
			arr1 = append(arr1, Rutas[i])
		}
	}
	Rutas = arr1
	fmt.Println("------------------------------------------Segundo Filtro--------------------------------")
	for i := 0; i < len(Rutas); i++ {
		fmt.Println(Rutas[i])
	}
	if len(Rutas) == 0 {
		Rutas = Repuesto
		var arr4 []Ruta
		fmt.Println("metodo Extremo")
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[i].Anteriores); j++ {
				fmt.Println("Getting Rutas: " + arr[i].Anteriores[j])
				arr3 := GetAllRutas(arr[i].Anteriores[j])
				for k := 0; k < len(arr3); k++ {
					fmt.Println(arr3[k].Anteriores[0])
					if arr3[k].Anteriores[0] == Inicial {
						arr4 = append(arr4, arr3[k])
					}
				}
			}
		}
		if len(arr4) == 0 {
			fmt.Println("-----------------------------------NO SE PUDO CALCULAR EL RECORRIDO :(---------------")
			sepudo = false

		} else {
			fmt.Println("Reporte")
			menor := arr4[0].Recorrido
			usar := new(Ruta)
			for i := 0; i < len(arr4); i++ {
				if arr4[i].Recorrido <= menor && arr4[i].Recorrido != 0 {
					*usar = arr4[i]
				}
			}
			fmt.Println("El menor es: ")
			fmt.Println(usar)

			var tttt []string
			for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i])
				for j := 0; j < len(arr[i].Anteriores); j++ {
					if usar.Actual == arr[i].Anteriores[j] {
						tttt = usar.Anteriores
						for k := j; k < len(arr[i].Anteriores); k++ {
							tttt = append(tttt, arr[i].Anteriores[k])
						}
						break
					}
				}
			}
			Route.Anteriores = tttt
			Route.Actual = Final
			Route.Recorrido = CalcularRecorrido(*Route)
			Rutas = append(Rutas, *Route)
		}
	} else {
		menor := Rutas[0].Recorrido
		for i := 0; i < len(Rutas); i++ {
			if Rutas[i].Anteriores[0] == Inicial && Rutas[i].Recorrido != 0 {
				if Rutas[i].Recorrido <= menor {
					menor = Rutas[i].Recorrido
					*Route = Rutas[i]
				}
			}
		}
	}
	fmt.Println("El camino mas corto es: " + Route.Actual)
	fmt.Print(Route.Anteriores)
	fmt.Println(Route.Actual)
	fmt.Println("Y recorrido: " + fmt.Sprintf("%f", Route.Recorrido))
	return *Route
}

func CalcularRecorrido(Ruta Ruta) float64 {
	rec := 0.0
	for i := 0; i < len(Ruta.Anteriores); i++ {
		for j := 0; j < len(Enlaces); j++ {
			if i < len(Ruta.Anteriores)-1 {
				if Enlaces[j].Nodo1 == Ruta.Anteriores[i] &&
					Enlaces[j].Nodo2 == Ruta.Anteriores[i+1] {
					rec += Enlaces[j].Distancia
				}
			} else {
				if Enlaces[j].Nodo1 == Ruta.Anteriores[i] &&
					Enlaces[j].Nodo2 == Ruta.Actual {
					rec += Enlaces[j].Distancia
				}
			}
		}
	}
	return rec
}

func Trayectoria(Ruta Ruta) {
	if sepudo {
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
							fmt.Sprintf("%f", Nodes[i].Enlaces[j].Distancia) +
							"\" color=red];\n"
					} else {
						cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
							strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
							" [label=\"" +
							fmt.Sprintf("%f", Nodes[i].Enlaces[j].Distancia) +
							"\" dir=both];\n"
					}
				}
			}
			cadena += "}"
			Llevar(Ruta.Anteriores[k])
			GenerarImagen(cadena, "Paso"+strconv.Itoa(NPaso))
			NPaso++
		}
		Actual = Ruta.Anteriores[len(Ruta.Anteriores)-1]
		EvaluarCaminos(Actual)
	}
}

//var Pedido []AVL.Producto1
var Actual string
var NPaso int
var AlmaPendientes []string
var Recorrido string
var Distancia float64
var Camion []AVL.Producto1
var PorRecoger []AVL.Producto1

type Paso struct {
	Numero     int
	Recorrido  string
	Distancia  float64
	Recogidos  []AVL.Producto1
	Pendientes []AVL.Producto1
}

func Llevar(Almacenamiento string) {
	//if sepudo {
	//Carga los Productos al Camion
	for i := 0; i < len(PorRecoger); i++ {
		if PorRecoger[i].Almacenamiento == Almacenamiento {
			Camion = append(Camion, PorRecoger[i])
		}
	}
	//Actualiza Productos Pendientes
	var nuevo []AVL.Producto1
	for i := 0; i < len(PorRecoger); i++ {
		if PorRecoger[i].Almacenamiento != Almacenamiento {
			nuevo = append(nuevo, PorRecoger[i])
		}
	}
	PorRecoger = nuevo
	//Se elimina de AlmaPendientes
	var nuevoo []string
	for i := 0; i < len(AlmaPendientes); i++ {
		if AlmaPendientes[i] != Almacenamiento {
			nuevoo = append(nuevoo, AlmaPendientes[i])
		}
	}
	AlmaPendientes = nuevoo
	Paso := new(Paso)
	Paso.Numero = NPaso
	Paso.Recorrido = Recorrido
	Paso.Distancia = Distancia
	Paso.Recogidos = Camion
	Paso.Pendientes = PorRecoger
	fmt.Println("Pasooooooooooooooooooooooooo a Recoger")
	fmt.Println(AlmaPendientes)
	fmt.Println("por recogeer")
	for i := 0; i < len(PorRecoger); i++ {
		fmt.Println(PorRecoger[i].Nombre)
		fmt.Println(PorRecoger[i].Almacenamiento)
	}
	Todo1.Pasos = append(Todo1.Pasos, *Paso)
	//}
}

func EvaluarCaminos(Actual string) {
	//if sepudo {
	if len(AlmaPendientes) == 0 {
		fmt.Println("finmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmnmnmnmnmnm")
		return
	}
	if Actual == "" {
		fmt.Println("RAROOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	}
	var Longitudes []float64
	for i := 0; i < len(AlmaPendientes); i++ {
		if Actual != AlmaPendientes[i] {
			a := CaminoMasCorto(Actual, AlmaPendientes[i]).Recorrido
			if a != 0 {
				Longitudes = append(Longitudes, a)
			} else {
				fmt.Println("Salio un ceroooooooooooooooooooooooooooooooooooooooo")
			}
		}
	}
	fmt.Println("Longituddddddddddddddddddddddddddddddddddddddddddddes: ")
	fmt.Println(Longitudes)
	//hallar el mas corto
	Destino := ""
	if len(Longitudes) != 0 {
		menor := Longitudes[0]
		for i := 0; i < len(Longitudes); i++ {
			if Longitudes[i] <= menor && Actual != AlmaPendientes[i] && Longitudes[i] != 0 {
				menor = Longitudes[i]
				Destino = AlmaPendientes[i]
			}
		}
		if Destino != "" {
			Recorrido += " --> " + Destino
			Distancia += menor
		}
	} else {
		Destino = Data.Entrega
		fmt.Println("Llego a estoooooooooooooooooooooooooooooooooooooooooooooooooooooo")
	}
	Cam := CaminoMasCorto(Actual, Destino)
	if Cam.Recorrido == 0 {
		fmt.Println("FIIIIIIINNNNNNNNNNNNNNNNNNNNNNNNNNNN")
		return
	} else {
		Trayectoria(CaminoMasCorto(Actual, Destino))
	}
	//}
}

func AgregarAlmaPendientes(Pedido []AVL.Producto1) {
	var agregados []string
	for i := 0; i < len(Pedido); i++ {
		if Pedido[i].Almacenamiento != Data.Entrega {
			existe := false
			for j := 0; j < len(agregados); j++ {
				if agregados[j] == Pedido[i].Almacenamiento {
					existe = true
					j = len(agregados)
				}
			}
			if !existe {
				AlmaPendientes = append(AlmaPendientes, Pedido[i].Almacenamiento)
				agregados = append(agregados, Pedido[i].Almacenamiento)
			}
		}
	}
	//comprobar
	fmt.Println("Almaceeeenes")
	for i := 0; i < len(AlmaPendientes); i++ {
		fmt.Println(AlmaPendientes[i])
	}
	PorRecoger = Pedido
}

func IniciarRecorrido(Pedido []AVL.Producto1) {
	NPaso = 0
	AlmaPendientes = nil
	sepudo = true
	Camion = nil
	PorRecoger = nil
	Distancia = 0
	Recorrido = Data.PosicionInicialRobot
	Actual = Data.PosicionInicialRobot
	AgregarAlmaPendientes(Pedido)
	Llevar(Data.PosicionInicialRobot)
	EvaluarCaminos(Data.PosicionInicialRobot)
}

type Todo struct {
	Pasos []Paso `json:"Pasos"`
}

var Todo1 Todo

func GenerarRecorrido(w http.ResponseWriter, r *http.Request) {
	var productos []AVL.Producto1
	Todo1 = *new(Todo)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &productos)
	fmt.Println("Productos recibidos")
	for i := 0; i < len(productos); i++ {
		fmt.Println(productos[i].Codigo)
	}
	IniciarRecorrido(productos)
	//COmprobar
	fmt.Println("COmprobarrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
	for i := 0; i < len(Todo1.Pasos); i++ {
		fmt.Println("Paso: " + strconv.Itoa(Todo1.Pasos[i].Numero))
		fmt.Println("Recorrido: " + Todo1.Pasos[i].Recorrido)
		fmt.Print(Todo1.Pasos[i].Distancia)
		fmt.Println("Pendientes: ")
		for j := 0; j < len(Todo1.Pasos[i].Pendientes); j++ {
			fmt.Println(Todo1.Pasos[i].Pendientes[j].Almacenamiento)
		}
		fmt.Println("Cargados: ")
		for j := 0; j < len(Todo1.Pasos[i].Recogidos); j++ {
			fmt.Println(Todo1.Pasos[i].Recogidos[j].Almacenamiento)
		}
	}

}

func GetRecorrido(w http.ResponseWriter, r *http.Request) {
	if sepudo {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Todo1)
	}
}
