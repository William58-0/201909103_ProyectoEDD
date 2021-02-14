package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"./Estructuras"
)

func Generardot() {
	if len(Vector) == 0 {
		fmt.Println("Primero cargue un archivo")
		return
	}
	cadena := "digraph grafo{\nfontname=\"Verdana\" color=red fontsize=22;\n" +
		"node [shape=record fontsize=8 fontname=\"Verdana\" style=filled];\n" +
		"edge [color=\"blue\"]\nsubgraph cluster{\nlabel = \"Vector\";\nbgcolor=\"yellow:dodgerblue\"\n" +
		"Vector[label=\""
	for i := 1; i <= len(Vector); i++ {
		cadena = cadena + "<" + strconv.Itoa(i) + ">" + strconv.Itoa(i) + "|"
		if i%7 == 0 {
			cadena = cadena + "\n"
		}
	}
	cadena = cadena + "\",width=4.5, fillcolor=\"aquamarine:violet\"];\n}\n"
	for j := 1; j <= len(Vector); j++ {
		if Vector[j-1].Primero != nil {
			contador := 1
			Lista := Vector[j-1]
			aux := Lista.Primero
			//primer nodo
			//id del nodo=posicion en el Vector + calificacion+posicion en la lista
			cadena = cadena + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) +
				"[label=\"" + aux.Nombre + " \\n " + aux.Contacto + " \\n " +
				strconv.Itoa(aux.Calificacion) + "\", fillcolor=\"yellowgreen:aquamarine\"];\n" +
				"Vector:" + strconv.Itoa(j) + "->" + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) +
				"[color=red]\n"
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
	//se escribe el archivo dot
	b := []byte(cadena)
	err := ioutil.WriteFile("grafo.dot", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//se crea la imagen
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png", cmd, os.FileMode(mode))
	//para abrir la imagen
}

////////////////////////////////////////////////////////////////////////////////////////////      FUNCIONES

//arreglo de departamentos
var depa []string

//Vector para Linealizar
var Vector []Estructuras.Lista

func Cargar() {
	lector, err := ioutil.ReadFile("pruebas.json")
	//lector, err := ioutil.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(lector))
	c := Estructuras.Data{}
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
				Lista := new(Estructuras.Lista)
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
							}
						}
					}
				}
				//si la lista queda vacia:
				if Lista.Tamanio == 0 {
					Lista.Primero = nil
					Lista.Ultimo = nil
				}
				//se agrega la lista al Vector
				Vector = append(Vector, *Lista)
			}
		}
		let++
	}
}

//Busqueda de Posicion especifica
func BuscarPosicion(Categoria, Nombre string, Calificacion int) {
	for _, Lista := range Vector {
		if Lista.Primero != nil {
			//Si encuentra la categoria, y la calificacion
			if Lista.Categoria == Categoria && Lista.Calificacion == Calificacion {
				//buscar en la lista
				if Lista.Buscar(Nombre, Calificacion) == true {
					fmt.Println("Encontrado")
					return
				}
			}
		}
	}
	fmt.Println("No encontrado")
}

//Eliminar Registro
func Eliminar(Categoria, Nombre string, Calificacion int) {
	cont := 0
	for _, Lista := range Vector {
		if Lista.Primero != nil {
			//Si encuentra la categoria, y la calificacion
			if Lista.Categoria == Categoria && Lista.Calificacion == Calificacion {
				//buscar en la lista
				if Lista.Eliminar(Nombre, Calificacion) == true {
					fmt.Println("Encontrado")
					fmt.Println(Lista.Primero)
					Vector[cont] = Lista
					return
				}
			}
		}
		cont++
	}
	fmt.Println("No encontrado")
}

//Guardar json
/*
func Guardar() {
cadena:="{\n\"Datos\": [\n{\n"
for i:=0;i<len(Vector);i++{
	if Vector[i].Indice!=""{
		cadena=cadena+"\"Indice\":"+"\""+Vector[i].Indice+"\",\n\"Departamentos\":"
	}
}
}
*/

func listasavectores() {
	var vector []Estructuras.Dep1
	for b := 0; b < len(depa); b++ {
		Departamento := new(Estructuras.Dep1)
		Departamento.Nombre = depa[b]
		var ListaTiendas []Estructuras.Tienda1
		for a := 0; a < 5; a++ {
			indice := (b * 5) + a
			//atributos de las tiendas
			aux := Vector[indice].Primero
			for aux != nil {
				Tienda1 := new(Estructuras.Tienda1)
				Tienda1.Nombre = aux.Nombre
				Tienda1.Descripcion = aux.Descripcion
				Tienda1.Contacto = aux.Contacto
				Tienda1.Calificacion = aux.Calificacion
				if Tienda1 != nil {
					ListaTiendas = append(ListaTiendas, *Tienda1)
				}
				aux = aux.Siguiente
			}
		}
		Departamento.Tiendas = ListaTiendas
		vector = append(vector, *Departamento)
	}
	cadenaJson, err := json.Marshal(vector)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println(string(cadenaJson))
	}
}

func Guardar() {
	mascotaComoJson, err := json.Marshal(Vector)
	if err != nil {
		fmt.Printf("Error codificando mascota: %v", err)
	} else {
		fmt.Println(string(mascotaComoJson))
	}
}
func main() {
	Cargar()
	listasavectores()
}
