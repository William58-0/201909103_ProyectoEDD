package MatrizDispersa

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

//-------------------------------------------------------------------------------------------------------------------------ESTRUCTURAS
//-------------------------------------------------------------------------------------------		    PARA LEER EL JSON
type Pedidos struct {
	Pedidos []*Principal `json:"Pedidos"`
}

type Principal struct {
	Fecha        string      `json:"Fecha"`
	Tienda       string      `json:"Tienda"`
	Departamento string      `json:"Departamento"`
	Calificacion int         `json:"Calificacion"`
	Productos    []*Producto `json:"Productos"`
}

type Producto struct {
	Nombre         string `json:"Nombre"`
	Codigo         int    `json:"Codigo"`
	Descripcion    string `json:"Descripcion"`
	Precio         string `json:"Precio"`
	Cantidad       int    `json:"Cantidad"`
	Imagen         string `json:"Imagen"`
	Almacenamiento string `json:"Almacenamiento"`
	//estos son extras
	Fecha        string `json:"Fecha"`
	Tienda       string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int    `json:"Calificacion"`
}

//--------------------------------------------------------------------------------------------			OBJETOS
type YEAR struct {
	Anio      string
	Meses     ListaM
	Siguiente *YEAR
	Anterior  *YEAR
}

type MONTH struct {
	Mes       string
	Productos []Producto
	Nodos     []NODO
	Siguiente *MONTH
	Anterior  *MONTH
	//Matriz matriz
}

type NODO struct {
	Nombre    string
	Tipo      string
	Cola      Cola
	Arriba    *NODO
	Abajo     *NODO
	Izquierda *NODO
	Derecha   *NODO
	Ultimo    *NODO //este se usa solo en valores
	URight    *NODO //este es para nodo0
	UDown     *NODO //este es para nodo0
}

func (Mes *MONTH) GetNodo(Nombre string) *NODO {
	nodo := new(NODO)
	for i := 0; i < len(Mes.Nodos); i++ {
		if Mes.Nodos[i].Nombre == Nombre {
			return &Mes.Nodos[i]
		}
	}
	return nodo
}

func (Mes *MONTH) ExisteNodo(Nombre string) bool {
	for i := 0; i < len(Mes.Nodos); i++ {
		if Mes.Nodos[i].Nombre == Nombre {
			return true
		}
	}
	return false
}

//----------------------------------------------------------------------------------------------                 LISTA
type ListaA struct {
	Primero *YEAR
	Ultimo  *YEAR
	Tamanio int
}

type ListaM struct {
	Primero *MONTH
	Ultimo  *MONTH
	Tamanio int
}

func OrdenarVec(vector []string) []string {
	//se ordenan los Meses
	var j int
	var aux string
	n := len(vector)
	for i := 1; i < n; i++ {
		j = i
		aux = vector[i]
		for j > 0 && aux < vector[j-1] {
			vector[j] = vector[j-1]
			j--
		}
		vector[j] = aux
	}
	return vector
}

