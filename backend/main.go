package main

import (
	"container/list"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"./AVL"
	"./ArbolB"
	"./Carrito"
	"./Estructuras"
	"./Grafos"
	"./MatrizDispersa"
	"./TablaHash"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//-------------------------------------------------------------------------------------------------------------GRAFO DE VECTOR
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

func Graficar() {
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
				cadena = cadena + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) +
					"[label=\"Nombre: " + aux.Nombre + " \\n Contacto: " + aux.Contacto + " \\n Calificacion: " +
					tipoCalif(aux.Calificacion) + "\", fillcolor=\"yellowgreen:aquamarine\"];\n" +
					"Vector:" + strconv.Itoa(j) + "->" + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) +
					"[color=red]\n"
				contador++
				aux = aux.Siguiente
				for aux != nil {
					cadena = cadena + strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "->" +
						strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador-1) + "\n" +
						strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador-1) + "->" +
						strconv.Itoa(j) + strconv.Itoa(aux.Calificacion) + strconv.Itoa(contador) + "\n" +
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
		b := []byte(cadena)
		err := ioutil.WriteFile("../frontend/src/assets/img/Vector"+strconv.Itoa(a)+".dot", b, 0644)
		if err != nil {
			log.Fatal(err)
		}
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "../frontend/src/assets/img/Vector"+strconv.Itoa(a)+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile("../frontend/src/assets/img/Vector"+strconv.Itoa(a)+".png", cmd, os.FileMode(mode))
		fmt.Println("Vector " + strconv.Itoa(a))
		a++
	}
	fmt.Println("Matriz Linealizada")
}

func Generardot(w http.ResponseWriter, r *http.Request) {
	if len(Vector) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("No hay datos cargados")
		return
	} else {
		Graficar()
	}
}

//--------------------------------------------------------------------------------------------------------------------      FUNCIONES
var depa []string
var ind []string
var Vector []Estructuras.Lista

type ListaTiendas struct {
	ListaTiendas []Estructuras.Tienda1 `json:"ListaTiendas"`
}

func GetTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todo)
}

func Corregir(palabra string) string {
	PrimeraLetra := strings.ToUpper(string(palabra[0]))
	Siguiente := strings.TrimLeft(palabra, string(palabra[0]))
	Nueva := PrimeraLetra + Siguiente
	return Nueva
}

var Todo Estructuras.Todo

