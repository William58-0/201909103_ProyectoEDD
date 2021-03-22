export class Producto {

    Nombre: string
    Codigo: number
    Descripcion: string
    Precio: string
    Cantidad: number
    Imagen: string
    //estos son extras
    Fecha: string
    Tienda: string
    Departamento: string
    Calificacion: number

    constructor(_Nombre: string, _Codigo: number, _Descripcion: string,
        _Precio: string, _Cantidad: number,_Imagen:string, _Fecha: string,
        _Tienda: string, _Departamento: string, _Calificacion: number) {
        this.Nombre = _Nombre
        this.Codigo = _Codigo
        this.Descripcion = _Descripcion
        this.Precio = _Precio
        this.Cantidad = _Cantidad
        this.Imagen = _Imagen
        this.Fecha = _Fecha
        this.Tienda = _Tienda
        this.Departamento = _Departamento
        this.Calificacion = _Calificacion
    }

}
