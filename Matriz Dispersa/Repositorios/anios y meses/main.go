package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./Estructuras"
)

type Pedidos struct {
	Pedidos []*Principal `json:"Pedidos"`
}

type Principal struct {
	Fecha        string      `json:"Fecha"`
	Tienda       string      `json:"Tienda"`
	Departamento string      `json:"Departamento"`
	Calificacion int         `json:"Calificacion"`
	Productos    []*Producto `json:"Productos"`
}

type Producto struct {
	Nombre      string `json:"Nombre"`
	Codigo      int    `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Precio      string `json:"Precio"`
	Cantidad    int    `json:"Cantidad"`
	//estos son extras
	Fecha        string
	Tienda       string
	Departamento string
	Calificacion int
}

var productos []Producto
var meses []string

func leer() {
	lector, err := ioutil.ReadFile("Pedidos.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Pedidos{}
	err = json.Unmarshal(lector, &c)
	if err != nil {
		log.Fatal(err)
	}
	//En listar solo productos
	for i := 0; i < len(c.Pedidos); i++ {
		for j := 0; j < len(c.Pedidos[i].Productos); j++ {
			Produc := c.Pedidos[i].Productos[j]
			Producto := new(Producto)
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
	//comprobar que los años tengan sus meses
	/*
		aux2 = ListaA.Primero
		for aux2 != nil {
			aux1 := aux2.Meses.Primero
			for aux1 != nil {
				fmt.Println("Año: " + aux2.Anio + "Mes: " + aux1.Mes)
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Siguiente
		}
	*/

	fmt.Println("vector de productos: ")
	fmt.Println(productos)
}

func main() {
	leer()
}
