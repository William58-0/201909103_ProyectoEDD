package main

import (
	"os"
	"io"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"os/exec"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//STRUCT para manejar los cursos aprobados
type NodoLista struct{
	Curso string `json:"curso"`
	Nota int `json:"nota"`
	Anio int `json:"anio"`
}


//STRUCT para registro de estudiantes
type Nodo struct{
	Carnet int `json:"carnet"`
	Nombres string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	CUI string `json:"cui"`
	Correo string `json:"correo"`
	Hizq* Nodo `json:"hizq"`
	Hder* Nodo `json:"hder"`
	ListaCursos[] NodoLista `json:"listaCursos"`
}

type Respuesta struct{
	Message string `json:"message"`
}

type listaCarnet struct{
	ListaCarnets[] int `json:"listacarnets"`
}

type asignacion struct{
	Carnet int `json:"carnet"`
	Course NodoLista `json:"curso"`
}

var raiz* Nodo
var contador int

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

//GET obtiene elarbol en json
func getArbol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(raiz)
}

//GET obtiene la lista de carnets en orden
func getListaCarnetsInorden(w http.ResponseWriter, r *http.Request) {
	var lista_carnets[] int
	lista_carnets=listaInorden(raiz,lista_carnets)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&listaCarnet{ListaCarnets:lista_carnets})
}

