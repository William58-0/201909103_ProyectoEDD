package Carrito

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../AVL"
)

var Carrito []AVL.Producto1

var Todo AVL.Todo

func RestarProducto(w http.ResponseWriter, r *http.Request) {
	var prod *AVL.Producto1
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &prod)
	Productos := AVL.Todo1.Productos
	for i := 0; i < len(Productos); i++ {
		if prod.Nombre == Productos[i].Nombre && prod.Codigo == Productos[i].Codigo &&
			prod.Descripcion == Productos[i].Descripcion && prod.Precio == Productos[i].Precio &&
			prod.Imagen == Productos[i].Imagen && prod.Tienda == Productos[i].Tienda &&
			prod.Departamento == Productos[i].Departamento && prod.Calificacion == Productos[i].Calificacion {
			fmt.Println("Nombre: " + prod.Nombre)
			fmt.Println("Cantidad de productos antes: " + strconv.Itoa(prod.Cantidad))
			if prod.Cantidad > 0 {
				Productos[i].Cantidad--
				prod.Cantidad--
				Carrito = append(Carrito, *prod)
				Todo.Productos = Carrito
				fmt.Println("Cantidad de productos luego: " + strconv.Itoa(Productos[i].Cantidad))
				break
			}
		}
	}
	fmt.Println("Comprobar Todo")
	for i := 0; i < len(Todo.Productos); i++ {
		fmt.Println(Todo.Productos[i].Nombre)
		fmt.Println(Todo.Productos[i].Codigo)
	}

	fmt.Println("Carrito Comprobar")
	for i := 0; i < len(Carrito); i++ {
		fmt.Println(Carrito[i].Nombre)
		fmt.Println(Carrito[i].Codigo)
	}
}

func SumarProducto(w http.ResponseWriter, r *http.Request) {
	var prod *AVL.Producto1
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &prod)
	Productos := AVL.Todo1.Productos
	for i := 0; i < len(Productos); i++ {
		if prod.Nombre == Productos[i].Nombre && prod.Codigo == Productos[i].Codigo &&
			prod.Descripcion == Productos[i].Descripcion && prod.Precio == Productos[i].Precio &&
			prod.Imagen == Productos[i].Imagen && prod.Tienda == Productos[i].Tienda &&
			prod.Departamento == Productos[i].Departamento && prod.Calificacion == Productos[i].Calificacion {
			fmt.Println("Nombre: " + prod.Nombre)
			fmt.Println("Cantidad de productos antes: " + strconv.Itoa(prod.Cantidad))
			if prod.Cantidad > 0 {
				Productos[i].Cantidad++
				prod.Cantidad++
				Carrito = append(Carrito[:i], Carrito[i+1:]...)
				Todo.Productos = Carrito
				fmt.Println("Cantidad de productos luego: " + strconv.Itoa(Productos[i].Cantidad))
				break
			}
		}
	}
	fmt.Println("Comprobar Todo")
	for i := 0; i < len(Todo.Productos); i++ {
		fmt.Println(Todo.Productos[i].Nombre)
		fmt.Println(Todo.Productos[i].Codigo)
	}

	fmt.Println("Carrito Comprobar")
	for i := 0; i < len(Carrito); i++ {
		fmt.Println(Carrito[i].Nombre)
		fmt.Println(Carrito[i].Codigo)
	}
}

func GetCarrito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todo)
}
