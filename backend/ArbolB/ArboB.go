package ArbolB

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
			if placed == false {
				nodo.Colocar(0, newKey)
				nodo.Keys[1].Izquierdo = newKey.Derecho
			}
			index = i
			break
		}
	}
	return index
}

//----------------------------------------------------------------------------vVerificar que los datos sean validos
/*
var Valido bool

funcvverificar(actual *Nodo, DPIbuscado int, ContraseniaBuscada string) {
	if actual != nil && !Valido {
		for i := 0; i < actual.Max; i++ {
			if actual.Keys[i] == nil {
				break
			}
			if actual.Keys[i] != nil {
				fmt.Println("Correo: ", actual.Keys[i].Correo, " DPI: ", actual.Keys[i].DPI)
				if actual.Keys[i].DPI == DPIbuscado && actual.Keys[i].Contrasenia == ContraseniaBuscada {
					fmt.Println("Valido")
					Valido = true
					return
				}
			}
		}
		for i := 0; i < actual.Max; i++ {
			if actual.Keys[i] != nil {
			vverificar(actual.Keys[i].Izquierdo, DPIbuscado, ContraseniaBuscada)
			vverificar(actual.Keys[i].Derecho, DPIbuscado, ContraseniaBuscada)
			} else {
				break
			}
		}
	}
	if Valido {
		return
	}
}

func (this *Arbol)vVerificar(DPIbuscado int, ContraseniaBuscada string) bool {
	Valido = false
vverificar(this.Raiz, DPIbuscado, ContraseniaBuscada)
	if Valido {
		return true
	} else {
		return false
	}
}
*/

//-----------------------------------------------------------------------------------Buscar si existe un usuario
/*
var Existe bool

func buscar(actual *Nodo, DPIbuscado int) {
	if actual != nil && !Existe {
		for i := 0; i < actual.Max; i++ {
			if actual.Keys[i] == nil {
				break
			}
			if actual.Keys[i] != nil {
				fmt.Println("Correo: ", actual.Keys[i].Correo, " DPI: ", actual.Keys[i].DPI)
				if actual.Keys[i].DPI == DPIbuscado {
					fmt.Println("Valido")
					Existe = true
					return
				}
			}
		}
		for i := 0; i < actual.Max; i++ {
			if actual.Keys[i] != nil {
				buscar(actual.Keys[i].Izquierdo, DPIbuscado)
				buscar(actual.Keys[i].Derecho, DPIbuscado)
			} else {
				break
			}
		}
	}
	if Existe {
		return
	}
}

func (this *Arbol) Buscar(DPIbuscado int) bool {
	Existe = false
	buscar(this.Raiz, DPIbuscado)
	if Existe {
		return true
	} else {
		return false
	}
}
*/

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

func iniciarsesion(DPI string, password string) Usuario {
	User := new(Usuario)
	for i := 0; i < len(Users); i++ {
		if strconv.Itoa(Users[i].Dpi) == DPI && Users[i].Password == password {
			fmt.Println(Users[i])
			return Users[i]
		}
	}
	return *User
}

func IniciarSesion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IniciarSesion")
	fmt.Println("Usuarios: ", Users)
	var usuario *Usuario1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	fmt.Println("Buscar:")
	fmt.Println(usuario.Dpi)
	fmt.Println(usuario.Password)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(iniciarsesion(usuario.Dpi, usuario.Password))
}

func getusuario(DPI string) Usuario {
	User := new(Usuario)
	for i := 0; i < len(Users); i++ {
		if strconv.Itoa(Users[i].Dpi) == DPI {
			fmt.Println(Users[i])
			return Users[i]
		}
	}
	return *User
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUsuario")
	fmt.Println("Usuarios: ", Users)
	var usuario *Usuario1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	fmt.Println("Buscar:")
	fmt.Println(usuario.Dpi)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getusuario(usuario.Dpi))
}

var Users []Usuario

func Cargar(w http.ResponseWriter, r *http.Request) {
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
		}
	}
	arbol.Graficar("Sin")
	arbol.Graficar("Cif")
	arbol.Graficar("CifSen")
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
	fmt.Println(DPI)
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

func Eliminar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registrar")
	fmt.Println("Usuarios Antes: ", Users)
	var usuario *Usuario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	fmt.Println(usuario)
	eliminar(usuario.Dpi)
	fmt.Println("Usuarios Despues: ", Users)
}

func Registrar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registrar")
	fmt.Println("Usuarios: ", Users)
	var usuario *Usuario1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &usuario)
	fmt.Println(usuario)
	User := new(Usuario)
	DPI, err := strconv.Atoi(usuario.Dpi)
	User.Dpi = DPI
	User.Nombre = usuario.Nombre
	User.Correo = usuario.Correo
	User.Password = usuario.Password
	User.Cuenta = "Usuario"
	registrar(*User)
	fmt.Println(Users)
}

//-------------------------------------------------------------------------------------------------Cifrado
func Encriptar(texto string) string {
	//key := []byte("keygopostmediumkeygopostmediumke")
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

/*
func Desencriptar(texto string) string {
	//key := []byte("keygopostmediumkeygopostmediumke")
	key := []byte("tercerafaseeddfechadeentregaelsabadoantesdemedianoche")
	ciphertext, _ := hex.DecodeString(texto)
	nonce := []byte("gopostmedium")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}
*/

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
	//fmt.Println("actual", &(*actual))
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
							"Correo: "+actual.Keys[i].Correo)
				} else if modo == "Cif" {
					fmt.Fprintf(cad, "<f%d>%s|", j,
						"DPI: "+Encriptar(strconv.Itoa(actual.Keys[i].DPI))+"\\n"+
							"Nombre: "+Encriptar(actual.Keys[i].Nombre)+"\\n"+
							"Correo: "+Encriptar(actual.Keys[i].Correo))
				} else {
					fmt.Fprintf(cad, "<f%d>%s|", j,
						"Nombre: "+actual.Keys[i].Nombre+"\\n"+
							"DPI: "+Encriptar(strconv.Itoa(actual.Keys[i].DPI))+"\\n"+
							"Correo: "+Encriptar(actual.Keys[i].Correo))
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
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
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
