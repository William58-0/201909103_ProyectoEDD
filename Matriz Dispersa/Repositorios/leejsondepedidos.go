package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
			/*
				Produc.Fecha = c.Pedidos[i].Fecha
				Produc.Tienda = c.Pedidos[i].Tienda
				Produc.Departamento = c.Pedidos[i].Departamento
				Produc.Calificacion = c.Pedidos[i].Calificacion
			*/
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
		}
	}
	fmt.Println("vector de productos: ")
	fmt.Println(productos)
}

func main() {
	leer()
}
