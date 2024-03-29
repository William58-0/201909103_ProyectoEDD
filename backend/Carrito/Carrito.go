package Carrito

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"../AVL"
	"../ArbolB"
	"../MatrizDispersa"
)

var Carrito []AVL.Producto1

var Todo AVL.Todo

func RestarProducto(w http.ResponseWriter, r *http.Request) {
	var prod *AVL.Producto1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &prod)
	if prod.Cantidad > 0 {
		Productos := AVL.Todo1.Productos
		for i := 0; i < len(Productos); i++ {
			if prod.Nombre == Productos[i].Nombre && prod.Codigo == Productos[i].Codigo &&
				prod.Tienda == Productos[i].Tienda && prod.Departamento == Productos[i].Departamento &&
				prod.Calificacion == Productos[i].Calificacion {
				fmt.Println("Nombre: " + prod.Nombre)
				fmt.Println("Cantidad de productos antes: " + strconv.Itoa(prod.Cantidad))

				Productos[i].Cantidad--
				prod.Cantidad--
				Carrito = append(Carrito, *prod)
				Todo.Productos = Carrito
				fmt.Println("Cantidad de productos luego: " + strconv.Itoa(Productos[i].Cantidad))
				break
			}
		}
	}
}

func Eliminar(Arr []AVL.Producto1, prod *AVL.Producto1) {
	Eliminado := false
	var nuevo []AVL.Producto1
	for i := 0; i < len(Arr); i++ {
		if Arr[i].Nombre == prod.Nombre && Arr[i].Tienda == prod.Tienda &&
			Arr[i].Departamento == prod.Departamento && !Eliminado {
			Eliminado = true
		} else {
			nuevo = append(nuevo, Arr[i])
		}
	}
	Carrito = nuevo
	Todo.Productos = Carrito
}

func SumarProducto(w http.ResponseWriter, r *http.Request) {
	var prod *AVL.Producto1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &prod)
	//if prod.Cantidad > 0 {
	Productos := AVL.Todo1.Productos
	for i := 0; i < len(Productos); i++ {
		if prod.Nombre == Productos[i].Nombre && prod.Codigo == Productos[i].Codigo &&
			prod.Tienda == Productos[i].Tienda && prod.Departamento == Productos[i].Departamento &&
			prod.Calificacion == Productos[i].Calificacion {
			//fmt.Println("Nombre: " + prod.Nombre)
			//fmt.Println("Cantidad de productos antes: " + strconv.Itoa(prod.Cantidad))

			Productos[i].Cantidad++
			prod.Cantidad++
			Eliminar(Carrito, prod)
			//fmt.Println("Cantidad de productos luego: " + strconv.Itoa(Productos[i].Cantidad))
			break

		}
	}
}

func GenerarPedido(w http.ResponseWriter, r *http.Request) {
	var productos []*MatrizDispersa.Producto
	var nuevo []MatrizDispersa.Producto
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &productos)
	for i := 0; i < len(productos); i++ {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02d",
			t.Day(), t.Month(), t.Year())
		productos[i].Fecha = fecha
		productos[i].Cliente = ArbolB.SesionActual
		MatrizDispersa.Productos = append(MatrizDispersa.Productos, *productos[i])
		nuevo = append(nuevo, *productos[i])
	}
	Carrito = nil
	Todo.Productos = Carrito
	MatrizDispersa.Actualizar(nuevo)
}

func GetCarrito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todo)
}