//-------------------------------------------------------------------------------------------------------------------------Matriz Linealizada
func Cargar(w http.ResponseWriter, r *http.Request) {
	arbolMerkle := NewArbol()
	Vector = nil
	AVL.Todo1.Productos = nil
	MatrizDispersa.Todo1.Fechas = nil
	MatrizDispersa.Pedd1.Productos = nil
	var ListaTiendas []Estructuras.Tienda1
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := Estructuras.Data{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(c.Datos); i++ {
		for j := 0; j < len(c.Datos[i].Departamentos); j++ {
			existe := false
			for _, a := range depa {
				if Corregir(c.Datos[i].Departamentos[j].Nombre) == a {
					existe = true
				}
			}
			if !existe {
				depa = append(depa, Corregir(c.Datos[i].Departamentos[j].Nombre))
			}
		}
	}
	for i := 0; i < len(c.Datos); i++ {
		existe := false
		for _, a := range ind {
			if Corregir(c.Datos[i].Indice) == a {
				existe = true
			}
		}
		if !existe {
			ind = append(ind, Corregir(c.Datos[i].Indice))
		}
	}
	for let := 0; let < len(ind); let++ {
		for dep := 0; dep < len(depa); dep++ {
			for calif := 1; calif <= 5; calif++ {
				Lista := new(Estructuras.Lista)
				Lista.Indice = ind[let]
				Lista.Categoria = depa[dep]
				Lista.Calificacion = calif
				for i := 0; i < len(c.Datos); i++ {
					for j := 0; j < len(c.Datos[i].Departamentos); j++ {
						for k := 0; k < len(c.Datos[i].Departamentos[j].Tiendas); k++ {
							prim := c.Datos[i]
							sec := prim.Departamentos[j]
							terc := sec.Tiendas[k]
							terc.Departamento = depa[dep]
							if Corregir(prim.Indice) == ind[let] && Corregir(sec.Nombre) == depa[dep] && terc.Calificacion == calif {
								Lista.Insertar(terc.Nombre, terc.Descripcion, terc.Contacto, terc.Calificacion, terc.Departamento, terc.Logo)
							}
						}
					}
				}
				Vector = append(Vector, *Lista)
				if Lista.Tamanio == 0 {
					Lista.Primero = nil
					Lista.Ultimo = nil
				}
				Nombres := Lista.Ordenar()
				ListaOrdenada := new(Estructuras.Lista)
				for _, o := range Nombres {
					aux1 := Lista.Primero
					for aux1 != nil {
						if aux1.Nombre == o {
							Tienda1 := new(Estructuras.Tienda1)
							Tienda1.Nombre = aux1.Nombre
							Tienda1.Descripcion = aux1.Descripcion
							Tienda1.Contacto = aux1.Contacto
							Tienda1.Calificacion = aux1.Calificacion
							Tienda1.Departamento = aux1.Departamento
							Tienda1.Logo = aux1.Logo
							if Tienda1 != nil {
								//-----------------------------------------------------------------Se insertan al arbol merkle
								//---------------------------------------------- HASH: NOMBRE + DEPARTAMENTO + CALIFICACION
								hash := Hash(aux1.Nombre + aux1.Departamento + strconv.Itoa(aux1.Calificacion))
								//---------------------------------------------- Insercion
								arbolMerkle.Insertar(Hash(hash), aux1.Nombre, aux1.Departamento, aux1.Calificacion)
								//----------------------------------------------------------------Se agregan a la lista
								ListaTiendas = append(ListaTiendas, *Tienda1)
								ListaOrdenada.Insertar(aux1.Nombre, aux1.Descripcion, aux1.Contacto, aux1.Calificacion, aux1.Departamento, aux1.Logo)
							}
						}
						aux1 = aux1.Siguiente
					}
				}
				ListaOrdenada.Indice = Lista.Indice
				ListaOrdenada.Calificacion = Lista.Calificacion
				ListaOrdenada.Categoria = Lista.Categoria
				indice := (let*len(depa)+dep)*5 + (calif - 1)
				Vector[indice] = *ListaOrdenada
			}
		}
		fmt.Println("Tiendas Cargadas")
		Todo.Tiendas = ListaTiendas
		AVL.Tiendas = ListaTiendas
	}
	w.Header().Set("Content-Type", "application/json")
	//----------------------------------------------------------Matriz Linealizada
	Graficar()
	//-----------------------------------------------------------Arbol Merkle
	arbolMerkle.GraficarMerkle()
	json.NewEncoder(w).Encode("Datos Cargados")
}

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
			if Lista.Categoria == Departamento && Lista.Calificacion == Calificacion {
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
			if Lista.Categoria == Categoria && Lista.Calificacion == Calificacion {
				if Lista.Eliminar(Nombre, Calificacion) {
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
		if conviene {
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
		b := []byte(cadenaJson)
		err := ioutil.WriteFile("guardado.json", b, 0644)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Se guardó json")
		if err != nil {
			log.Fatal(err)
		}
	}
}

//--------------------------------------------------------------------------------------ARBOL TIENDAS
type Nodo struct {
	Hash         string
	Nombre       string
	Departamento string
	Calificacion int
	Derecha      *Nodo
	Izquierda    *Nodo
}

type Arbol struct {
	Raiz *Nodo
}

func newNodo(Hash string, Nombre, Departamento string, Calificacion int, Derecha *Nodo, Izquierda *Nodo) *Nodo {
	return &Nodo{Hash, Nombre, Departamento, Calificacion, Derecha, Izquierda}
}

func NewArbol() *Arbol {
	return &Arbol{}
}

func (this *Arbol) Insertar(Hash1, Nombre, Departamento string, Calificacion int) {
	n := newNodo(Hash1, Nombre, Departamento, Calificacion, nil, nil)
	if this.Raiz == nil {
		lista := list.New()
		lista.PushBack(n)
		lista.PushBack(newNodo(Hash(""), "", "", -1, nil, nil))
		this.construirArbol(lista)
	} else {
		lista := this.obtenerLista()
		lista.PushBack(n)
		this.construirArbol(lista)
	}
}

func (this *Arbol) obtenerLista() *list.List {
	lista := list.New()
	obtenerLista(lista, this.Raiz.Izquierda)
	obtenerLista(lista, this.Raiz.Derecha)
	return lista
}

func obtenerLista(lista *list.List, actual *Nodo) {
	if actual != nil {
		obtenerLista(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Hash != Hash("") {
			lista.PushBack(actual)
		}
		obtenerLista(lista, actual.Derecha)
	}
}

func (this *Arbol) construirArbol(lista *list.List) {
	size := float64(lista.Len())
	cant := 1
	for (size / 2) > 1 {
		cant++
		size = size / 2
	}
	nodostot := math.Pow(2, float64(cant))
	for lista.Len() < int(nodostot) {
		lista.PushBack(newNodo(Hash(""), "", "", -1, nil, nil))
	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*Nodo)
		nodo2 := segundo.Value.(*Nodo)
		h := ""
		if nodo2.Hash != "" {
			h = nodo1.Hash + "\\n" + nodo2.Hash
		} else {
			h = nodo1.Hash
		}
		a := Hash(h)
		nuevo := newNodo(a, h, "", -1, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*Nodo)
}

func (this *Arbol) GraficarMerkle() {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\", style=\"filled\"];\n")
	if this.Raiz != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"{%s | %s}\", fillcolor=\"green\"];\n", &(*this.Raiz), this.Raiz.Hash, this.Raiz.Nombre)
		this.generar(&cadena, (this.Raiz), this.Raiz.Izquierda, (this.Raiz))
		this.generar(&cadena, (this.Raiz), this.Raiz.Derecha, (this.Raiz))
	}
	fmt.Fprintf(&cadena, "}\n")
	//hacer el dot y la imagen
	b := []byte(cadena.String())
	err := ioutil.WriteFile("../frontend/src/assets/img/MerkleTiendas.dot", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "../frontend/src/assets/img/MerkleTiendas.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("../frontend/src/assets/img/MerkleTiendas.png", cmd, os.FileMode(mode))
	fmt.Println("MerkleTiendas")
}

func (this *Arbol) generar(cadena *strings.Builder, padre *Nodo, actual *Nodo, Raiz *Nodo) {
	if actual != nil {
		if actual.Hash != Hash("") {
			if actual.Calificacion >= 0 {
				fmt.Fprintf(cadena, "node%p[label=\"{%s |Nombre: %s \\nDep: %s \\nCalif: %v}\", fillcolor=\"green\"];\n",
					&(*actual), actual.Hash, actual.Nombre, actual.Departamento, actual.Calificacion)
			} else {
				fmt.Fprintf(cadena, "node%p[label=\"{%s | %s}\", fillcolor=\"green\"];\n", &(*actual), actual.Hash, actual.Nombre)
			}
		} else {
			fmt.Fprintf(cadena, "node%p[label=\"{%s |%s \\n %s \\n %v}\", fillcolor=\"gray\" ,color=\"red\"];\n",
				&(*actual), actual.Hash, actual.Nombre, actual.Departamento, actual.Calificacion)
		}
		fmt.Fprintf(cadena, "node%p->node%p [dir=back]\n", &(*padre), &(*actual))
		this.generar(cadena, actual, actual.Izquierda, Raiz)
		this.generar(cadena, actual, actual.Derecha, Raiz)
	}
}

func Hash(texto string) string {
	hash := sha256.Sum256([]byte(texto))
	return fmt.Sprintf("%x", hash)
}

//--------------------------------------------------------------------------------------ENDPOINTS
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "William Alejandro Borrayo Alarcon_201909103")
}