//----------------------------------------------------------------------------------------------				FUNCIONES DE LISTAS
//insertar año
func (Lista *ListaA) InsertarA(Anio string) {
	nuevo := new(YEAR)
	nuevo.Anio = Anio
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

//insertar mes
func (Lista *ListaM) InsertarM(Mes string) {
	nuevo := new(MONTH)
	nuevo.Mes = Mes
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

//buscar año
func (Lista *ListaA) BuscarA(Anio string) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Anio == Anio {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}

//buscar mes
func (Lista *ListaM) BuscarM(Mes string) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Mes == Mes {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}

//--------------------------------------------------------------------------------------------------------------------COLA
type Casilla struct {
	Producto  Producto
	Siguiente *Casilla
}

type Cola struct {
	Nombre  string
	Tamanio int
	Primero *Casilla
	Ultimo  *Casilla
}

func (Cola *Cola) Insertar(Producto *Producto) {
	nuevo := new(Casilla)
	nuevo.Producto = *Producto
	if Cola.Primero != nil {
		Cola.Ultimo.Siguiente = nuevo
		Cola.Ultimo = nuevo
	} else {
		Cola.Primero = nuevo
		Cola.Ultimo = nuevo
	}
	Cola.Tamanio++
}

func (Cola *Cola) Extraer() *Casilla {
	aux := Cola.Primero
	if aux.Siguiente != nil {
		Cola.Primero = aux.Siguiente
		aux.Siguiente = nil
	}
	Cola.Tamanio--
	return aux
}

//-------------------------------------------------------------------------------------------------------------------------FUNCION
var Productos []Producto
var Meses []string

type Todo struct {
	Fechas []string `json:"Fechas"`
}

type Pedd struct {
	Productos []Producto `json:"Productos"`
}

var Todo1 Todo
var Pedd1 Pedd

var Leido bool
var Inicial []Producto

func Leer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Leer")
	Leido = false
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := Pedidos{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Meses = nil
	//Enlistar solo Productos
	for i := 0; i < len(c.Pedidos); i++ {
		//fmt.Println(strconv.Itoa(i))
		for j := 0; j < len(c.Pedidos[i].Productos); j++ {
			//Produc := c.Pedidos[i].Productos[j]
			if !AVL.ExisteProducto(c.Pedidos[i].Productos[j].Codigo) {
				j++
			} else {
				Produc := AVL.GetProducto(c.Pedidos[i].Productos[j].Codigo)
				Producto := new(Producto)
				Producto.Nombre = Produc.Nombre
				Producto.Codigo = Produc.Codigo
				Producto.Descripcion = Produc.Descripcion
				Producto.Precio = strconv.FormatFloat(Produc.Precio, 'g', 1, 64)
				Producto.Cantidad = Produc.Cantidad
				Producto.Imagen = Produc.Imagen
				Producto.Almacenamiento = Produc.Almacenamiento
				Producto.Fecha = c.Pedidos[i].Fecha
				Producto.Tienda = c.Pedidos[i].Tienda
				Producto.Departamento = c.Pedidos[i].Departamento
				Producto.Calificacion = c.Pedidos[i].Calificacion
				Productos = append(Productos, *Producto)
				//agregar el mes a Meses[]
				mes := strings.Split(c.Pedidos[i].Fecha, "-")[2] + "-" + strings.Split(c.Pedidos[i].Fecha, "-")[1]
				existe := false
				for k := 0; k < len(Meses); k++ {
					if Meses[k] == mes {
						existe = true
					}
				}
				if !existe {
					Meses = append(Meses, mes)
				}
			}
		}
	}
	Inicial = Productos
	//se ordenan los Meses
	Meses = OrdenarVec(Meses)
	Todo1.Fechas = Meses
	Pedd1.Productos = Productos
	if !Leido {
		Estructurar()
	}
}

func Actualizar() {
	fmt.Println("Actualizar")
	Leido = false
	Meses = nil
	//Enlistar solo Productos
	for i := 0; i < len(Productos); i++ {
		//agregar el mes a Meses[]
		mes := strings.Split(Productos[i].Fecha, "-")[2] + "-" + strings.Split(Productos[i].Fecha, "-")[1]
		existe := false
		for k := 0; k < len(Meses); k++ {
			if Meses[k] == mes {
				existe = true
			}
		}
		if !existe {
			Meses = append(Meses, mes)
		}
	}
	//se ordenan los Meses
	Meses = OrdenarVec(Meses)
	Todo1.Fechas = Meses
	Pedd1.Productos = Productos
	if !Leido {
		Estructurar()
	}
}

func Estructurar() {
	//se crean las estructuras
	//se crea la lista de años
	fmt.Println("Estructurar")
	anio := ""
	ListaA := new(ListaA)
	for i := 0; i < len(Meses); i++ {
		if strings.Split(Meses[i], "-")[0] != anio {
			anio = strings.Split(Meses[i], "-")[0]
			ListaA.InsertarA(anio)
		}
	}
	//se añaden los Meses a los años
	aux2 := ListaA.Primero
	for aux2 != nil {
		ListaM := new(ListaM)
		for i := 0; i < len(Meses); i++ {
			if strings.Split(Meses[i], "-")[0] == aux2.Anio {
				ListaM.InsertarM(strings.Split(Meses[i], "-")[1])
			}
		}
		aux2.Meses = *ListaM
		aux2 = aux2.Siguiente
	}
	aux2 = ListaA.Primero
	for aux2 != nil {
		aux1 := aux2.Meses.Primero
		for aux1 != nil {
			for p := 0; p < len(Productos); p++ {
				anio := strings.Split(Productos[p].Fecha, "-")[2]
				mes := strings.Split(Productos[p].Fecha, "-")[1]
				if anio == aux2.Anio && mes == aux1.Mes {
					aux1.Productos = append(aux1.Productos, Productos[p])
				}
			}
			aux1 = aux1.Siguiente
		}
		aux2 = aux2.Siguiente
	}
	//Arbol de años
	ListaA.Arbol()
	//se crean los nodos generales de cada matriz
	auxA := ListaA.Primero
	//if !Leido {
	for auxA != nil {
		auxM := auxA.Meses.Primero
		for auxM != nil {
			//se crea un slice para los dias
			var dias []string
			existeDia := false
			//se crea un slice para los departamentos
			var departamentos []string
			existeDep := false
			nodo0 := new(NODO)
			nodo0.Nombre = "Matriz_D"
			nodo0.Tipo = "Dia"
			auxM.Nodos = append(auxM.Nodos, *nodo0)
			for i := 0; i < len(auxM.Productos); i++ {
				for j := 0; j < len(dias); j++ {
					if dias[j] == strings.Split(auxM.Productos[i].Fecha, "-")[0] {
						existeDia = true
						j = len(dias) + 1
					} else {
						existeDia = false
					}
				}
				for k := 0; k < len(departamentos); k++ {
					if departamentos[k] == auxM.Productos[i].Departamento {
						existeDep = true
						k = len(departamentos) + 1
					} else {
						existeDep = false
					}
				}
				if !existeDia {
					dias = append(dias, strings.Split(auxM.Productos[i].Fecha, "-")[0])
				}
				if !existeDep {
					departamentos = append(departamentos, auxM.Productos[i].Departamento)
				}
			}
			departamentos = OrdenarVec(departamentos)
			dias = OrdenarVec(dias)
			nodoAntDia := new(NODO)
			nodoAntDep := new(NODO)
			var diasagregados []string
			var depagregados []string
			for i := 0; i < len(dias); i++ {
				if len(diasagregados) <= 0 {
					if !auxM.ExisteNodo(dias[i]) {
						nodo := new(NODO)
						nodo.Nombre = dias[i]
						nodo.Tipo = "Dia"
						nodo.Izquierda = nodo0
						nodo0.Derecha = nodo
						nodo0.URight = nodo
						nodo.Ultimo = nodo
						auxM.Nodos = append(auxM.Nodos, *nodo)
						nodoAntDia = nodo
					}
				} else {
					if !auxM.ExisteNodo(dias[i]) {
						nodo := new(NODO)
						nodo.Nombre = dias[i]
						nodo.Tipo = "Dia"
						nodo.Izquierda = nodoAntDia
						nodoAntDia.Derecha = nodo
						nodo.Ultimo = nodo
						nodo.Izquierda = nodo0.URight
						nodo0.URight = nodo
						auxM.Nodos = append(auxM.Nodos, *nodo)
						nodoAntDia = nodo
					}
				}
				diasagregados = append(diasagregados, strings.Split(auxM.Productos[i].Fecha, "-")[0])
			}
			for i := 0; i < len(departamentos); i++ {
				if len(depagregados) <= 0 {
					if !auxM.ExisteNodo(departamentos[i]) {
						nodo := new(NODO)
						nodo.Nombre = departamentos[i]
						nodo.Tipo = "Dep"
						nodo.Arriba = nodo0
						nodo0.Abajo = nodo
						nodo.Ultimo = nodo
						auxM.Nodos = append(auxM.Nodos, *nodo)
						nodoAntDep = nodo
						nodo0.UDown = nodo
					}
				} else {
					if !auxM.ExisteNodo(departamentos[i]) {
						nodo := new(NODO)
						nodo.Nombre = departamentos[i]
						nodo.Tipo = "Dep"
						nodo.Arriba = nodo0.UDown
						nodo0.UDown.Abajo = nodo
						nodoAntDep.Abajo = nodo
						nodo.Ultimo = nodo
						auxM.Nodos = append(auxM.Nodos, *nodo)
						nodoAntDep = nodo
						nodo0.UDown = nodo
					}
				}
				depagregados = append(depagregados, departamentos[i])
			}
			for i := 0; i < len(auxM.Productos); i++ {
				Day := auxM.GetNodo(strings.Split(auxM.Productos[i].Fecha, "-")[0])
				Depa := auxM.GetNodo(auxM.Productos[i].Departamento)
				buscar := Day.Nombre + Depa.Nombre
				if auxM.ExisteNodo(buscar) {
					if auxM.GetNodo(buscar).Cola.Nombre == buscar {
						auxM.GetNodo(buscar).Cola.Insertar(&auxM.Productos[i])
					}
				} else {
					nodo := new(NODO)
					cola := new(Cola)
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

			//se crea el grafo
			cadena := ""
			cadena1 := ""
			rankdir := "{ rank=same; "
			//fmt.Println("\n INFO DE NODOS: \n")
			for i := 0; i < len(auxM.Nodos); i++ {
				a := strings.ReplaceAll(auxM.Nodos[i].Nombre, " ", "_")
				//si es un nodo con cola
				if auxM.Nodos[i].Cola.Tamanio != 0 {
					cadena += "nodo" + a + " [label=\"" + strconv.Itoa(auxM.Nodos[i].Cola.Tamanio) +
						"\" shape=circle fillcolor=lightgoldenrod]\n"
				} else {
					cadena += "nodo" + a + " [label=\"" + auxM.Nodos[i].Nombre + "\" fillcolor=aquamarine]\n"
				}
				if auxM.Nodos[i].Arriba != nil {
					b := strings.ReplaceAll(auxM.Nodos[i].Arriba.Nombre, " ", "_")
					cadena += "nodo" + a + "->nodo" + b + " [dir=both]\n"
				}
				if auxM.Nodos[i].Abajo != nil {
					b := strings.ReplaceAll(auxM.Nodos[i].Abajo.Nombre, " ", "_")
					cadena += "nodo" + a + "->nodo" + b + " [dir=both]\n"
				}
				if auxM.Nodos[i].Derecha != nil {
					b := strings.ReplaceAll(auxM.Nodos[i].Derecha.Nombre, " ", "_")
					if auxM.Nodos[i].Derecha.Tipo != "Dia" {
						cadena += "nodo" + b + "->nodo" + a + " [constraint=false; dir=both]\n"
					} else {
						cadena += "nodo" + b + "->nodo" + a + "  [dir=both]\n"
					}
					cadena1 += "{ rank=same; " + "nodo" + a + "; nodo" + b + "; }\n"
				}
				if auxM.Nodos[i].Izquierda != nil {
					b := strings.ReplaceAll(auxM.Nodos[i].Izquierda.Nombre, " ", "_")
					if auxM.Nodos[i].Izquierda.Tipo != "Dia" {
						cadena += "nodo" + b + "->nodo" + a + " [constraint=false; dir=both]\n"
					} else {
						cadena += "nodo" + b + "->nodo" + a + " [dir=both]\n"
					}
					cadena1 += "{ rank=same; " + "nodo" + b + "; nodo" + a + ";}\n "
				}
				if auxM.Nodos[i].Tipo == "Dia" {
					rankdir += "nodo" + a + "; "
				}
			}
			rankdir += "}"
			cadena = "digraph {\nrankdir = BT;\nnode [shape=rectangle style=filled];\ngraph[ nodesep = 0.5];\n" +
				cadena1 + cadena + rankdir + "\n }"
			//se escribe el archivo dot
			b := []byte(cadena)
			err := ioutil.WriteFile("../frontend/src/assets/img/"+auxA.Anio+"-"+auxM.Mes+".dot", b, 0644)
			if err != nil {
				log.Fatal(err)
			}
			//se crea la imagen
			path, _ := exec.LookPath("dot")
			cmd, _ := exec.Command(path, "-Tpng", "../frontend/src/assets/img/"+auxA.Anio+"-"+auxM.Mes+".dot").Output()
			mode := int(0777)
			ioutil.WriteFile("../frontend/src/assets/img/"+auxA.Anio+"-"+auxM.Mes+".png", cmd, os.FileMode(mode))
			fmt.Println("Se creó calendario")
			fmt.Println(auxM.Mes)
			auxM = auxM.Siguiente
			fmt.Println(auxA.Anio)
		}
		auxA = auxA.Siguiente
		if auxA == nil {
			fmt.Println("Se acabó")
			Leido = true
			return
		}
	}
	//}
}

func (ListaA *ListaA) Arbol() {
	auxA := ListaA.Primero
	arbol := AVL.New_ABB()
	for auxA != nil {
		a, err := strconv.Atoi(auxA.Anio)
		if err != nil {
		}
		AVL.Insertar(arbol, auxA.Anio, a, "", 0.0, 0, "", "", "", "", 0)
		auxA = auxA.Siguiente
	}
	fmt.Println("Generando arbol de años")
	AVL.Generar_Grafo(arbol, "arbolAnios")
}

func GetFechas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todo1)
}

func GetPedidos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Pedd1)
}
