package Estructuras

//-----------------------------------------------------------------------------------                  ESTRUCTURAS ORDENADAS
type Data struct {
	Datos []*Principal `json:"Datos"`
}

type Principal struct {
	Indice        string `json:"Indice"`
	Departamentos []*Dep `json:"Departamentos"`
}

type Dep struct {
	Nombre  string    `json:"Nombre"`
	Tiendas []*Tienda `json:"Tiendas"`
}

//un sinonimo de nodo
type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
	Departamento string `json:"Departamento"`
	Siguiente    *Tienda
	Anterior     *Tienda
}

//---------------------------------------------------------------------------------------     		para cargar tiendas en react
type Todo struct {
	Tiendas []Tienda1 `json:"Tiendas"`
}

//----------------------------------------------------------------------------------------					para buscary/o eliminar
type Objetivo struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

type ObjetivoE struct {
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	Calificacion int    `json:"Calificacion"`
}

type Salida struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
}

//---------------------------------------------------------------------------------------------				para generar json
type Data1 struct {
	Datos []Principal1
}

type Principal1 struct {
	Indice        string
	Departamentos []Dep1
}

type Dep1 struct {
	Nombre  string
	Tiendas []Tienda1
}

type Tienda1 struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Departamento string `json:"Departamento"`
	Logo         string `json:"Logo"`
}

//-------------------------------------------------------------------------------------------                 LISTA
type Lista struct {
	Categoria    string
	Indice       string
	Calificacion int
	Tamanio      int
	Primero      *Tienda
	Ultimo       *Tienda
}

func (Lista *Lista) Insertar(Nombre string, Descripcion string, Contacto string, Calificacion int, Departamento string, Logo string) {
	nuevo := new(Tienda)
	nuevo.Nombre = Nombre
	nuevo.Descripcion = Descripcion
	nuevo.Contacto = Contacto
	nuevo.Calificacion = Calificacion
	nuevo.Departamento = Departamento
	nuevo.Logo = Logo
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

func (Lista *Lista) Ordenar() []string {
	aux1 := Lista.Primero
	var vector []string
	for aux1 != nil {
		vector = append(vector, aux1.Nombre)
		aux1 = aux1.Siguiente
	}
	var j int
	var aux string
	n := len(vector)
	for i := 1; i < n; i++ {
		j = i
		aux = vector[i]

		for j > 0 && aux < vector[j-1] {
			vector[j] = vector[j-1]
			j--
		}
		vector[j] = aux
	}
	return vector
}

func (Lista *Lista) Eliminar(Nombre string, Calificacion int) bool {
	aux := Lista.Primero
	for aux != nil {
		if aux.Nombre == Nombre && aux.Calificacion == Calificacion {
			if Lista.Tamanio == 1 {
				Lista.Primero = nil
				Lista.Tamanio--
				return true
			}
			if aux == Lista.Primero {
				Lista.Primero = aux.Siguiente
				aux.Siguiente = nil
				Lista.Primero.Anterior = nil
				Lista.Tamanio--
				return true
			} else if aux == Lista.Ultimo {
				Lista.Ultimo = aux.Anterior
				aux.Anterior = nil
				Lista.Ultimo.Siguiente = nil
				Lista.Tamanio--
				return true
			} else {
				aux.Anterior.Siguiente = aux.Siguiente
				aux.Siguiente.Anterior = aux.Anterior
				Lista.Tamanio--
				return true
			}
		}
		aux = aux.Siguiente
	}
	return false
}

func (Lista *Lista) Buscar(Nombre string, Calificacion int) Salida {
	Salida := new(Salida)
	aux := Lista.Primero
	for aux != nil {
		if aux.Nombre == Nombre && aux.Calificacion == Calificacion {
			Salida.Nombre = aux.Nombre
			Salida.Descripcion = aux.Descripcion
			Salida.Contacto = aux.Contacto
			Salida.Calificacion = aux.Calificacion
			return *Salida
		}
		aux = aux.Siguiente
	}
	return *Salida
}
