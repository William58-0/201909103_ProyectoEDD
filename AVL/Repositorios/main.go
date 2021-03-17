package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

//------------------------------------------------------------------------Estructuras btnode
type Producto struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      string
	Cantidad    int
	Imagen      string
	Izq         *Producto
	Der         *Producto
	Peso        int
}

const (
	Izq_Peso   = -1
	Equilibrio = 0
	Der_Peso   = +1
)

func New_Producto(Nombre string, Codigo int, Descripcion string, Precio string, Cantidad int, Imagen string) *Producto {
	return &Producto{Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, nil, nil, 0}
}

func New_Producto_2(Nombre string, Codigo int, Descripcion string, Precio string, Cantidad int, Imagen string, hizq *Producto, hder *Producto) *Producto {
	return &Producto{Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, hizq, hder, 0}
}

//--------------------------------------------------------------------------Estructuras binary_tree
var increase = false

type ABB struct {
	Raiz *Producto
}

func New_ABB() *ABB {
	return &ABB{nil}
}

func New_ABB_2(hizq *ABB, hder *ABB, Nombre string, Codigo int, Descripcion string, Precio string, Cantidad int, Imagen string) *ABB {
	return &ABB{New_Producto_2(Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, hizq.Raiz, hder.Raiz)}
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

func Insertar(arbol *ABB, Nombre string, Codigo int, Descripcion string, Precio string, Cantidad int, Imagen string) bool {
	increase = false
	return Insertar2(&arbol.Raiz, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, &increase)
}

func Insertar2(local_Raiz **Producto, Nombre string, Codigo int, Descripcion string, Precio string, Cantidad int, Imagen string, increase *bool) bool {
	if *local_Raiz == nil {
		*local_Raiz = New_Producto(Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen)
		*increase = true
		return true
	} else {
		if Codigo < (*local_Raiz).Codigo {
			return_value := Insertar2(&(*local_Raiz).Izq, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, increase)
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
			return_value := Insertar2(&(*local_Raiz).Der, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, increase)
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
func Generar_Grafo(arbol *ABB) {
	cadena := "digraph G{\nnode [shape=circle];\n"
	acum := ""
	if arbol.Raiz != nil {
		Recorrer_Arbol(&arbol.Raiz, &acum)
	}
	cadena += acum + "\n}\n"
	path := "grafo.dot"
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
	cmd, _ := exec.Command(path2, "-Tpng", "grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png", cmd, os.FileMode(mode))
}

func Recorrer_Arbol(actual **Producto, acum *string) {
	if *actual != nil {
		fmt.Println("Nombre: " + (*actual).Nombre + "Codigo: " + strconv.Itoa((*actual).Codigo))
		//Producto ACTUAL
		*acum += "\"" + fmt.Sprint(&(*actual)) + "\"[label=\"" + strconv.Itoa((*actual).Codigo) + "\"];\n"
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

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	arbol := New_ABB()
	/*
		Nombre
		Código
		Descripción
		Precio
		Cantidad
		Imagen
	*/
	Insertar(arbol, "Nombre1", 8, "Desc1", "Precio1", 1, "Imagen1")
	Insertar(arbol, "Nombre2", 7, "Desc2", "Precio2", 2, "Imagen2")
	Insertar(arbol, "Nombre3", 6, "Desc3", "Precio3", 3, "Imagen3")
	Insertar(arbol, "Nombre4", 4, "Desc4", "Precio4", 4, "Imagen4")
	Insertar(arbol, "Nombre5", 3, "Desc5", "Precio5", 5, "Imagen5")
	Insertar(arbol, "Nombre6", 2, "Desc6", "Precio6", 6, "Imagen6")
	Insertar(arbol, "Nombre7", 1, "Desc7", "Precio7", 7, "Imagen7")
	Generar_Grafo(arbol)
	fmt.Println("hola")
}