//POST registra un nuevo estudiante
func createNode(w http.ResponseWriter, r *http.Request) {
	var newNode* Nodo
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &newNode)
	//fmt.Printf("%d",newNode.Carnet)
	//insertamos la raiz
	raiz=crearNodo(raiz,newNode)
	escribir,err2:=json.Marshal(raiz)
	if err2 != nil {
        log.Fatal(err2)
    }
	data := []byte(escribir)
    err = ioutil.WriteFile("persiste.json", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("----------------")
	//preorden(raiz)
	//createDot(raiz)
	//Si todo ha salido bien, devolvemos un status code 201 y el arbol
	w.Header().Set("Content-Type", "application/json")
	respuesta:= &Respuesta{Message:"Alumno creado exitosamente"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}

func crearNodo(nodo* Nodo, nodoInsertar* Nodo) *Nodo{
	//insertamos el primer nodo
	if raiz==nil{
		nuevoNodo:=&Nodo{
			Carnet:nodoInsertar.Carnet, 
			Nombres: nodoInsertar.Nombres,
			Apellidos: nodoInsertar.Apellidos,
			Correo: nodoInsertar.Correo,
			CUI: nodoInsertar.CUI,
			Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
		raiz=nuevoNodo
	} else{
		//sino insertamos en donde corresponda
		raiz=insertarNodo(raiz,nodoInsertar)
	}
	return raiz
}

// función que recorre el arbol para poder insertar el nodo
func insertarNodo(nodo* Nodo, nodoInsertar *Nodo) *Nodo{
	//Si es menor que el carnet actual
	if nodoInsertar.Carnet < nodo.Carnet{
		//si es null el hizq insertamos 
		if nodo.Hizq==nil{
			nuevoNodo:=&Nodo{
				Carnet:nodoInsertar.Carnet, 
				Nombres: nodoInsertar.Nombres,
				Apellidos: nodoInsertar.Apellidos,
				Correo: nodoInsertar.Correo,
				CUI: nodoInsertar.CUI,
				Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
			nodo.Hizq=nuevoNodo
		} else{
			//sino seguimos recorriendo hasta insertar
			nodo.Hizq=insertarNodo(nodo.Hizq, nodoInsertar)
		}
	} else {
		//Si es mayor que el carnet entra a este conjuntp de instrucciones
		//Si el hder es null, insertamos
		if nodo.Hder==nil{
			nuevoNodo:=&Nodo{
				Carnet:nodoInsertar.Carnet, 
				Nombres: nodoInsertar.Nombres,
				Apellidos: nodoInsertar.Apellidos,
				Correo: nodoInsertar.Correo,
				CUI: nodoInsertar.CUI,
				Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
			
			nodo.Hder=nuevoNodo
		} else{
			//de lo contrario seguimos recorriendo hasta insertar
			nodo.Hder=insertarNodo(nodo.Hder, nodoInsertar)
		}
	}
	return nodo 
}

//POST registra un nuevo curso aprobado
func insertarCurso(w http.ResponseWriter, r *http.Request){
	//var cursos[] NodoLista
	var asign* asignacion
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &asign)
	buscado := listaCursosEstudiante(raiz, asign.Carnet)
	buscado.ListaCursos = append(buscado.ListaCursos, asign.Course)
	escribir,err2:=json.Marshal(raiz)
	if err2 != nil {
        log.Fatal(err2)
    }
	data := []byte(escribir)
    err = ioutil.WriteFile("persiste.json", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	estudiante:=&Nodo{
		Carnet:buscado.Carnet, 
		Nombres: buscado.Nombres,
		Apellidos: buscado.Apellidos,
		Correo: buscado.Correo,
		CUI: buscado.CUI,
		Hizq:nil, Hder:nil, ListaCursos:buscado.ListaCursos,
	}
	w.Header().Set("Content-Type", "application/json")
	//respuesta:= &Respuesta{Message:"Curso Creado"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(estudiante)
}

//POST obtiene los cursos de un estudiante
func getListaCursos(w http.ResponseWriter, r *http.Request){
	//var cursos[] NodoLista
	var newNode* Nodo
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &newNode)
	buscado := listaCursosEstudiante(raiz, newNode.Carnet)
	estudiante:=&Nodo{
		Carnet:buscado.Carnet, 
		Nombres: buscado.Nombres,
		Apellidos: buscado.Apellidos,
		Correo: buscado.Correo,
		CUI: buscado.CUI,
		Hizq:nil, Hder:nil, ListaCursos:buscado.ListaCursos,
	}
	//fmt.Println(estudiante)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(estudiante)
}

func listaCursosEstudiante(nodo* Nodo, carnet int) *Nodo{
	var es* Nodo
	if nodo!=nil{
		if nodo.Carnet==carnet{
			es=nodo
			fmt.Println("ENCONTRADOOOOOOOOO")
			return es
		}else{
			if nodo.Carnet>carnet{
				es=listaCursosEstudiante(nodo.Hizq,carnet)
			}else{
				es=listaCursosEstudiante(nodo.Hder,carnet)
			}
		}
	}
	return es
}

//Recorrido en inorden
func listaInorden(nodo* Nodo, lista[] int) []int {
	if nodo.Hizq != nil{
		lista=listaInorden(nodo.Hizq, lista)
	}
	lista = append(lista, nodo.Carnet)
	if nodo.Hder != nil{
		lista=listaInorden(nodo.Hder, lista)
	}
	return lista
}

//Recorrido en preorden
func preorden(nodo* Nodo) {
	fmt.Println(nodo.Carnet)
	if nodo.Hizq != nil{
		preorden(nodo.Hizq)
	}
	if nodo.Hder != nil{
		preorden(nodo.Hder)
	}
}

//GET genera el arbol a partir de la variable raíz y envía la imagen al frontend
func getImagenArbol(w http.ResponseWriter, r *http.Request){
	//Generamos el archivo dot
	data := []byte(createDot(raiz))
    err := ioutil.WriteFile("grafo.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearGrafo.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./grafo.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)

}

func recorrerArbol(nombrePadre string, hijo* Nodo, textoActual string) string{
	if hijo.Hizq!=nil{
		nombreHijo := "Nodo"
		nombreHijo+=strconv.FormatInt(int64(contador), 10)
		contador+=1
		textoActual+=nombreHijo
		textoActual+=`[shape=none label=<
		`
		textoActual+=`<table cellspacing="0" border="0" cellborder="1">
		<tr>
			<td colspan="2"><img src="avatar.png" /></td>
		</tr>`
		textoActual+="<tr><td colspan=\"2\">Nombre: "+ hijo.Hizq.Nombres+" "+hijo.Hizq.Apellidos+"</td></tr>"
		textoActual+="<tr><td>Carnet: "+strconv.FormatInt(int64(hijo.Hizq.Carnet), 10)+"</td><td>CUI: "+hijo.Hizq.CUI+"</td></tr>"
		textoActual+="<tr><td colspan=\"2\">Correo: "+hijo.Hizq.Correo+"</td></tr></table>"
		textoActual+=`
		>];
		`
		textoActual+=nombrePadre+"->"+nombreHijo+";\n"
		textoActual=recorrerArbol(nombreHijo,hijo.Hizq, textoActual)
	}
	if hijo.Hder!=nil{
		nombreHijo := "Nodo"
		nombreHijo+=strconv.FormatInt(int64(contador), 10)
		contador+=1
		textoActual+=nombreHijo
		textoActual+=`[shape=none label=<
		`
		textoActual+=`<table cellspacing="0" border="0" cellborder="1">
		<tr>
			<td colspan="2"><img src="avatar.png" /></td>
		</tr>`
		textoActual+="<tr><td colspan=\"2\">Nombre: "+ hijo.Hder.Nombres+" "+hijo.Hder.Apellidos+"</td></tr>"
		textoActual+="<tr><td>Carnet: "+strconv.FormatInt(int64(hijo.Hder.Carnet), 10)+"</td><td>CUI: "+hijo.Hder.CUI+"</td></tr>"
		textoActual+="<tr><td colspan=\"2\">Correo: "+hijo.Hder.Correo+"</td></tr></table>"
		textoActual+=`
		>];
		`
		textoActual+=nombrePadre+"->"+nombreHijo+";\n"
		textoActual=recorrerArbol(nombreHijo,hijo.Hder, textoActual)
	}
	return textoActual
}

func createDot(nodo* Nodo) string{
	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=`Nodo0[shape=none label=<
	`
	grafo+=`<table cellspacing="0" border="0" cellborder="1">
	<tr>
		<td colspan="2"><img src="avatar.png" /></td>
	</tr>`
	grafo+="<tr><td colspan=\"2\">Nombre: "+ nodo.Nombres+" "+nodo.Apellidos+"</td></tr>"
	grafo+="<tr><td>Carnet: "+strconv.FormatInt(int64(nodo.Carnet), 10)+"</td><td>CUI: "+nodo.CUI+"</td></tr>"
	grafo+="<tr><td colspan=\"2\">Correo: "+nodo.Correo+"</td></tr></table>"
	grafo+=`
	>];
	`
	contador=1
	grafo=recorrerArbol("Nodo0", nodo, grafo)
	grafo+="}"
	return grafo
	//fmt.Println(grafo)
}

//Carga los datos del arbol que se encuentran almacenados en un archivo json
func cargarDatos() *Nodo{
	var arbol *Nodo
	datosArchivo, err := ioutil.ReadFile("./persiste.json")
    if err != nil {
        log.Fatal(err)
    }
	err=json.Unmarshal(datosArchivo, &arbol)
	if err != nil {
        log.Fatal(err)
    }
	return arbol
}

func main() {
	raiz=cargarDatos()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/arbol", getArbol).Methods("GET")
	router.HandleFunc("/crearEstudiante", createNode).Methods("POST")
	router.HandleFunc("/arbolBinario", getImagenArbol).Methods("GET")
	router.HandleFunc("/listaEstudiantes", getListaCarnetsInorden).Methods("GET")
	router.HandleFunc("/cursosEstudiante", getListaCursos).Methods("POST")
	router.HandleFunc("/insertarCurso", insertarCurso).Methods("POST")
	//Se agregaron cors
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}