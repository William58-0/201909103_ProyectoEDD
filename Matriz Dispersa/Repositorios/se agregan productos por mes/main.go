package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./Estructuras"
)

var productos []Estructuras.Producto
var meses []string

func leer() {
	lector, err := ioutil.ReadFile("Pedidos.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Estructuras.Pedidos{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	//En listar solo productos
	for i := 0; i < len(c.Pedidos); i++ {
		for j := 0; j < len(c.Pedidos[i].Productos); j++ {
			Produc := c.Pedidos[i].Productos[j]
			Producto := new(Estructuras.Producto)
			Producto.Nombre = Produc.Nombre
			Producto.Codigo = Produc.Codigo
			Producto.Descripcion = Produc.Descripcion
			Producto.Precio = Produc.Precio
			Producto.Cantidad = Produc.Cantidad
			Producto.Fecha = c.Pedidos[i].Fecha
			Producto.Tienda = c.Pedidos[i].Tienda
			Producto.Departamento = c.Pedidos[i].Departamento
			Producto.Calificacion = c.Pedidos[i].Calificacion
			productos = append(productos, *Producto)
			//agregar el mes a meses[]

			mes := strings.Split(c.Pedidos[i].Fecha, "-")[2] + "-" + strings.Split(c.Pedidos[i].Fecha, "-")[1]
			existe := false
			for k := 0; k < len(meses); k++ {
				if meses[k] == mes {
					existe = true
				}
			}
			if !existe {
				meses = append(meses, mes)
			}
		}
	}
	//se ordenan los meses
	var j int
	var aux string
	n := len(meses)
	for i := 1; i < n; i++ {
		j = i
		aux = meses[i]

		for j > 0 && aux < meses[j-1] {
			meses[j] = meses[j-1]
			j--
		}
		meses[j] = aux
	}
	//se crean las estructuras
	//se crea la lista de años
	anio := ""
	ListaA := new(Estructuras.ListaA)
	for i := 0; i < len(meses); i++ {
		if strings.Split(meses[i], "-")[0] != anio {
			anio = strings.Split(meses[i], "-")[0]
			ListaA.InsertarA(anio)
		}
	}
	//se añaden los meses a los años
	aux2 := ListaA.Primero
	for aux2 != nil {
		ListaM := new(Estructuras.ListaM)
		for i := 0; i < len(meses); i++ {
			if strings.Split(meses[i], "-")[0] == aux2.Anio {
				ListaM.InsertarM(strings.Split(meses[i], "-")[1])
			}
		}
		aux2.Meses = *ListaM
		aux2 = aux2.Siguiente
	}
	aux2 = ListaA.Primero
	for aux2 != nil {
		aux1 := aux2.Meses.Primero
		for aux1 != nil {
			for p := 0; p < len(productos); p++ {
				anio := strings.Split(productos[p].Fecha, "-")[2]
				mes := strings.Split(productos[p].Fecha, "-")[1]
				if anio == aux2.Anio && mes == aux1.Mes {
					aux1.Productos = append(aux1.Productos, productos[p])
					//fmt.Println("Se agregó producto: " + strconv.Itoa(productos[p].Codigo) + " a mes: " + mes + " y año: " + anio)
					//fmt.Println("Fecha: ", productos[p].Fecha)
				}
			}
			aux1 = aux1.Siguiente
		}
		aux2 = aux2.Siguiente
	}

	fmt.Println("vector de productos: ")
	fmt.Println(productos)
}

func main() {
	leer()
}
