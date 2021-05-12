package ArbolB

import (
	"container/list"
	"crypto/aes"
	"crypto/cipher"
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
)

//-----------------------------------------------------------------------------------------------------KEY
type Key struct {
	DPI         int
	Correo      string
	Contrasenia string
	Nombre      string
	Tipo        string
	Izquierdo   *Nodo
	Derecho     *Nodo
}

func NewKey(DPI int, Correo string, Contrasenia string, Nombre string, Tipo string) *Key {
	k := Key{DPI, Correo, Contrasenia, Nombre, Tipo, nil, nil}
	return &k
}

//----------------------------------------------------------------------------------------------------NODO
type Nodo struct {
	Max       int
	NodoPadre *Nodo
	Keys      []*Key
}

func NewNodo(max int) *Nodo {
	keys := make([]*Key, max)
	n := Nodo{max, nil, keys}
	return &n
}

func (this *Nodo) Colocar(i int, llave *Key) {
	this.Keys[i] = llave
}

//---------------------------------------------------------------------------------------------------ARBOL
type Arbol struct {
	k    int
	Raiz *Nodo
}

func NewArbol(nivel int) *Arbol {
	a := Arbol{nivel, nil}
	nodoRaiz := NewNodo(nivel)
	a.Raiz = nodoRaiz
	return &a
}

