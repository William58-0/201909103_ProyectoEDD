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
	"strings"

	"./Estructuras"
	"github.com/gorilla/mux"
)

func tipoCalif(calif int) string {
	switch ej := calif; ej {
	case 1:
		return "Regular (1)"
	case 2:
		return "Buena (2)"
	case 3:
		return "Muy Buena (3)"
	case 4:
		return "Excelente (4)"
	default:
		return "Magnifica (5)"
	}
}

func Generardot(w http.ResponseWriter, r *http.Request) {
	if len(Vector) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("No hay datos cargados")
		return
	}
	a := 0
	k := 0
	q := 0
	for a < len(Vector)/5+len(Vector)%5 && k < len(Vector) && q < len(Vector) {
		cadena := "digraph grafo{\nfontname=\"Verdana\" color=red fontsize=22;\n" +
			"node [shape=record fontsize=8 fontname=\"Verdana\" style=filled];\n" +
			"edge [color=\"blue\"]\nsubgraph cluster{\nlabel = \"Vector\";\nbgcolor=\"yellow:dodgerblue\"\n" +
			"Vector[label=\""
		for i := k; i < len(Vector); i++ {
			cadena = cadena + "<" + strconv.Itoa(i) + ">" +
				"Posicion: " + strconv.Itoa(i+1) +
				"\\n Indice: " + Vector[i].Indice +
				"\\n Categoria: " + Vector[i].Categoria +
				"\\n Calificacion: " + strconv.Itoa(Vector[i].Calificacion) + "|"
			if (i+1)%5 == 0 {
				k = i + 1
				break
			}
		}
		cadena = strings.TrimSuffix(cadena, "|")
		cadena = cadena + "\",width=20, fillcolor=\"aquamarine:violet\"];\n}\n"
		for j := q; j < len(Vector); j++ {
			if Vector[j].Primero != nil {
				contador := 1
				Lista := Vector[j]
				aux := Lista.Primero
				//primer nodo
				//id del nodo=posicion en el Vector + calificacion + posicion en la lista
				cadena = cadena + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) +
					"[label=\"Nombre: " + aux.Nombre + " \\n Contacto: " + aux.Contacto + " \\n Calificacion: " +
					tipoCalif(aux.Calificacion) + "\", fillcolor=\"yellowgreen:aquamarine\"];\n" +
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
						strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "[label=\"Nombre: " + aux.Nombre +
						" \\n Contacto: " + aux.Contacto + " \\n Calificacion: " + tipoCalif(aux.Calificacion) + "\", fillcolor=\"yellowgreen:aquamarine\"];\n"
					contador++
					aux = aux.Siguiente
				}
			}
			if (j+1)%5 == 0 {
				q = j + 1
				break
			}
		}
		cadena = cadena + "}"
		//se escribe el archivo dot
		b := []byte(cadena)
		err := ioutil.WriteFile("grafo"+strconv.Itoa(a)+".dot", b, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//se crea la imagen
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "grafo"+strconv.Itoa(a)+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile("grafo"+strconv.Itoa(a)+".png", cmd, os.FileMode(mode))
		a++
	}
}

////////////////////////////////////////////////////////////////////////////////////////////      FUNCIONES

//arreglo de departamentos
var depa []string

//arreglo de indices
var ind []string

//Vector para Linealizar
var Vector []Estructuras.Lista

func Cargar(w http.ResponseWriter, r *http.Request) {
	Vector = nil
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
	//Contar numero de indices
	for i := 0; i < len(c.Datos); i++ {
		existe := false
		for _, a := range ind {
			if c.Datos[i].Indice == a {
				existe = true
			}
		}
		if existe == false {
			ind = append(ind, c.Datos[i].Indice)
		}
	}
	//indices
	for _, let := range ind {
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
				//Se ordena la lista
				Nombres := Lista.Ordenar()
				ListaOrdenada := new(Estructuras.Lista)
				for _, o := range Nombres {
					aux1 := Lista.Primero
					for aux1 != nil {
						if aux1.Nombre == o {
							ListaOrdenada.Insertar(aux1.Nombre, aux1.Descripcion, aux1.Contacto, aux1.Calificacion)
						}
						aux1 = aux1.Siguiente
					}
				}
				ListaOrdenada.Indice = Lista.Indice
				ListaOrdenada.Calificacion = Lista.Calificacion
				ListaOrdenada.Categoria = Lista.Categoria
				//se agrega la lista al Vector
				Vector = append(Vector, *ListaOrdenada)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Datos Cargados")
}

//Busqueda de Posicion especifica
func BuscarPosicion(w http.ResponseWriter, r *http.Request) {
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := Estructuras.Objetivo{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Departamento := c.Departamento
	Nombre := c.Nombre
	Calificacion := c.Calificacion
	for _, Lista := range Vector {
		if Lista.Primero != nil {
			//Si encuentra la categoria, y la calificacion
			if Lista.Categoria == Departamento && Lista.Calificacion == Calificacion {
				//buscar en la lista
				a := Lista.Buscar(Nombre, Calificacion)
				if a.Nombre != "" {
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
	c := Estructuras.ObjetivoE{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Nombre := c.Nombre
	Categoria := c.Categoria
	Calificacion := c.Calificacion
	cont := 0
	for _, Lista := range Vector {
		if Lista.Primero != nil {
			//Si encuentra la categoria, y la calificacion
			if Lista.Categoria == Categoria && Lista.Calificacion == Calificacion {
				//buscar en la lista
				if Lista.Eliminar(Nombre, Calificacion) == true {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode("Eliminado")
					Vector[cont] = Lista
					return
				}
			}
		}
		cont++
	}
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
		a.Contacto = aux.Contacto
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
		json.NewEncoder(w).Encode(i)
	}
}

func GuardarJson(w http.ResponseWriter, r *http.Request) {
	Data := new(Estructuras.Data1)
	var ListaPrincipal []Estructuras.Principal1
	for let := 0; let < len(ind); let++ {
		Principal := new(Estructuras.Principal1)
		Principal.Indice = string(ind[let])
		var vector []Estructuras.Dep1
		for b := 0; b < len(depa); b++ {
			Departamento := new(Estructuras.Dep1)
			Departamento.Nombre = depa[b]
			var ListaTiendas []Estructuras.Tienda1
			for a := 0; a < 5; a++ {
				indice := let*5*len(depa) + (b * 5) + a
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
		conviene := false
		for _, o := range vector {
			if o.Tiendas != nil {
				conviene = true
			}
		}
		if conviene == true {
			Principal.Departamentos = vector
			ListaPrincipal = append(ListaPrincipal, *Principal)
		}
	}
	Data.Datos = ListaPrincipal
	cadenaJson, err := json.Marshal(Data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Data)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		//se escribe el archivo json
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
	fmt.Fprintf(w, "William Alejandro Borrayo Alarcon_201909103")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", Cargar).Methods("POST")
	router.HandleFunc("/getArreglo", Generardot).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", BuscarPosicion).Methods("POST")
	router.HandleFunc("/id/{id}", Buscarenvector).Methods("GET")
	router.HandleFunc("/Eliminar", Eliminar).Methods("DELETE")
	router.HandleFunc("/Guardar", GuardarJson).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
