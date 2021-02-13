package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

////////////////////////////////////////////////////////////////////////////////////////////      GRAPHVIZ
func generardot() {
	if len(vector) == 0 {
		fmt.Println("Primero cargue un archivo")
		return
	}

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

}
