package Estructuras

//-------------------------------------------------------------------------------------------		    PARA LEER EL JSON
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

//--------------------------------------------------------------------------------------------			OBJETOS
type YEAR struct {
	Anio      string
	Meses     ListaM
	Siguiente *YEAR
	Anterior  *YEAR
}

type MONTH struct {
	Mes       string
	Productos []Producto
	Nodos     []NODO
	Siguiente *MONTH
	Anterior  *MONTH
	//Matriz matriz
}

type NODO struct {
	Fecha        string
	Departamento string
	Cola         Cola
	Arriba       *NODO
	Abajo        *NODO
	Izquierda    *NODO
	Derecha      *NODO
	//este se usa solo en valores
	Ultimo *NODO
}

//-------------------------------------------------------------------------------------------                 LISTA
type ListaA struct {
	Primero *YEAR
	Ultimo  *YEAR
	Tamanio int
}

type ListaM struct {
	Primero *MONTH
	Ultimo  *MONTH
	Tamanio int
}

//----------------------------------------------------------------------------------------------				FUNCIONES DE LISTAS
//insertar año
func (Lista *ListaA) InsertarA(Anio string) {
	nuevo := new(YEAR)
	nuevo.Anio = Anio
	if Lista.Primero == nil {
		Lista.Primero = nuevo
		Lista.Ultimo = nuevo
	} else {
		Lista.Ultimo.Siguiente = nuevo
		Lista.Ultimo.Siguiente.Anterior = Lista.Ultimo
		Lista.Ultimo = nuevo
	}
	Lista.Tamanio++
}

//insertar mes
func (Lista *ListaM) InsertarM(Mes string) {
	nuevo := new(MONTH)
	nuevo.Mes = Mes
	if Lista.Primero == nil {
		Lista.Primero = nuevo
		Lista.Ultimo = nuevo
	} else {
		Lista.Ultimo.Siguiente = nuevo
		Lista.Ultimo.Siguiente.Anterior = Lista.Ultimo
		Lista.Ultimo = nuevo
	}
	Lista.Tamanio++
}

//obtener año
func (Lista *ListaA) getA(Anio string) YEAR {
	aux := Lista.Primero
	for aux != nil {
		if aux.Anio == Anio {
			return *aux
		}
		aux = aux.Siguiente
	}
	return *aux
}

//obtener mes
func (Lista *ListaM) getM(Mes string) MONTH {
	aux := Lista.Primero
	for aux != nil {
		if aux.Mes == Mes {
			return *aux
		}
		aux = aux.Siguiente
	}
	return *aux
}

//buscar año
func (Lista *ListaA) BuscarA(Anio string) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Anio == Anio {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}

//buscar mes
func (Lista *ListaM) BuscarM(Mes string) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Mes == Mes {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}

//--------------------------------------------------------------------------------------------------------------------COLA
type Casilla struct {
	Producto  Producto
	siguiente *Casilla
}

type Cola struct {
	Fecha        string
	Departamento string
	primero      *Casilla
	ultimo       *Casilla
}

func (Lista *Cola) Insertar(Producto Producto) {
	nuevo := new(Casilla)
	nuevo.Producto = Producto
	if Lista.primero == nil {
		Lista.primero = nuevo
		Lista.ultimo = nuevo
	} else {
		Lista.ultimo.siguiente = nuevo
		Lista.ultimo = nuevo
	}
}

func (Lista *Cola) Extraer() *Casilla {
	aux := Lista.primero
	if aux.siguiente != nil {
		Lista.primero = aux.siguiente
		aux.siguiente = nil
	}
	return aux
}
