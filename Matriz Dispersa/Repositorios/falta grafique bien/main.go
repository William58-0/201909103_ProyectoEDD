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
	//fmt.Println("vector de productos: ")
	//fmt.Println(productos)
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
			//fmt.Println("AÑO:     ----------------           " + auxA.Anio)
			//fmt.Println("MES:     ----------------           " + auxM.Mes)
			//fmt.Println(auxM.Productos)
			//se crea un slice para los dias
			var dias []string
			existeDia := false
			//se crea un slice para los departamentos
			var departamentos []string
			existeDep := false
			nodo0 := new(Estructuras.NODO)
			nodo0.Nombre = "0"
			auxM.Nodos = append(auxM.Nodos, *nodo0)
			for i := 0; i < len(auxM.Productos); i++ {
				//fmt.Println("se analiza producto: ", auxM.Productos[i])
				for j := 0; j < len(dias); j++ {
					//fmt.Println(dias[j])
					if dias[j] == strings.Split(auxM.Productos[i].Fecha, "-")[0] {
						//fmt.Println("Dia " + dias[j] + " ya existe!")
						existeDia = true
						j = len(dias) + 1
					} else {
						existeDia = false
					}
				}
				if !existeDia {
					//fmt.Println(strings.Split(auxM.Productos[i].Fecha, "-")[0] + " no existe")
				}
				//fmt.Println("Departamento: ", auxM.Productos[i].Departamento)
				//fmt.Println(strconv.Itoa(len(departamentos)))
				for k := 0; k < len(departamentos); k++ {
					//fmt.Println(departamentos[k])
					if departamentos[k] == auxM.Productos[i].Departamento {
						fmt.Println("Dep " + departamentos[k] + " ya existe!")
						existeDep = true
						j = len(departamentos) + 1
					} else {
						existeDep = false
					}
					//fmt.Println(departamentos[k] + " no es igual a " + auxM.Productos[i].Departamento)
				}
				fmt.Println(existeDep)
				if !existeDep {
					//fmt.Println(auxM.Productos[i].Departamento + " no existe")
				}
				nodoAntDia := new(Estructuras.NODO)
				nodoAntDep := new(Estructuras.NODO)
				if !existeDia {
					if len(dias) <= 0 {
						if !auxM.ExisteNodo(strings.Split(auxM.Productos[i].Fecha, "-")[0]) {
							//fmt.Println((strings.Split(auxM.Productos[i].Fecha, "-")[0]) + " no existia antes")
							nodo := new(Estructuras.NODO)
							nodo.Nombre = strings.Split(auxM.Productos[i].Fecha, "-")[0]
							nodo.Tipo = "Dia"
							nodo.Izquierda = nodo0
							nodo0.Derecha = nodo
							nodo0.URight = nodo
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDia = nodo
						} else {
							//fmt.Println((strings.Split(auxM.Productos[i].Fecha, "-")[0]) + " ya existe")
						}
					} else {
						if !auxM.ExisteNodo(strings.Split(auxM.Productos[i].Fecha, "-")[0]) {
							//fmt.Println((strings.Split(auxM.Productos[i].Fecha, "-")[0]) + " no existia antes")
							nodo := new(Estructuras.NODO)
							nodo.Nombre = strings.Split(auxM.Productos[i].Fecha, "-")[0]
							nodo.Tipo = "Dia"
							nodo.Izquierda = nodoAntDia
							nodoAntDia.Derecha = nodo
							nodo.Ultimo = nodo
							nodo.Izquierda = nodo0.URight
							nodo0.URight = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDia = nodo
						} else {
							//fmt.Println((strings.Split(auxM.Productos[i].Fecha, "-")[0]) + " ya existe")
						}
					}
					dias = append(dias, strings.Split(auxM.Productos[i].Fecha, "-")[0])
				}
				if !existeDep {
					if len(departamentos) <= 0 {
						if !auxM.ExisteNodo(auxM.Productos[i].Departamento) {
							//fmt.Println(auxM.Productos[i].Departamento + " no existía antes")
							nodo := new(Estructuras.NODO)
							nodo.Nombre = auxM.Productos[i].Departamento
							nodo.Tipo = "Dep"
							nodo.Arriba = nodo0
							nodo0.Abajo = nodo
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDep = nodo
							nodo0.UDown = nodo
						} else {
							//fmt.Println(auxM.Productos[i].Departamento + " ya existe")
						}
					} else {
						if !auxM.ExisteNodo(auxM.Productos[i].Departamento) {
							fmt.Println(auxM.Productos[i].Departamento + " no existía antes")
							nodo := new(Estructuras.NODO)
							nodo.Nombre = auxM.Productos[i].Departamento
							nodo.Tipo = "Dep"
							nodo.Arriba = nodo0.UDown
							nodo0.UDown.Abajo = nodo ////
							nodoAntDep.Abajo = nodo
							nodo.Ultimo = nodo
							auxM.Nodos = append(auxM.Nodos, *nodo)
							nodoAntDep = nodo
							nodo0.UDown = nodo
						} else {
							//fmt.Println(auxM.Productos[i].Departamento + " ya existe")
						}
					}
					departamentos = append(departamentos, auxM.Productos[i].Departamento)
				}
				//fmt.Println(dias)
				//fmt.Println(departamentos)
			}
			//borrar luego
			//fmt.Println("EJES-----------------------------------------EJES")
			for u := 0; u < len(auxM.Nodos); u++ {
				fmt.Println("Nombre: " + auxM.Nodos[u].Nombre)
			}
			//fmt.Println("")
			//fmt.Println("PRODUCTOS-----------------------------------------PRODUCTOS")
			for u := 0; u < len(auxM.Productos); u++ {
				fmt.Println("Departamento: " + auxM.Productos[u].Departamento)
				fmt.Println("Fecha: " + auxM.Productos[u].Fecha)
			}
			fmt.Println("")
			//el nodo inicial
			//nodo0 := new(Estructuras.NODO)
			//para los valores dentro de la matriz
			//var created []string
			for i := 0; i < len(auxM.Productos); i++ {
				Day := auxM.GetNodo(strings.Split(auxM.Productos[i].Fecha, "-")[0])
				//fmt.Println("NodoNombre: " + Day.Nombre)
				if Day.Ultimo != nil {
					//fmt.Println("Ultimo: " + Day.Ultimo.Nombre)
				} else {
					//fmt.Println("NodoNombre: " + Day.Nombre + " no tiene ultimo :-(")
				}
				//fmt.Println("Nodo: " + Day.Nombre + " Ultimo: " + Day.Ultimo.Nombre)
				Depa := auxM.GetNodo(auxM.Productos[i].Departamento)
				//fmt.Println("NodoNombre: " + Depa.Nombre)
				if Depa.Ultimo != nil {
					//fmt.Println("Ultimo: " + Depa.Ultimo.Tipo)
				} else {
					//fmt.Println("NodoNombre: " + Depa.Nombre + " no tiene ultimo :-(")
				}
				//fmt.Println("Nodo: " + Depa.Nombre + " Ultimo: " + Depa.Ultimo.Nombre)
				buscar := Day.Nombre + Depa.Nombre
				if auxM.ExisteNodo(buscar) {
					auxM.GetNodo(buscar).Cola.Insertar(&auxM.Productos[i])
					//fmt.Println("busca: ")
					//fmt.Println("ENCONTRADO: ", auxM.GetNodo(buscar).Nombre)
				} else {
					nodo := new(Estructuras.NODO)
					cola := new(Estructuras.Cola)
					cola.Nombre = Day.Nombre + Depa.Nombre
					nodo.Nombre = Day.Nombre + Depa.Nombre
					nodo.Tipo = "Val"
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

			cadena := ""
			rankdir := "{ rank=same; "
			fmt.Println("\nINFO DE NODOS: \n")
			for i := 0; i < len(auxM.Nodos); i++ {
				cadena += "nodo" + auxM.Nodos[i].Nombre + " [label=\"" + auxM.Nodos[i].Nombre + "\"]\n"
				if auxM.Nodos[i].Arriba != nil {
					cadena += "nodo" + auxM.Nodos[i].Nombre + "->nodo" + auxM.Nodos[i].Arriba.Nombre + "\n"
				}
				if auxM.Nodos[i].Abajo != nil {
					cadena += "nodo" + auxM.Nodos[i].Nombre + "->nodo" + auxM.Nodos[i].Abajo.Nombre + "\n"
				}
				if auxM.Nodos[i].Derecha != nil {
					cadena += "nodo" + auxM.Nodos[i].Nombre + "->nodo" + auxM.Nodos[i].Derecha.Nombre + " [constraint=false]\n"
					cadena += "{ rank=same; " + "nodo" + auxM.Nodos[i].Nombre + "; nodo" + auxM.Nodos[i].Derecha.Nombre + "; }\n"
				}
				if auxM.Nodos[i].Izquierda != nil {
					cadena += "nodo" + auxM.Nodos[i].Nombre + "->nodo" + auxM.Nodos[i].Izquierda.Nombre + " [constraint=false]\n"
					cadena += "{ rank=same; " + "nodo" + auxM.Nodos[i].Izquierda.Nombre + "; nodo" + auxM.Nodos[i].Nombre + ";}\n "
				}
				if auxM.Nodos[i].Tipo == "Dia" {
					rankdir += "nodo" + auxM.Nodos[i].Nombre + "; "
				}
			}
			rankdir += "}"
			fmt.Println(cadena)
			fmt.Println(rankdir)
			auxM = auxM.Siguiente
		}
		auxA = auxA.Siguiente
	}
}

func main() {
	leer()
}