func (this *Arbol) Insertar(newKey *Key) {
	if this.Raiz.Keys[0] == nil {
		this.Raiz.Colocar(0, newKey)
	} else if this.Raiz.Keys[0].Izquierdo == nil {
		lugarinsertado := -1
		node := this.Raiz
		lugarinsertado = this.colocarNodo(node, newKey)
		if lugarinsertado != -1 {
			if lugarinsertado == node.Max-1 {
				mid := node.Max / 2
				llavecentral := node.Keys[mid]
				derecho := NewNodo(this.k)
				izquierdo := NewNodo(this.k)
				indiceizq := 0
				indicedere := 0
				for j := 0; j < node.Max; j++ {
					if node.Keys[j].DPI < llavecentral.DPI {
						izquierdo.Colocar(indiceizq, node.Keys[j])
						indiceizq++
						node.Colocar(j, nil)
					} else if node.Keys[j].DPI > llavecentral.DPI {
						derecho.Colocar(indicedere, node.Keys[j])
						indicedere++
						node.Colocar(j, nil)
					}
				}
				node.Colocar(mid, nil)
				this.Raiz = node
				this.Raiz.Colocar(0, llavecentral)
				izquierdo.NodoPadre = this.Raiz
				derecho.NodoPadre = this.Raiz
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
			}
		}
	} else if this.Raiz.Keys[0].Izquierdo != nil {
		node := this.Raiz
		for node.Keys[0].Izquierdo != nil {
			loop := 0
			for i := 0; i < node.Max; i, loop = i+1, loop+1 {
				if node.Keys[i] != nil {
					if node.Keys[i].DPI > newKey.DPI {
						node = node.Keys[i].Izquierdo
						break
					}
				} else {
					node = node.Keys[i-1].Derecho
					break
				}
			}
			if loop == node.Max {
				node = node.Keys[loop-1].Derecho
			}
		}
		indiceColocado := this.colocarNodo(node, newKey)
		if indiceColocado == node.Max-1 {
			for node.NodoPadre != nil {
				indicemedio := node.Max / 2
				llavecentral := node.Keys[indicemedio]
				izquierdo := NewNodo(this.k)
				derecho := NewNodo(this.k)
				indiceizquierdo, indicederecho := 0, 0
				for i := 0; i < node.Max; i++ {
					if node.Keys[i].DPI < llavecentral.DPI {
						izquierdo.Colocar(indiceizquierdo, node.Keys[i])
						indiceizquierdo++
						node.Colocar(i, nil)
					} else if node.Keys[i].DPI > llavecentral.DPI {
						derecho.Colocar(indicederecho, node.Keys[i])
						indicederecho++
						node.Colocar(i, nil)
					}
				}
				node.Colocar(indicemedio, nil)
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
				node = node.NodoPadre
				izquierdo.NodoPadre = node
				derecho.NodoPadre = node
				for i := 0; i < izquierdo.Max; i++ {
					if izquierdo.Keys[i] != nil {
						if izquierdo.Keys[i].Izquierdo != nil {
							izquierdo.Keys[i].Izquierdo.NodoPadre = izquierdo
						}
						if izquierdo.Keys[i].Derecho != nil {
							izquierdo.Keys[i].Derecho.NodoPadre = izquierdo
						}
					}
				}
				for i := 0; i < derecho.Max; i++ {
					if derecho.Keys[i] != nil {
						if derecho.Keys[i].Izquierdo != nil {
							derecho.Keys[i].Izquierdo.NodoPadre = derecho
						}
						if derecho.Keys[i].Derecho != nil {
							derecho.Keys[i].Derecho.NodoPadre = derecho
						}
					}
				}
				lugarcolocado := this.colocarNodo(node, llavecentral)
				if lugarcolocado == node.Max-1 {
					if node.NodoPadre == nil {
						indicecentralraiz := node.Max / 2
						llavecentralraiz := node.Keys[indicecentralraiz]
						izquierdoraiz := NewNodo(this.k)
						derechoraiz := NewNodo(this.k)
						indicederechoraiz, indiceizquierdoraiz := 0, 0
						for i := 0; i < node.Max; i++ {
							if node.Keys[i].DPI < llavecentralraiz.DPI {
								izquierdoraiz.Colocar(indiceizquierdoraiz, node.Keys[i])
								indiceizquierdoraiz++
								node.Colocar(i, nil)

							} else if node.Keys[i].DPI > llavecentralraiz.DPI {
								derechoraiz.Colocar(indicederechoraiz, node.Keys[i])
								indicederechoraiz++
								node.Colocar(i, nil)
							}
						}
						node.Colocar(indicecentralraiz, nil)
						node.Colocar(0, llavecentralraiz)
						for i := 0; i < this.k; i++ {
							if izquierdoraiz.Keys[i] != nil {
								izquierdoraiz.Keys[i].Izquierdo.NodoPadre = izquierdoraiz
								izquierdoraiz.Keys[i].Derecho.NodoPadre = izquierdoraiz
							}

						}
						for i := 0; i < this.k; i++ {
							if derechoraiz.Keys[i] != nil {
								derechoraiz.Keys[i].Izquierdo.NodoPadre = derechoraiz
								derechoraiz.Keys[i].Derecho.NodoPadre = derechoraiz
							}
						}
						llavecentralraiz.Izquierdo = izquierdoraiz
						llavecentralraiz.Derecho = derechoraiz
						izquierdoraiz.NodoPadre = node
						derechoraiz.NodoPadre = node
						this.Raiz = node
					}
					continue
				} else {
					break
				}
			}
		}
	}
}

func (this *Arbol) colocarNodo(nodo *Nodo, newKey *Key) int {
	index := -1
	for i := 0; i < nodo.Max; i++ {
		if nodo.Keys[i] == nil {
			placed := false
			for j := i - 1; j >= 0; j-- {
				if nodo.Keys[j].DPI > newKey.DPI {
					nodo.Colocar(j+1, nodo.Keys[j])
				} else {
					nodo.Colocar(j+1, newKey)
					nodo.Keys[j].Derecho = newKey.Izquierdo
					if (j+2) < this.k && nodo.Keys[j+2] != nil {
						nodo.Keys[j+2].Izquierdo = newKey.Derecho
					}
					placed = true
					break
				}
			}
			if !placed {
				nodo.Colocar(0, newKey)
				nodo.Keys[1].Izquierdo = newKey.Derecho
			}
			index = i
			break
		}
	}
	return index
}

func Buscar(DPI int) bool {
	for i := 0; i < len(Users); i++ {
		if Users[i].Dpi == DPI {
			return true
		}
	}
	return false
}

