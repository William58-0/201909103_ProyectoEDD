package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"./Estructuras"
	"github.com/gorilla/mux"
)

func Generardot(w http.ResponseWriter, r *http.Request) {
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

func Cargar(w http.ResponseWriter, r *http.Request) {
	lector, err := ioutil.ReadAll(r.Body)
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
func BuscarPosicion(w http.ResponseWriter, r *http.Request) {
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(lector))
	c := Estructuras.Objetivo{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Nombre := c.Nombre
	Departamento := c.Departamento
	Calificacion := c.Calificacion
	for _, Lista := range Vector {
		if Lista.Primero != nil {
			//Si encuentra la categoria, y la calificacion
			if Lista.Categoria == Departamento && Lista.Calificacion == Calificacion {
				//buscar en la lista
				a := Lista.Buscar(Nombre, Calificacion)
				if a.Nombre != "" {
					fmt.Println("Encontrado")
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(a)
					return
				}
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("No encontrado")
}

//Eliminar Registro
func Eliminar(w http.ResponseWriter, r *http.Request) {
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(lector))
	c := Estructuras.Objetivo{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Nombre := c.Nombre
	Departamento := c.Departamento
	Calificacion := c.Calificacion
	cont := 0
	for _, Lista := range Vector {
		if Lista.Primero != nil {
			//Si encuentra la categoria, y la calificacion
			if Lista.Categoria == Departamento && Lista.Calificacion == Calificacion {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("No encontrado")
}

func Buscarenvector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	obj, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	var lista []Estructuras.Salida
	aux := Vector[obj].Primero
	for aux != nil {
		a := new(Estructuras.Salida)
		a.Nombre = aux.Nombre
		a.Descripcion = aux.Descripcion
		a.Contacto = aux.Descripcion
		a.Calificacion = aux.Calificacion
		lista = append(lista, *a)
		aux = aux.Siguiente
	}
	if len(lista) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("No hay tiendas en este indice")
	}
	for _, i := range lista {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(i.Nombre)
		json.NewEncoder(w).Encode(i.Descripcion)
		json.NewEncoder(w).Encode(i.Contacto)
		json.NewEncoder(w).Encode(i.Calificacion)
	}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(Vector[obj])
}

func GuardarJson(w http.ResponseWriter, r *http.Request) {
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
		//se escribe el archivo dot
		b := []byte(cadenaJson)
		err := ioutil.WriteFile("guardado.json", b, 0644)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Se guardó json")
		if err != nil {
			log.Fatal(err)
		}

	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "201909103")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", Cargar).Methods("POST")
	router.HandleFunc("/getArreglo", Generardot).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", BuscarPosicion).Methods("POST")
	router.HandleFunc("/id:/{id}", Buscarenvector).Methods("GET")
	router.HandleFunc("/Eliminar", Eliminar).Methods("POST")
	router.HandleFunc("/Guardar", GuardarJson).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
