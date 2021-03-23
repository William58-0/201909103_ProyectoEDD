package AVL

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

//-------------------------------------------------------------------------Para leer el json
type Inv struct {
	Inventarios []*Principal `json:"Inventarios"`
}

type Principal struct {
	Tienda       string      `json:"Tienda"`
	Departamento string      `json:"Departamento"`
	Calificacion int         `json:"Calificacion"`
	Productos    []*Producto `json:"Productos"`
}

//------------------------------------------------------------------------Estructuras btnode
type Producto struct {
	Nombre      string  `json:"Nombre"`
	Codigo      int     `json:"Codigo"`
	Descripcion string  `json:"Descripcion"`
	Precio      float64 `json:"Precio"`
	Cantidad    int     `json:"Cantidad"`
	Imagen      string  `json:"Imagen"`
	//
	Tienda       string
	Departamento string
	Calificacion int
	//
	Izq  *Producto
	Der  *Producto
	Peso int
}

//----------------------------------------------------------------------------Para enviar el json
type Todo struct {
	Productos []Producto1 `json:"Productos"`
}

type Mostrar struct {
	Productos []Producto1 `json:"Productos"`
}

type Producto1 struct {
	Nombre       string  `json:"Nombre"`
	Codigo       int     `json:"Codigo"`
	Descripcion  string  `json:"Descripcion"`
	Precio       float64 `json:"Precio"`
	Cantidad     int     `json:"Cantidad"`
	Imagen       string  `json:"Imagen"`
	Tienda       string  `json:"Tienda"`
	Departamento string  `json:"Departamento"`
	Calificacion int     `json:"Calificacion"`
}

//-----------------------------------------------------------------------------Para buscar tienda
type Buscar struct {
	Tienda       string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int    `json:"Calificacion"`
}

var Todo1 Todo
var Mostrar1 Mostrar