//--------------------------------------------------------------------------------------------Cargar Datos
type Usuarios struct {
	Usuarios []*Usuario `json:"Usuarios"`
}

type Usuario struct {
	Dpi      int    `json:"Dpi"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"`
}

func Administrador(this *Arbol) {
	Admin := new(Usuario)
	Admin.Dpi = 1234567890101
	Admin.Correo = "auxiliar@edd.com"
	Admin.Password = "1234"
	Admin.Nombre = "EDD2021"
	Admin.Cuenta = "Admin"
	this.Insertar(NewKey(1234567890101, "auxiliar@edd.com", "1234", "EDD2021", "Admin"))
	Users = append(Users, *Admin)
}

type Usuario1 struct {
	Dpi      string `json:"Dpi"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"`
}

var SesionActual int

func iniciarsesion(DPI string, password string) Usuario {
	SesionActual = 0
	User := new(Usuario)
	for i := 0; i < len(Users); i++ {
		if strconv.Itoa(Users[i].Dpi) == DPI && Users[i].Password == password {
			SesionActual = Users[i].Dpi
			*User = Users[i]
			return Users[i]
		}
	}
	return *User
}

func IniciarSesion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IniciarSesion")
	var usuario *Usuario1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(iniciarsesion(usuario.Dpi, usuario.Password))
}

func Ggetusuario(DPI string) Usuario {
	User := new(Usuario)
	for i := 0; i < len(Users); i++ {
		if strconv.Itoa(Users[i].Dpi) == DPI {
			return Users[i]
		}
	}
	return *User
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario *Usuario1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Ggetusuario(usuario.Dpi))
}

var Users []Usuario

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var Dpis []int
	for i := 0; i < len(Users); i++ {
		Dpis = append(Dpis, Users[i].Dpi)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Dpis)
}

func Cargar(w http.ResponseWriter, r *http.Request) {
	ArbolMerkle := NewArbolMerkle()
	c := Usuarios{}
	Users = nil
	arbol := NewArbol(5)
	Administrador(arbol)
	lector, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(c.Usuarios); i++ {
		if !Buscar(c.Usuarios[i].Dpi) {
			arbol.Insertar(NewKey(c.Usuarios[i].Dpi, c.Usuarios[i].Correo,
				c.Usuarios[i].Password, c.Usuarios[i].Nombre, c.Usuarios[i].Cuenta))
			Users = append(Users, *c.Usuarios[i])
			//--------------------------------------------------------------------------------------------------Insercion a arbol merkle
			//---------------------------------------------------------------------------- HASH: DPI +  NOMBRE + CORREO
			hash := Hash(strconv.Itoa(c.Usuarios[i].Dpi) + c.Usuarios[i].Nombre + c.Usuarios[i].Correo)
			ArbolMerkle.Insertar(Hash(hash), c.Usuarios[i].Nombre, c.Usuarios[i].Correo, c.Usuarios[i].Dpi)
		}
	}
	arbol.Graficar("Sin")
	arbol.Graficar("Cif")
	arbol.Graficar("CifSen")
	ArbolMerkle.GraficarMerkle()
}

func registrar(Usuario Usuario) {
	if !Buscar(Usuario.Dpi) {
		arbol := NewArbol(5)
		Users = append(Users, Usuario)
		for i := 0; i < len(Users); i++ {
			arbol.Insertar(NewKey(Users[i].Dpi, Users[i].Correo,
				Users[i].Password, Users[i].Nombre, Users[i].Cuenta))
		}
		arbol.Graficar("Sin")
		arbol.Graficar("Cif")
		arbol.Graficar("CifSen")
	} else {
		fmt.Println("Este usuario ya existe")
	}
}

