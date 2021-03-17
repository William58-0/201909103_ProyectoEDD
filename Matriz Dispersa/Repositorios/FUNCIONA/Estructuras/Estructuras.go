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
	Nombre    string
	Tipo      string
	Cola      Cola
	Arriba    *NODO
	Abajo     *NODO
	Izquierda *NODO
	Derecha   *NODO
	Ultimo    *NODO //este se usa solo en valores
	URight    *NODO //este es para nodo0
	UDown     *NODO //este es para nodo0
}

func (Mes *MONTH) GetNodo(Nombre string) *NODO {
	nodo := new(NODO)
	for i := 0; i < len(Mes.Nodos); i++ {
		if Mes.Nodos[i].Nombre == Nombre {
			return &Mes.Nodos[i]
		}
	}
	return nodo
}

func (Mes *MONTH) ExisteNodo(Nombre string) bool {
	for i := 0; i < len(Mes.Nodos); i++ {
		if Mes.Nodos[i].Nombre == Nombre {
			return true
		}
	}
	return false
}

//----------------------------------------------------------------------------------------------                 LISTA
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
	Siguiente *Casilla
}

type Cola struct {
	Nombre  string
	Tamanio int
	Primero *Casilla
	Ultimo  *Casilla
}

func (Cola *Cola) Insertar(Producto *Producto) {
	nuevo := new(Casilla)
	nuevo.Producto = *Producto
	/*
		if Cola.Primero == nil {
			Cola.Primero = nuevo
			Cola.Ultimo = nuevo
		} else {
			Cola.Ultimo.Siguiente = nuevo
			Cola.Ultimo = nuevo
		}
	*/
	if Cola.Primero != nil {
		Cola.Ultimo.Siguiente = nuevo
		Cola.Ultimo = nuevo
	} else {
		Cola.Primero = nuevo
		Cola.Ultimo = nuevo
	}
	Cola.Tamanio++
}

func (Cola *Cola) Extraer() *Casilla {
	aux := Cola.Primero
	if aux.Siguiente != nil {
		Cola.Primero = aux.Siguiente
		aux.Siguiente = nil
	}
	Cola.Tamanio--
	return aux
}
