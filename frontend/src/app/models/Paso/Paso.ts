import { Producto } from "../Producto/Producto";

export class Paso {

    Numero: number
    Recorrido: string
    Distancia: number
    Recogidos: Producto[]
    Pendientes: Producto[]

    constructor(_Numero: number, _Recorrido: string,
        _Distancia: number, _Recogidos: Producto[], 
        _Pendientes: Producto[]) {
        this.Numero = _Numero
        this.Recorrido = _Recorrido
        this.Distancia = _Distancia
        this.Recogidos = _Recogidos
        this.Pendientes = _Pendientes
    }

}