func eliminar(DPI int) {
	var Nuevo []Usuario
	if Buscar(DPI) {
		arbol := NewArbol(5)
		for i := 0; i < len(Users); i++ {
			if Users[i].Dpi != DPI {
				arbol.Insertar(NewKey(Users[i].Dpi, Users[i].Correo,
					Users[i].Password, Users[i].Nombre, Users[i].Cuenta))
				Nuevo = append(Nuevo, Users[i])
			}
		}
		arbol.Graficar("Sin")
		arbol.Graficar("Cif")
		arbol.Graficar("CifSen")
	} else {
		fmt.Println("Este usuario no existe")
	}
	Users = Nuevo
}

var Dpis []int

func Eliminar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registrar")
	var usuario *Usuario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	eliminar(usuario.Dpi)
	for i := 0; i < len(Users); i++ {
		Dpis = append(Dpis, Users[i].Dpi)
	}
}

func Registrar(w http.ResponseWriter, r *http.Request) {
	var usuario *Usuario1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	User := new(Usuario)
	DPI, err := strconv.Atoi(usuario.Dpi)
	User.Dpi = DPI
	User.Nombre = usuario.Nombre
	User.Correo = usuario.Correo
	User.Password = usuario.Password
	User.Cuenta = "Usuario"
	registrar(*User)
}

//-------------------------------------------------------------------------------------------------Cifrado
func Encriptar(texto string) string {
	key := []byte("tercerafaseeddfechadeentregaelsa")
	plaintext := []byte(texto)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("gopostmedium")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func EncripCont(texto string) string {
	sum := sha256.Sum256([]byte(texto))
	return fmt.Sprintf("%x", sum)
}

//------------------------------------------------------------------------------------------------Graficar
func graficar(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int, modo string) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++
			} else {
				if modo == "Sin" {
					fmt.Fprintf(cad, "<f%d>%s|", j,
						"DPI: "+strconv.Itoa(actual.Keys[i].DPI)+"\\n"+
							"Nombre: "+actual.Keys[i].Nombre+"\\n"+
							"Correo: "+actual.Keys[i].Correo+"\\n"+
							"Contraseña: "+actual.Keys[i].Contrasenia)
				} else if modo == "Cif" {
					fmt.Fprintf(cad, "<f%d>%s|", j,
						"DPI: "+Encriptar(strconv.Itoa(actual.Keys[i].DPI))+"\\n"+
							"Nombre: "+Encriptar(actual.Keys[i].Nombre)+"\\n"+
							"Correo: "+Encriptar(actual.Keys[i].Correo)+"\\n"+
							"Contraseña: "+EncripCont(actual.Keys[i].Contrasenia))
				} else {
					fmt.Fprintf(cad, "<f%d>%s|", j,
						"Nombre: "+actual.Keys[i].Nombre+"\\n"+
							"DPI: "+Encriptar(strconv.Itoa(actual.Keys[i].DPI))+"\\n"+
							"Correo: "+Encriptar(actual.Keys[i].Correo)+"\\n"+
							"Contraseña: "+EncripCont(actual.Keys[i].Contrasenia))
				}
				j++
				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficar(actual.Keys[i].Izquierdo, cad, arr, actual, ji, modo)
		ji++
		ji++
		graficar(actual.Keys[i].Derecho, cad, arr, actual, ji, modo)
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}

func (this *Arbol) Graficar(modo string) {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\n")
	m := make(map[string]*Nodo)
	graficar(this.Raiz, &builder, m, nil, 0, modo)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("diagrama.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println("Arbol B")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./diagrama.dot").Output()
	mode := int(0772)
	if modo == "Sin" {
		ioutil.WriteFile("../frontend/src/assets/img/ArbolCuentasSin.png", cmd, os.FileMode(mode))
	} else if modo == "Cif" {
		ioutil.WriteFile("../frontend/src/assets/img/ArbolCuentasCif.png", cmd, os.FileMode(mode))
	} else {
		ioutil.WriteFile("../frontend/src/assets/img/ArbolCuentasCifSen.png", cmd, os.FileMode(mode))
	}
}

