package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./Estructuras"
)

var productos []Estructuras.Producto
var meses []string

func leer() {
	lector, err := ioutil.ReadFile("Pedidos.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Estructuras.Pedidos{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	//Enlistar solo productos
	for i := 0; i < len(c.Pedidos); i++ {
		for j := 0; j < len(c.Pedidos[i].Productos); j++ {
			Produc := c.Pedidos[i].Productos[j]
			Producto := new(Estructuras.Producto)
			Producto.Nombre = Produc.Nombre
			Producto.Codigo = Produc.Codigo
			Producto.Descripcion = Produc.Descripcion
			Producto.Precio = Produc.Precio
			Producto.Cantidad = Produc.Cantidad
			Producto.Fecha = c.Pedidos[i].Fecha
			Producto.Tienda = c.Pedidos[i].Tienda
			Producto.Departamento = c.Pedidos[i].Departamento
			Producto.Calificacion = c.Pedidos[i].Calificacion
			productos = append(productos, *Producto)
			//agregar el mes a meses[]

			mes := strings.Split(c.Pedidos[i].Fecha, "-")[2] + "-" + strings.Split(c.Pedidos[i].Fecha, "-")[1]
			existe := false
			for k := 0; k < len(meses); k++ {
				if meses[k] == mes {
					existe = true
				}
			}
			if !existe {
				meses = append(meses, mes)
			}
		}
	}
	fmt.Println("vector de productos: ")
	fmt.Println(productos)
	//se ordenan los meses
	var j int
	var aux string
	n := len(meses)
	for i := 1; i < n; i++ {
		j = i
		aux = meses[i]

		for j > 0 && aux < meses[j-1] {
			meses[j] = meses[j-1]
			j--
		}
		meses[j] = aux
	}
	//se crean las estructuras
	//se crea la lista de años
	anio := ""
	ListaA := new(Estructuras.ListaA)
	for i := 0; i < len(meses); i++ {
		if strings.Split(meses[i], "-")[0] != anio {
			anio = strings.Split(meses[i], "-")[0]
			ListaA.InsertarA(anio)
		}
	}
	//se añaden los meses a los años
	aux2 := ListaA.Primero
	for aux2 != nil {
		ListaM := new(Estructuras.ListaM)
		for i := 0; i < len(meses); i++ {
			if strings.Split(meses[i], "-")[0] == aux2.Anio {
				ListaM.InsertarM(strings.Split(meses[i], "-")[1])
			}
		}
		aux2.Meses = *ListaM
		aux2 = aux2.Siguiente
	}
	aux2 = ListaA.Primero
	for aux2 != nil {
		aux1 := aux2.Meses.Primero
		for aux1 != nil {
			for p := 0; p < len(productos); p++ {
				anio := strings.Split(productos[p].Fecha, "-")[2]
				mes := strings.Split(productos[p].Fecha, "-")[1]
				if anio == aux2.Anio && mes == aux1.Mes {
					aux1.Productos = append(aux1.Productos, productos[p])
				}
			}
			aux1 = aux1.Siguiente
		}
		aux2 = aux2.Siguiente
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//se crean los nodos generales de cada matriz
	auxA := ListaA.Primero
	for auxA != nil {
		auxM := auxA.Meses.Primero
		for auxM != nil {
			//se crea un slice para los dias
			var dias []string
			existeDia := false
			//se crea un slice para los departamentos
			var departamentos []string
			existeDep := false
			nodo0 := new(Estructuras.NODO)
			nodo0.Nombre = "**"
			auxM.Nodos = append(auxM.Nodos, *nodo0)
			for i := 0; i < len(auxM.Productos); i++ {
				for j := 0; j < len(dias); j++ {
					if dias[j] == strings.Split(auxM.Productos[i].Fecha, "-")[0] {
						existeDia = true
					}
				}
				for j := 0; j < len(departamentos); j++ {
					if departamentos[j] == auxM.Productos[i].Departamento {
						existeDep = true
					}
				}
				nodoAntDia := new(Estructuras.NODO)
				nodoAntDep := new(Estructuras.NODO)
				if !existeDia {
					if len(dias) <= 0 {
						if !auxM.ExisteNodo(strings.Split(auxM.Productos[i].Fecha, "_")[0]) {
							nodo := new(Estructuras.NODO)
							nodo.Nombre = strings.Split(auxM.Productos[i].Fecha, "_")[0]
							nodo.Izquierda = nodo0
							nodo0.Derecha = nodo
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDia = nodo
						}
					} else {
						if !auxM.ExisteNodo(strings.Split(auxM.Productos[i].Fecha, "_")[0]) {
							nodo := new(Estructuras.NODO)
							nodo.Nombre = strings.Split(auxM.Productos[i].Fecha, "_")[0]
							nodo.Izquierda = nodoAntDia
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDia = nodo
						}
					}
					dias = append(dias, strings.Split(auxM.Productos[i].Fecha, "-")[0])
				}
				if !existeDep {
					if len(departamentos) <= 0 {
						if !auxM.ExisteNodo(auxM.Productos[i].Departamento) {
							nodo := new(Estructuras.NODO)
							nodo.Nombre = auxM.Productos[i].Departamento
							nodo0.Abajo = nodo
							nodo.Arriba = nodo0
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDep = nodo
						}
					} else {
						if !auxM.ExisteNodo(auxM.Productos[i].Departamento) {
							nodo := new(Estructuras.NODO)
							nodo.Nombre = auxM.Productos[i].Departamento
							nodo.Arriba = nodoAntDep
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDep = nodo
						}
					}
					departamentos = append(departamentos, auxM.Productos[i].Departamento)
				}
			}
			//el nodo inicial
			//nodo0 := new(Estructuras.NODO)
			//para los valores dentro de la matriz
			//var created []string
			for i := 0; i < len(auxM.Productos); i++ {
				Day := auxM.GetNodo(strings.Split(auxM.Productos[i].Fecha, "_")[0])
				Depa := auxM.GetNodo(auxM.Productos[i].Departamento)
				buscar := Day.Nombre + Depa.Nombre
				if auxM.ExisteNodo(buscar) {
					auxM.GetNodo(buscar).Cola.Insertar(&auxM.Productos[i])
					fmt.Println("busca: ")
					fmt.Println("ENCONTRADO: ", auxM.GetNodo(buscar).Nombre)
				} else {
					nodo := new(Estructuras.NODO)
					cola := new(Estructuras.Cola)
					cola.Nombre = Day.Nombre + Depa.Nombre
					nodo.Nombre = Day.Nombre + Depa.Nombre
					cola.Insertar(&auxM.Productos[i])
					nodo.Cola = *cola
					if Day.Ultimo != nil {
						nodo.Arriba = Day.Ultimo
						Day.Ultimo.Abajo = nodo
						Day.Ultimo = nodo
					}
					if Depa.Ultimo != nil {
						nodo.Izquierda = Depa.Ultimo
						Depa.Ultimo.Derecha = nodo
						Depa.Ultimo = nodo
					}
					auxM.Nodos = append(auxM.Nodos, *nodo)
				}

			}
			fmt.Println("\nNODOS\n")
			for i := 0; i < len(auxM.Nodos); i++ {
				fmt.Println(auxM.Nodos[i].Nombre)
			}

			fmt.Println("\nINFO DE NODOS: \n")
			for i := 0; i < len(auxM.Nodos); i++ {
				fmt.Println("Nombre: " + auxM.Nodos[i].Nombre)
				if auxM.Nodos[i].Arriba != nil {
					fmt.Println("Arriba: ", auxM.Nodos[i].Arriba.Nombre)
				}
				if auxM.Nodos[i].Abajo != nil {
					fmt.Println("Abajo: ", auxM.Nodos[i].Abajo.Nombre)
				}
				if auxM.Nodos[i].Derecha != nil {
					fmt.Println("Derecha: ", auxM.Nodos[i].Derecha.Nombre)
				}
				if auxM.Nodos[i].Izquierda != nil {
					fmt.Println("Izquierda: ", auxM.Nodos[i].Izquierda.Nombre)
				}
			}

			auxM = auxM.Siguiente
		}
		auxA = auxA.Siguiente
	}
}

func main() {
	leer()
}
