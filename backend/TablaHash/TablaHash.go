package TablaHash

//William Alejandro Borrayo Alarcón
//201909103

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

//--------------------------------------------------------------------------------------------------------------------------------- Operaciones
func EsPrimo(num int) bool {
	Divisores := 0
	for i := 1; i <= int(num/2); i++ {
		if num%i == 0 {
			Divisores++
		}
		if Divisores > 1 {
			return false
		}
	}
	return true
}

func MetMultiplicacion(m, k int) int {
	A := 3 / 7
	//multiplicacion
	return int(m * ((k * A) % 1))
}

func ExpCuadratica(m, k int, i float64) int {
	//Exploracion Cuadrática
	return ((k % m) + int(math.Pow(i, 2))) % m
}

//----------------------------------------------------------------------------------------------------------------------------------- Tabla
type Tabla struct {
	Id          string
	Tamanio     int
	Valores     []Casilla
	FactorCarga int
}

type Casilla struct {
	Posicion       int
	Usuario        int //DPI
	NombreUsuario  string
	Comentario     string
	SubComentarios *Tabla
}

func (Tabla Tabla) Ocupada(pos int) bool {
	ocupado := false
	for i := 0; i < len(Tabla.Valores); i++ {
		if Tabla.Valores[i].Posicion == pos {
			ocupado = true
			break
		}
	}
	return ocupado
}

func (Tabla Tabla) Ordenar() Tabla {
	var Valores []Casilla
	for i := 0; i < Tabla.Tamanio; i++ {
		for j := 0; j < len(Tabla.Valores); j++ {
			if Tabla.Valores[j].Posicion == i {
				Valores = append(Valores, Tabla.Valores[j])
			}
		}
	}
	Tabla.Valores = Valores
	return Tabla
}

func (Tabla1 *Tabla) InsertarHash(Comentario1 Comentario) {
	factorCarga := (float32(Tabla1.FactorCarga) / float32(Tabla1.Tamanio)) * 100

	if factorCarga < 51 && factorCarga >= 0 { //significa que posee el 50%
		Tabla1.Insertar(Comentario1)
	} else {
		//ReHashing
		fmt.Println("ReHashing...")
		NuevaTabla := new(Tabla)
		Tam := Tabla1.Tamanio + 2*Tabla1.FactorCarga
		for !EsPrimo(Tam) {
			Tam++
		}
		NuevaTabla.Tamanio = Tam
		for i := 0; i < len(Tabla1.Valores); i++ {
			Comentario := &Comentario{"Id", Tabla1.Valores[i].Usuario, Tabla1.Valores[i].NombreUsuario, Tabla1.Valores[i].Comentario}
			NuevaTabla.Insertar(*Comentario)
		}
		Tabla1.Valores = NuevaTabla.Valores
		Tabla1.Tamanio = NuevaTabla.Tamanio
		Tabla1.InsertarHash(Comentario1)
	}
}

func (Tabla1 *Tabla) Insertar(Comentario Comentario) {
	//Metodo de Multiplicacion
	entrada := Comentario.Usuario
	pos := MetMultiplicacion(Tabla1.Tamanio, entrada)
	if !Tabla1.Ocupada(pos) {
		Casilla := new(Casilla)
		Casilla.Posicion = pos
		Casilla.Usuario = entrada
		Casilla.NombreUsuario = Comentario.Nombre
		Casilla.Comentario = Comentario.Mensaje
		TablaNueva := NuevaTabla(Tabla1.Id+"--->"+Comentario.Id, Comentario)
		Casilla.SubComentarios = &TablaNueva
		Tabla1.Valores = append(Tabla1.Valores, *Casilla)
		ActualizarTablas(*Tabla1)
		fmt.Println(Tabla1.Valores)
		Tabla1.FactorCarga++
	} else {
		intento := 1.0
		for Tabla1.Ocupada(pos) {
			//Exploracion Cuadrática
			pos = ExpCuadratica(Tabla1.Tamanio, entrada, intento)
			intento++
		}
		Casilla := new(Casilla)
		Casilla.Posicion = pos
		Casilla.Usuario = entrada
		Casilla.NombreUsuario = Comentario.Nombre
		Casilla.Comentario = Comentario.Mensaje
		TablaNueva := NuevaTabla(Tabla1.Id+"--->"+Comentario.Id, Comentario)
		Casilla.SubComentarios = &TablaNueva
		Tabla1.Valores = append(Tabla1.Valores, *Casilla)
		ActualizarTablas(*Tabla1)
		Tabla1.FactorCarga++
	}
}

func NuevaTabla(id string, Comentario1 Comentario) Tabla {
	Tabla := new(Tabla)
	Tabla.Id = id
	Tabla.Tamanio = 7
	Tabla.FactorCarga = 0
	Tablas = append(Tablas, *Tabla)
	return *Tabla
}

var Tablas []Tabla
var Comentarios []ComentarioJSON