func Leer(w http.ResponseWriter, r *http.Request) {
	var Produc []Producto1
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := Inv{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(c.Inventarios); i++ {
		arbol := New_ABB()
		for j := 0; j < len(c.Inventarios[i].Productos); j++ {
			Producto := c.Inventarios[i].Productos[j]
			Producto.Tienda = c.Inventarios[i].Tienda
			Producto.Departamento = c.Inventarios[i].Departamento
			Producto.Calificacion = c.Inventarios[i].Calificacion
			P := new(Producto1)
			P.Nombre = Producto.Nombre
			P.Codigo = Producto.Codigo
			P.Descripcion = Producto.Descripcion
			P.Precio = Producto.Precio
			P.Cantidad = Producto.Cantidad
			P.Imagen = Producto.Imagen
			P.Tienda = c.Inventarios[i].Tienda
			P.Departamento = c.Inventarios[i].Departamento
			P.Calificacion = c.Inventarios[i].Calificacion
			Insertar(arbol, Producto.Nombre, Producto.Codigo, Producto.Descripcion, Producto.Precio, Producto.Cantidad, Producto.Imagen, Producto.Tienda, Producto.Departamento, Producto.Calificacion)
			Produc = append(Produc, *P)
		}
		fmt.Println("Generando grafo")
		Generar_Grafo(arbol, c.Inventarios[i].Tienda+"---"+c.Inventarios[i].Departamento+"---"+strconv.Itoa(c.Inventarios[i].Calificacion))
	}
	Todo1.Productos = Produc
}

const (
	Izq_Peso   = -1
	Equilibrio = 0
	Der_Peso   = +1
)

func New_Producto(Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string, Tienda string, Departamento string, Calificacion int) *Producto {
	return &Producto{Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion, nil, nil, 0}
}

func New_Producto_2(Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string, Tienda string, Departamento string, Calificacion int, hizq *Producto, hder *Producto) *Producto {
	return &Producto{Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion, hizq, hder, 0}
}

//--------------------------------------------------------------------------Estructuras binary_tree
var increase = false

type ABB struct {
	Raiz *Producto
}

func New_ABB() *ABB {
	return &ABB{nil}
}

func New_ABB_2(hizq *ABB, hder *ABB, Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string, Tienda string, Departamento string, Calificacion int) *ABB {
	return &ABB{New_Producto_2(Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion, hizq.Raiz, hder.Raiz)}
}

func get_Izq_subarbol(arbol *ABB) *Producto {
	if arbol.Raiz == nil {
		log.Fatal("get Izq subtee on empty arbol")
		return nil
	}
	return arbol.Raiz.Izq
}

func get_Der_subarbol(arbol *ABB) *Producto {
	if arbol.Raiz == nil {
		log.Fatal("get Der subtee on empty arbol")
		return nil
	}
	return arbol.Raiz.Der
}

func get_data(arbol *ABB) int {
	return arbol.Raiz.Codigo
}

func Insertar(arbol *ABB, Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string, Tienda string, Departamento string, Calificacion int) bool {
	increase = false
	return Insertar2(&arbol.Raiz, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion, &increase)
}

func Insertar2(local_Raiz **Producto, Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string, Tienda string, Departamento string, Calificacion int, increase *bool) bool {
	if *local_Raiz == nil {
		*local_Raiz = New_Producto(Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion)
		*increase = true
		return true
	} else {
		if Codigo < (*local_Raiz).Codigo {
			return_value := Insertar2(&(*local_Raiz).Izq, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion, increase)
			if *increase {
				switch (*local_Raiz).Peso {
				case Equilibrio:
					(*local_Raiz).Peso = Izq_Peso
				case Der_Peso:
					(*local_Raiz).Peso = Equilibrio
					*increase = false
					break
				case Izq_Peso:
					Rebalance_Izq(&*local_Raiz)
					*increase = false
					break
				}
			}
			return return_value
		} else if Codigo > (*local_Raiz).Codigo {
			return_value := Insertar2(&(*local_Raiz).Der, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Departamento, Calificacion, increase)
			if *increase {
				switch (*local_Raiz).Peso {
				case Equilibrio:
					(*local_Raiz).Peso = Der_Peso
				case Izq_Peso:
					(*local_Raiz).Peso = Equilibrio
					*increase = false
					break
				case Der_Peso:
					Rebalance_Der(&*local_Raiz)
					*increase = false
					break
				}
			}
			return return_value
		} else {
			return false
		}
	}
}

func Encontrar(arbol *ABB, Codigo int) int {
	return Encontrar2(arbol.Raiz, Codigo)
}

func Encontrar2(local_Raiz *Producto, Codigo int) int {
	if local_Raiz == nil {
		return -1
	}
	if Codigo < local_Raiz.Codigo {
		return Encontrar2(local_Raiz.Izq, Codigo)
	} else if Codigo > local_Raiz.Codigo {
		return Encontrar2(local_Raiz.Der, Codigo)
	} else {
		return local_Raiz.Codigo
	}
}

func PreOrden(arbol *ABB) {
	PreOrden_2(arbol.Raiz)
}

func PreOrden_2(Raiz *Producto) {
	if Raiz != nil {
		fmt.Println(Raiz.Codigo)
		PreOrden_2(Raiz.Izq)
		PreOrden_2(Raiz.Der)
	}
}

func InOrden(arbol *ABB) {
	InOrden_2(arbol.Raiz)
}

func InOrden_2(Raiz *Producto) {
	if Raiz != nil {
		InOrden_2(Raiz.Izq)
		fmt.Println(Raiz.Codigo)
		InOrden_2(Raiz.Der)
	}
}

func Rebalance_Izq(local_Raiz **Producto) {
	hizq := (*local_Raiz).Izq
	if hizq.Peso == Der_Peso {
		hizqder := hizq.Der
		if hizqder.Peso == Izq_Peso {
			hizq.Peso = Equilibrio
			hizqder.Peso = Equilibrio
			(*local_Raiz).Peso = Der_Peso
		} else if hizqder.Peso == Equilibrio {
			hizq.Peso = Equilibrio
			hizqder.Peso = Equilibrio
			(*local_Raiz).Peso = Equilibrio
		} else {
			hizq.Peso = Izq_Peso
			hizqder.Peso = Equilibrio
			(*local_Raiz).Peso = Equilibrio
		}
		RotIzq(&(*local_Raiz).Izq)
	} else {
		hizq.Peso = Equilibrio
		(*local_Raiz).Peso = Equilibrio
	}
	RotDer(&*local_Raiz)
}

func Rebalance_Der(local_Raiz **Producto) {
	hder := (*local_Raiz).Der
	if hder.Peso == Izq_Peso {
		derhizq := hder.Izq
		if derhizq.Peso == Der_Peso {
			hder.Peso = Equilibrio
			derhizq.Peso = Equilibrio
			(*local_Raiz).Peso = Izq_Peso
		} else if derhizq.Peso == Equilibrio {
			hder.Peso = Equilibrio
			derhizq.Peso = Equilibrio
			(*local_Raiz).Peso = Equilibrio
		} else {
			hder.Peso = Der_Peso
			derhizq.Peso = Equilibrio
			(*local_Raiz).Peso = Equilibrio
		}
		RotDer(&(*local_Raiz).Der)
	} else {
		hder.Peso = Equilibrio
		(*local_Raiz).Peso = Equilibrio
	}
	RotIzq(&*local_Raiz)
}

func RotDer(local_Raiz **Producto) {
	tmp := (*local_Raiz).Izq
	(*local_Raiz).Izq = tmp.Der
	tmp.Der = *local_Raiz
	*local_Raiz = tmp
}

func RotIzq(local_Raiz **Producto) {
	tmp := (*local_Raiz).Der
	(*local_Raiz).Der = tmp.Izq
	tmp.Izq = *local_Raiz
	*local_Raiz = tmp
}

//-----------------------------------------------------------------------GRAFO
func Generar_Grafo(arbol *ABB, nombre string) {
	cadena := "digraph G{\nnode [shape=circle style=filled];\n"
	acum := ""
	if arbol.Raiz != nil {
		Recorrer_Arbol(&arbol.Raiz, &acum)
	}
	cadena += acum + "\n}\n"
	path := "../frontend/src/assets/img/" + nombre + ".dot"
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
	}

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err2) {
		return
	}
	defer file.Close()
	_, err = file.WriteString(cadena)
	if existeError(err) {
		return
	}
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Se creó grafo.")
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", "../frontend/src/assets/img/"+nombre+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile("../frontend/src/assets/img/"+nombre+".png", cmd, os.FileMode(mode))
}

