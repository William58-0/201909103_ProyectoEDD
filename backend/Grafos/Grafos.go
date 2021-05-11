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
				" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"darkturquoise\"];\n"
		} else if Nodes[i].Nombre == Data.PosicionInicialRobot {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
				" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"salmon\"];\n"
		} else if Nodes[i].Nombre == Data.Entrega {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
				" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"green\"];\n"
		}
		for j := 0; j < len(Nodes[i].Enlaces); j++ {
			cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
				strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
				" [label=\"" +
				fmt.Sprintf("%g", Nodes[i].Enlaces[j].Distancia) +
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

func GetAllRutas(Primero, Siguiente string) []Ruta {
	arr = nil
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Actual == Siguiente && Rutas[i].Anteriores[0] == Primero {
			arr = append(arr, Rutas[i])
		}
	}
	return arr
}

func GetRuta(Primero, Siguiente string) Ruta {
	arr = GetAllRutas(Primero, Siguiente)
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

func Camino(Primero, Actual string) {
	if sepudo {
		for i := 0; i < len(Enlaces); i++ {
			if Enlaces[i].Nodo1 == Actual {
				Route := GetRuta(Primero, Enlaces[i].Nodo1)
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
	fmt.Println("------------------------Calculando camino mas corto")
	fmt.Println("De: " + Inicial + " a " + Final)
	Route := new(Ruta)
	Rutas = nil
	if Final == Data.Entrega {
		fmt.Println("El final es el nodo de entrega!")
	}
	if Inicial == "" || Final == "" {
		fmt.Println("No existe ese nodo")
		return *Route
	}
	if Inicial == Final {
		fmt.Println("No se puede ir de Inicial a Inicial")
		return *Route
	}
	Camino(Inicial, Inicial)
	for k := 0; k < 50; k++ {
		for i := 0; i < len(Nodes); i++ {
			if Nodes[i].Nombre != Inicial {
				Camino(Inicial, Nodes[i].Nombre)
			}
		}
	}
	Camino(Inicial, Final)
	var Repuesto = Rutas
	if len(Rutas) == 0 {
		fmt.Println("finnn")
		return *Route
	}
	var arr []Ruta
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Actual == Final {
			arr = append(arr, Rutas[i])
		}
	}
	Rutas = arr
	var arr1 []Ruta
	for i := 0; i < len(Rutas); i++ {
		if Rutas[i].Anteriores[0] == Inicial {
			arr1 = append(arr1, Rutas[i])
		}
	}
	Rutas = arr1
	if len(Rutas) == 0 {
		Rutas = Repuesto
		var arr4 []Ruta
		fmt.Println("Intentando Unir rutas")
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[i].Anteriores); j++ {
				fmt.Println("Getting Rutas: " + arr[i].Anteriores[j])
				arr3 := GetAllRutas(Inicial, arr[i].Anteriores[j])
				for k := 0; k < len(arr3); k++ {
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
			menor := arr4[0].Recorrido
			usar := new(Ruta)
			for i := 0; i < len(arr4); i++ {
				if arr4[i].Recorrido <= menor && arr4[i].Recorrido != 0 {
					*usar = arr4[i]
				}
			}

			var tttt []string
			for i := 0; i < len(arr); i++ {
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
			Route.Recorrido = CalcularDistancia(*Route)
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
	return *Route
}

func Trayectoria(Ruta Ruta) {
	if sepudo {
		var arr []string
		arr = Ruta.Anteriores
		arr = append(arr, Ruta.Actual)
		Ruta.Anteriores = arr
		for k := 0; k < len(Ruta.Anteriores); k++ {
			sumo := false
			cadena := "digraph G\n{\nnode [shape=circle style=filled];\n"
			for i := 0; i < len(Nodes); i++ {
				if Nodes[i].Nombre != Data.PosicionInicialRobot &&
					Nodes[i].Nombre != Data.Entrega &&
					Nodes[i].Nombre != Ruta.Anteriores[k] {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"darkturquoise\"];\n"
				} else if Nodes[i].Nombre == Data.PosicionInicialRobot &&
					Nodes[i].Nombre != Ruta.Anteriores[k] {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"salmon\"];\n"
				} else if Nodes[i].Nombre == Data.Entrega &&
					Nodes[i].Nombre != Ruta.Anteriores[k] {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"green\"];\n"
				} else if Nodes[i].Nombre == Ruta.Anteriores[k] {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"yellow\"];\n"
				}
				for j := 0; j < len(Nodes[i].Enlaces); j++ {
					if k < len(Ruta.Anteriores)-1 {
						if (Ruta.Anteriores[k] == Nodes[i].Nombre &&
							Ruta.Anteriores[k+1] == Nodes[i].Enlaces[j].Nombre) ||
							(Ruta.Anteriores[k] == Nodes[i].Enlaces[j].Nombre &&
								Ruta.Anteriores[k+1] == Nodes[i].Nombre) && !sumo {
							Distancia += Nodes[i].Enlaces[j].Distancia
							sumo = true
							cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
								strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
								" [label=\"" +
								fmt.Sprintf("%g", Nodes[i].Enlaces[j].Distancia) +
								"\" color=red dir=both];\n"
						} else {
							cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
								strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
								" [label=\"" +
								fmt.Sprintf("%g", Nodes[i].Enlaces[j].Distancia) +
								"\" dir=both];\n"
						}
					} else {
						cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
							strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
							" [label=\"" +
							fmt.Sprintf("%g", Nodes[i].Enlaces[j].Distancia) +
							"\" dir=both];\n"
					}
				}
			}
			cadena += "}"
			Llevar(Ruta.Anteriores[k])
			Paso := new(Paso)
			Paso.Numero = NPaso
			Paso.Distancia = Distancia
			tttt := strings.Split(Recorrido, " --> ")
			if len(tttt) != 0 {
				if tttt[len(tttt)-1] != Ruta.Anteriores[k] && Ruta.Anteriores[k] != "" {
					Recorrido += " --> " + Ruta.Anteriores[k]
				}
			}
			Paso.Recorrido = Recorrido
			Paso.Recogidos = Camion
			Paso.Pendientes = PorRecoger
			Todo1.Pasos = append(Todo1.Pasos, *Paso)
			GenerarImagen(cadena, "Paso"+strconv.Itoa(NPaso))
			NPaso++
		}
		Actual = Ruta.Anteriores[len(Ruta.Anteriores)-1]
		EvaluarCaminos(Actual)
	}
}

func CalcularDistancia(Ruta Ruta) float64 {
	Dist := 0.0
	for i := 0; i < len(Ruta.Anteriores); i++ {
		for j := 0; j < len(Enlaces); j++ {
			if i < len(Ruta.Anteriores)-1 {
				if Enlaces[j].Nodo1 == Ruta.Anteriores[i] &&
					Enlaces[j].Nodo2 == Ruta.Anteriores[i+1] {
					Dist += Enlaces[j].Distancia
				}
			} else {
				if Enlaces[j].Nodo1 == Ruta.Anteriores[i] &&
					Enlaces[j].Nodo2 == Ruta.Actual {
					Dist += Enlaces[j].Distancia
				}
			}
		}
	}
	return Dist
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
	if sepudo {
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
		fmt.Print("Intento Recoger productos\nAlmacenamientos pendientes: ")
		fmt.Println(len(AlmaPendientes))
		fmt.Print("Productos por Recoger: ")
		fmt.Println(len(PorRecoger))
	}
}

func EvaluarCaminos(Actual string) {
	if sepudo {
		if len(AlmaPendientes) == 0 {
			if !Entregado {
				fmt.Println("----------------------------Ir al despacho")
				if Actual != Data.Entrega {
					AlmaPendientes = append(AlmaPendientes, Data.Entrega)
					Trayectoria(CaminoMasCorto(Actual, Data.Entrega))
				}
				Entregado = true
			}
			if !Finalizado {
				fmt.Println("----------------------------Ir a posicion inicial")
				if Actual != Data.PosicionInicialRobot {
					AlmaPendientes = append(AlmaPendientes, Data.PosicionInicialRobot)
					Trayectoria(CaminoMasCorto(Data.Entrega, Data.PosicionInicialRobot))
				}
				RecorridoCompleto(Recorrido)
				Finalizado = true
				return
			}
		}
		if Actual == "" {
			fmt.Println("El nodo actual no existe")
		}
		var Longitudes []float64
		for i := 0; i < len(AlmaPendientes); i++ {
			if Actual != AlmaPendientes[i] {
				a := CaminoMasCorto(Actual, AlmaPendientes[i]).Recorrido
				if a != 0 {
					Longitudes = append(Longitudes, a)
				} else {
					fmt.Println("No se pueden recorridos de longitud 0")
				}
			}
		}
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
		} else if !Finalizado {
			Destino = Data.PosicionInicialRobot
			//fmt.Println("Algo raro pasó")
		}
		Cam := CaminoMasCorto(Actual, Destino)
		if Cam.Recorrido == 0 {
			//fmt.Println("Algo raro pasó")
			return
		} else {
			Trayectoria(CaminoMasCorto(Actual, Destino))
		}
	}
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
	//Llevar(Data.PosicionInicialRobot)
	EvaluarCaminos(Data.PosicionInicialRobot)
}

type Todo struct {
	Pasos []Paso `json:"Pasos"`
}

var Todo1 Todo
var Entregado bool
var Finalizado bool

func GenerarRecorrido(w http.ResponseWriter, r *http.Request) {
	Entregado = false
	Finalizado = false
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
		fmt.Println(productos[i].Cliente)
	}
	IniciarRecorrido(productos)
	fmt.Println("--------------------------Recorrido generado")
}

func GetRecorrido(w http.ResponseWriter, r *http.Request) {
	if sepudo {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Todo1)
	}
}

//Crea la ruta completa del recorrido
func RecorridoCompleto(Recorrido string) {
	if sepudo {
		Ruta := new(Ruta)
		almacenes := strings.Split(Recorrido, " --> ")
		Ruta.Actual = almacenes[0]
		for i := 0; i < len(almacenes); i++ {
			Ruta.Anteriores = append(Ruta.Anteriores, almacenes[i])
			//si no es el ultimo almacen del recorrido
			if i < len(almacenes)-1 {
				//buscar el enlace
				for j := 0; j < len(Enlaces); j++ {
					if Enlaces[j].Nodo1 == almacenes[i] &&
						Enlaces[j].Nodo2 == almacenes[i+1] {
						Ruta.Recorrido += Enlaces[j].Distancia
						Ruta.Actual = Enlaces[j].Nodo2
						j = len(Enlaces)
					}
				}
			}
		}
		TrayectoriaCompleta(*Ruta)
	}
}

func FueUsado(nodo string, nodos []string) bool {
	respuesta := false
	for i := 0; i < len(nodos); i++ {
		if nodo == nodos[i] {
			respuesta = true
			return respuesta
		}
	}
	return respuesta
}

func TrayectoriaCompleta(Ruta Ruta) {
	var escritos []string
	if sepudo {
		var arr []string
		arr = Ruta.Anteriores
		if arr[len(arr)-1] != Ruta.Actual {
			arr = append(arr, Ruta.Actual)
			Ruta.Anteriores = arr
		}
		cadena := "digraph G\n{\nnode [shape=circle style=filled];\n"
		for k := 0; k < len(Ruta.Anteriores); k++ {
			for i := 0; i < len(Nodes); i++ {
				SeUso := FueUsado(Nodes[i].Nombre, Ruta.Anteriores)
				if Nodes[i].Nombre != Data.PosicionInicialRobot &&
					Nodes[i].Nombre != Data.Entrega && SeUso {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"yellow\"];\n"
				} else if Nodes[i].Nombre == Data.PosicionInicialRobot {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"salmon\"];\n"
				} else if Nodes[i].Nombre == Data.Entrega {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"green\"];\n"
				} else {
					cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") +
						" [label=\"" + Nodes[i].Nombre + "\" fillcolor=\"darkturquoise\"];\n"
				}
				for j := 0; j < len(Nodes[i].Enlaces); j++ {
					if (strings.Contains(Recorrido, Nodes[i].Nombre+" --> "+Nodes[i].Enlaces[j].Nombre) ||
						strings.Contains(Recorrido, Nodes[i].Enlaces[j].Nombre+" --> "+Nodes[i].Nombre)) &&
						(!FueUsado(Nodes[i].Nombre+" --> "+Nodes[i].Enlaces[j].Nombre, escritos) &&
							!FueUsado(Nodes[i].Enlaces[j].Nombre+" --> "+Nodes[i].Nombre, escritos)) {
						cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
							strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
							" [label=\"" +
							fmt.Sprintf("%g", Nodes[i].Enlaces[j].Distancia) +
							"\" color=red dir=both];\n"
						escritos = append(escritos, Nodes[i].Enlaces[j].Nombre+" --> "+Nodes[i].Nombre)
						escritos = append(escritos, Nodes[i].Nombre+" --> "+Nodes[i].Enlaces[j].Nombre)
					} else if !FueUsado(Nodes[i].Nombre+" --> "+Nodes[i].Enlaces[j].Nombre, escritos) &&
						!FueUsado(Nodes[i].Enlaces[j].Nombre+" --> "+Nodes[i].Nombre, escritos) {
						cadena += strings.ReplaceAll(Nodes[i].Nombre, " ", "_") + " -> " +
							strings.ReplaceAll(Nodes[i].Enlaces[j].Nombre, " ", "_") +
							" [label=\"" +
							fmt.Sprintf("%g", Nodes[i].Enlaces[j].Distancia) +
							"\" dir=both];\n"
						escritos = append(escritos, Nodes[i].Enlaces[j].Nombre+" --> "+Nodes[i].Nombre)
						escritos = append(escritos, Nodes[i].Nombre+" --> "+Nodes[i].Enlaces[j].Nombre)
					}
				}

			}
		}
		cadena += "}"
		GenerarImagen(cadena, "RecorridoCompleto")
	}

}