func main() {
	//Load() //para cargar las tiendas
	//MatrizDispersa.Leer()
	router := mux.NewRouter().StrictSlash(true)
	arbol := ArbolB.NewArbol(5)
	ArbolB.Administrador(arbol)
	arbol.Graficar("Sin")
	arbol.Graficar("Cif")
	arbol.Graficar("CifSen")
	//------------------------------------------------------------------FASE 1
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", Cargar).Methods("POST")
	router.HandleFunc("/getArreglo", Generardot).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", BuscarPosicion).Methods("POST")
	router.HandleFunc("/id/{id}", Buscarenvector).Methods("GET")
	router.HandleFunc("/Eliminar", Eliminar).Methods("DELETE")
	router.HandleFunc("/Guardar", GuardarJson).Methods("POST")
	//------------------------------------------------------------------FASE 2
	router.HandleFunc("/LoadTiendas", Cargar).Methods("POST")
	router.HandleFunc("/GetTiendas", GetTiendas).Methods("GET")
	router.HandleFunc("/LoadInventario", AVL.Leer).Methods("POST")
	router.HandleFunc("/GetInventario", AVL.GetInventario).Methods("POST")
	router.HandleFunc("/LoadFechas", MatrizDispersa.Leer).Methods("POST")
	router.HandleFunc("/GetFechas", MatrizDispersa.GetFechas).Methods("GET")
	router.HandleFunc("/GetPedidos", MatrizDispersa.GetPedidos).Methods("GET")
	router.HandleFunc("/Comprar", Carrito.RestarProducto).Methods("POST")
	router.HandleFunc("/Devolver", Carrito.SumarProducto).Methods("POST")
	router.HandleFunc("/CargarCarro", Carrito.GetCarrito).Methods("GET")
	router.HandleFunc("/GenerarPedido", Carrito.GenerarPedido).Methods("POST")
	//------------------------------------------------------------------FASE 3
	router.HandleFunc("/LoadUsuarios", ArbolB.Cargar).Methods("POST")
	router.HandleFunc("/IniciarSesion", ArbolB.IniciarSesion).Methods("POST")
	router.HandleFunc("/GetUsuario", ArbolB.GetUsuario).Methods("POST")
	router.HandleFunc("/Registrar", ArbolB.Registrar).Methods("POST")
	router.HandleFunc("/Eliminar", ArbolB.Eliminar).Methods("POST")
	router.HandleFunc("/GrafoInicial", Grafos.GrafoInicial).Methods("POST")
	router.HandleFunc("/GenerarRecorrido", Grafos.GenerarRecorrido).Methods("POST")
	router.HandleFunc("/GetRecorrido", Grafos.GetRecorrido).Methods("GET")
	router.HandleFunc("/GetUsuarios", ArbolB.GetUsuarios).Methods("GET")
	//-------------------------------------------------------------------FASE 4
	router.HandleFunc("/SendComentario", TablaHash.SendComentario).Methods("POST")
	router.HandleFunc("/GetComentarios", TablaHash.GetComentarios).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