func Recorrer_Arbol(actual **Producto, acum *string) {
	if *actual != nil {
		//fmt.Println("Nombre: " + (*actual).Nombre + " Codigo: " + strconv.Itoa((*actual).Codigo))
		//Producto ACTUAL
		*acum += "\"" + fmt.Sprint(&(*actual)) + "\"[label=\"" + (*actual).Nombre + "\n" + strconv.Itoa((*actual).Codigo) + "\" fillcolor=aquamarine];\n"
		//HIZQ
		if (*actual).Izq != nil {
			*acum += "\"" + fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).Izq) + "\";\n"
		}
		//HDER
		if (*actual).Der != nil {
			*acum += "\"" + fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).Der) + "\";\n"
		}
		Recorrer_Arbol(&(*actual).Izq, acum)
		Recorrer_Arbol(&(*actual).Der, acum)
	}
}

func filtrar(Tienda string, Departamento string, Calificacion int) Mostrar {
	var nn []Producto1
	for i := 0; i < len(Todo1.Productos); i++ {
		Producto := Todo1.Productos[i]
		if Tienda == Producto.Tienda && Departamento == Producto.Departamento && Calificacion == Producto.Calificacion {
			nn = append(nn, Producto)
		}
	}
	Mostrar1.Productos = nn
	return Mostrar1
}

//POST registra un nuevo curso aprobado
func GetInventario(w http.ResponseWriter, r *http.Request) {
	//var cursos[] NodoLista
	var buscar *Buscar
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &buscar)
	buscado := filtrar(buscar.Tienda, buscar.Departamento, buscar.Calificacion)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buscado)
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "William Alejandro Borrayo Alarcon_201909103")
}