type EnviarJSON struct {
	Comentarios []ComentarioJSON
}

var EnvioJSON EnviarJSON

func NuevoComentario(Comentario Comentario, Padre string) {
	fmt.Println(Padre)
	Tabla1 := *GetTabla(Padre)
	if Tabla1.Id == "" {
		fmt.Println("No existe")
		return
	} else {
		Tabla1.InsertarHash(Comentario)
		ActualizarTablas(Tabla1)
		ComentarioJSON1 := Comentario.AJSON()
		ComentarioJSON1.Ruta = Padre
		Comentarios = append(Comentarios, ComentarioJSON1)
	}
}

func (Comentario *Comentario) AJSON() ComentarioJSON {
	nuevo := new(ComentarioJSON)
	nuevo.Id = Comentario.Id
	nuevo.Mensaje = Comentario.Mensaje
	nuevo.Usuario = Comentario.Usuario
	nuevo.Nombre = Comentario.Nombre
	return *nuevo
}

/*
func GetComentario(Padre string) ComentarioJSON {
	Comment := new(ComentarioJSON)
	for i := 0; i < len(Comentarios); i++ {
		if Padre == Comentarios[i].Ruta {
			Comment = &Comentarios[i]
			break
		}
	}
	return *Comment
}
*/

func GetTabla(id string) *Tabla {
	TablaBuscada := new(Tabla)
	for i := 0; i < len(Tablas); i++ {
		fmt.Println(id + " Comparando con " + Tablas[i].Id)

		if id == Tablas[i].Id {
			TablaBuscada = &Tablas[i]
			break
		}
	}
	return TablaBuscada
}

type ComentarioJSON struct {
	Id      string
	Ruta    string
	Usuario int //DPI
	Nombre  string
	Mensaje string
}

type Comentario struct {
	Id      string
	Usuario int //DPI
	Nombre  string
	Mensaje string
}

func ActualizarTablas(Tabla1 Tabla) {
	var Nuevo []Tabla
	for i := 0; i < len(Tablas); i++ {
		if Tabla1.Id == Tablas[i].Id {
			Nuevo = append(Nuevo, Tabla1)
		} else {
			Nuevo = append(Nuevo, Tablas[i])
		}
	}
	Tablas = Nuevo
}

func SendComentario(w http.ResponseWriter, r *http.Request) {
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(lector)
	c := ComentarioJSON{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	Comentarios = append(Comentarios, c)
	EnvioJSON.Comentarios = Comentarios
}

func GetComentarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EnvioJSON)
}

/*
func main() {

	Tabla0 := new(Tabla)
	Tabla0.Tamanio = 5
	Tabla0.Id = "Tienda"
	Tablas = append(Tablas, *Tabla0)
	//Comentarioni := &ComentarioJSON{Id: "Primero", Ruta: "Tabla", Usuario: 12, Nombre: "yo1", Mensaje: "Prueba1"}
	//Comentarios = append(Comentarios, *Comentarioni)
	Comentario0 := &Comentario{Id: "Primero", Usuario: 12, Nombre: "yo1", Mensaje: "Prueba1"}
	Comentario1 := &Comentario{Id: "Segundo", Usuario: 13, Nombre: "yo2", Mensaje: "Prueba2"}
	Comentario2 := &Comentario{Id: "Tercero", Usuario: 14, Nombre: "yo3", Mensaje: "Prueba3"}
	Comentario3 := &Comentario{Id: "Cuarto", Usuario: 15, Nombre: "yo4", Mensaje: "Prueba4"}
	Comentario4 := &Comentario{Id: "Extra", Usuario: 15, Nombre: "yo5", Mensaje: "Extra"}

	Tabla0.InsertarHash(*Comentario0)

	NuevoComentario(*Comentario0, "Tienda")
	fmt.Println("uno")
	fmt.Println(Comentarios)
	NuevoComentario(*Comentario1, "Tienda--->Primero")
	fmt.Println("dos")
	fmt.Println(Comentarios)
	NuevoComentario(*Comentario2, "Tienda--->Primero--->Segundo")
	fmt.Println("tres")
	fmt.Println(Comentarios)
	NuevoComentario(*Comentario3, "Tienda--->Primero--->Tercero") //no existe
	fmt.Println("cuatro")
	fmt.Println(Comentarios)
	NuevoComentario(*Comentario4, "Tienda--->Primero")
	Prueba := GetTabla("Tienda")
	fmt.Println("\nTabla Final")
	for i := 0; i < len(Prueba.Valores); i++ {
		fmt.Println(Prueba.Valores[i])
	}
	videojuegosComoJson, err := json.Marshal(Comentarios)
	if err != nil {
		fmt.Printf("Error codificando videojuegos: %v", err)
	} else {
		correcto := strings.ReplaceAll(string(videojuegosComoJson), "\\u003e", ">")
		fmt.Println(correcto)
	}

}
*/