//-------------------------------------------------------------------------------------------------- ARBOL MERKLE
/*
type Key struct {
	DPI         int
	Correo      string
	Contrasenia string
	Nombre      string
	Tipo        string
	Izquierdo   *Nodo
	Derecho     *Nodo
}
*/
type NodoA struct {
	Hash      string
	Nombre    string
	Correo    string
	DPI       int
	Derecha   *NodoA
	Izquierda *NodoA
}

type ArbolMerkle struct {
	Raiz *NodoA
}

func newNodoA(Hash string, Nombre, Correo string, DPI int, Derecha *NodoA, Izquierda *NodoA) *NodoA {
	return &NodoA{Hash, Nombre, Correo, DPI, Derecha, Izquierda}
}

func NewArbolMerkle() *ArbolMerkle {
	return &ArbolMerkle{}
}

func (this *ArbolMerkle) Insertar(Hash1, Nombre, Correo string, DPI int) {
	n := newNodoA(Hash1, Nombre, Correo, DPI, nil, nil)
	if this.Raiz == nil {
		lista := list.New()
		lista.PushBack(n)
		lista.PushBack(newNodoA(Hash(""), "", "", -1, nil, nil))
		this.construirArbolMerkle(lista)
	} else {
		lista := this.obtenerLista()
		lista.PushBack(n)
		this.construirArbolMerkle(lista)
	}
}

func (this *ArbolMerkle) obtenerLista() *list.List {
	lista := list.New()
	obtenerLista(lista, this.Raiz.Izquierda)
	obtenerLista(lista, this.Raiz.Derecha)
	return lista
}

func obtenerLista(lista *list.List, actual *NodoA) {
	if actual != nil {
		obtenerLista(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Hash != Hash("") {
			lista.PushBack(actual)
		}
		obtenerLista(lista, actual.Derecha)
	}
}

func (this *ArbolMerkle) construirArbolMerkle(lista *list.List) {
	size := float64(lista.Len())
	cant := 1
	for (size / 2) > 1 {
		cant++
		size = size / 2
	}
	NodoAstot := math.Pow(2, float64(cant))
	for lista.Len() < int(NodoAstot) {
		lista.PushBack(newNodoA(Hash(""), "", "", -1, nil, nil))
	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		NodoA1 := primero.Value.(*NodoA)
		NodoA2 := segundo.Value.(*NodoA)
		h := ""
		if NodoA2.Hash != "" {
			h = NodoA1.Hash + "\\n" + NodoA2.Hash
		} else {
			h = NodoA1.Hash
		}
		a := Hash(h)
		nuevo := newNodoA(a, h, "", -1, NodoA2, NodoA1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoA)
}

func (this *ArbolMerkle) GraficarMerkle() {
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
	err := ioutil.WriteFile("../frontend/src/assets/img/MerkleUsuarios.dot", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "../frontend/src/assets/img/MerkleUsuarios.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("../frontend/src/assets/img/MerkleUsuarios.png", cmd, os.FileMode(mode))
	fmt.Println("MerkleUsuarios")
}

func (this *ArbolMerkle) generar(cadena *strings.Builder, padre *NodoA, actual *NodoA, Raiz *NodoA) {
	if actual != nil {
		if actual.Hash != Hash("") {
			if actual.DPI >= 0 {
				fmt.Fprintf(cadena, "node%p[label=\"{%s |DPI: %v \\nNombre: %s \\nCorreo: %s}\", fillcolor=\"green\"];\n",
					&(*actual), actual.Hash, actual.DPI, actual.Nombre, actual.Correo)
			} else {
				fmt.Fprintf(cadena, "node%p[label=\"{%s | %s}\", fillcolor=\"green\"];\n", &(*actual), actual.Hash, actual.Nombre)
			}
		} else {
			fmt.Fprintf(cadena, "node%p[label=\"{%s |%s \\n %s \\n %v}\", fillcolor=\"gray\", color=\"red\"];\n",
				&(*actual), actual.Hash, actual.Nombre, actual.Correo, actual.DPI)
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
