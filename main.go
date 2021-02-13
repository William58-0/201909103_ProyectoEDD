package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (Lista *Lista) eliminar(Nombre string, categoria string, Calificacion int) {
	aux := Lista.Primero
	for aux != nil {
		if aux.Nombre == Nombre && aux.Calificacion == Calificacion {
			if Lista.Tamanio == 1 {
				Lista.Primero = nil
				break
			}
			if aux == Lista.Primero {
				Lista.Primero = aux.Siguiente
				aux.Siguiente = nil
				Lista.Primero.Anterior = nil
				Lista.Tamanio--
				break
			} else if aux == Lista.Ultimo {
				Lista.Ultimo = aux.Anterior
				aux.Anterior = nil
				Lista.Ultimo.Siguiente = nil
				Lista.Tamanio--
				break
			} else {
				aux.Anterior.Siguiente = aux.Siguiente
				aux.Siguiente.Anterior = aux.Anterior
				Lista.Tamanio--
				break
			}
		}
		aux = aux.Siguiente
	}
}

////////////////////////////////////////////////////////////////////////////////////////////      GRAPHVIZ
func generardot() {
	if len(vector) == 0 {
		fmt.Println("Primero cargue un archivo")
		return
	}
	cadena := "digraph grafo{\nfontname=\"Verdana\" color=red fontsize=22;" +
		"node [shape=record fontsize=12 fontname=\"Verdana\" style=filled];" +
		"edge [color=\"blue\"]\nsubgraph cluster{\nlabel = \"Vector\";\nbgcolor=\"yellow:dodgerblue\"" +
		"vector[label=\""
	for i := 1; i <= len(vector); i++ {
		cadena = cadena + "<" + strconv.Itoa(i) + ">" + strconv.Itoa(i) + "|"
		if i%7 == 0 {
			cadena = cadena + "\n"
		}
	}
	cadena = cadena + "\",width=4.5, fillcolor=\"aquamarine:violet\"];\n}\n"
	for j := 1; j <= len(vector); j++ {
		if vector[j-1].Primero != nil {
			contador := 1
			Lista := vector[j-1]
			aux := Lista.Primero
			//primer nodo
			//id del nodo=posicion en el vector + calificacion+posicion en la lista
			cadena = cadena + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) +
				"[label=\"" + aux.Nombre + " \\n " + aux.Contacto + " \\n " +
				strconv.Itoa(aux.Calificacion) + "\", fillcolor=\"yellowgreen:aquamarine\"];\n" +
				"vector:" + strconv.Itoa(j) + "->" + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "\n"
			contador++
			aux = aux.Siguiente
			//el resto de nodos
			for aux != nil {
				cadena = cadena + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "->" +
					strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador-1) + "\n" + //nodo actual->nodo anterior
					strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador-1) + "->" +
					strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "\n" + //nodo anterior->nodo actual
					//se crea nuevo nodo
					strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "[label=\"" + aux.Nombre + " \\n " +
					aux.Contacto + strconv.Itoa(aux.Calificacion) + "\", fillcolor=\"yellowgreen:aquamarine\"];\n"
				contador++
				aux = aux.Siguiente
			}
		}
	}
	cadena = cadena + "}"
	fmt.Println(cadena)
}

////////////////////////////////////////////////////////////////////////////////////////////      FUNCIONES

//arreglo de departamentos
var depa []string

//vector para linealizar
var vector []Lista

func Linealizar() {
	lector, err := ioutil.ReadFile("pruebas.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Data{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	//Contar numero de departamentos
	for i := 0; i < len(c.Datos); i++ {
		for j := 0; j < len(c.Datos[i].Departamentos); j++ {
			existe := false
			for _, a := range depa {
				if c.Datos[i].Departamentos[j].Nombre == a {
					existe = true
				}
			}
			if existe == false {
				depa = append(depa, c.Datos[i].Departamentos[j].Nombre)
			}
		}
	}
	//alfabeto
	let := 'A'
	//indices
	for abc := 0; abc < 26; abc++ {
		//departamentos
		for _, dep := range depa {
			//calificaciones
			for calif := 1; calif <= 5; calif++ {
				//se crea la lista con las caracteristicas correspondientes
				Lista := new(Lista)
				Lista.Indice = string(let)
				Lista.Categoria = dep
				Lista.Calificacion = calif
				//recorrer datos del json para comparar
				for i := 0; i < len(c.Datos); i++ {
					for j := 0; j < len(c.Datos[i].Departamentos); j++ {
						for k := 0; k < len(c.Datos[i].Departamentos[j].Tiendas); k++ {
							prim := c.Datos[i]
							sec := prim.Departamentos[j]
							terc := sec.Tiendas[k]
							//si una tienda cargada desde el json cumple con las caracteristicas, se agrega a la lista
							if prim.Indice == string(let) && sec.Nombre == dep && terc.Calificacion == calif {
								Lista.Insertar(terc.Nombre, terc.Descripcion, terc.Contacto, terc.Calificacion)
								fmt.Println("Se inserta: " + terc.Nombre)
							}
						}
					}
				}
				//si la lista queda vacia:
				if Lista.Tamanio == 0 {
					Lista.Primero = nil
					Lista.Ultimo = nil
				}
				//se agrega la lista al vector
				vector = append(vector, *Lista)
			}
		}
		let++
	}

}

func main() {
	Linealizar()
	generardot()
}
