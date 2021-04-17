import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";
import { Usuario } from "../../models/Usuario/Usuario";
import { Paso } from "../../models/Paso/Paso";

@Component({
  selector: 'app-Carrito',
  templateUrl: './Carrito.component.html',
  styleUrls: ['./Carrito.component.css']
})
export class CarritoComponent implements OnInit {

  Productos: Producto[] = []
  Usuario: Usuario;
  Estado: string;
  //Para el recorrido
  Grafo: string="GrafoInicial";
  NumeroPaso: number;
  TiposGrafo = ["Grafo Inicial", "Pasos", "Recorrido Completo"]
  TipoGrafo = "Grafo Inicial"
  Pasos: Paso[] = []
  NPaso: number = 0
  Pendientes: Producto[] = []
  Recogidos: Producto[] = []
  Recorrido: string = ""
  Distancia: number = 0

  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {
    this.DatosService.CargarCarro().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      this.Estado = "Carrito"
      //console.log(this.Productos[0])>
    }, (err) => {
      console.log("error")
    })
    this.DatosService.GetRecorrido().subscribe((dataList: any) => {
      this.Pasos = dataList.Pasos
      console.log(dataList)
      //console.log(this.Productos[0])>
    }, (err) => {
      console.log("error")
    })
  }

  ngOnInit(): void {
    this.GetUsuario(this.route.snapshot.paramMap.get('Dpi'))
  }

  GetUsuario(Dpi) {
    var Busqueda = {
      Dpi: Dpi
    }
    this.DatosService.GetUsuario(Busqueda).subscribe(Usuario => {
      this.Usuario = Usuario;
      console.log(Usuario)
      /*
      if (data.Usuario != null) {

      } else {
        alert("No existe usuario")
      }
      */
    },
      error => {
        console.log(error);
      });
  }

  Devolver(Producto: Producto) {
    this.DatosService.Devolver(Producto).subscribe((res: any) => {
      Producto.Cantidad++
      this.Productos = this.removeItemFromArr(this.Productos, Producto)
    }, (err) => {
      console.log("Error")
    })
  }

  ARecorrido() {
    this.Estado = "Recorrido"
  }

  removeItemFromArr(arr: Producto[], item: Producto) {
    var i = arr.indexOf(item);
    if (i !== -1) {
      arr.splice(i, 1);
    }
    return arr
  }

  GenerarPedido(Productos: Producto[]) {
    this.DatosService.GenerarPedido(Productos).subscribe((res: any) => {
      this.Productos = null
    }, (err) => {
      console.log("Error")
    })
  }

  GenerarRecorrido(Productos: Producto[]) {
    this.DatosService.GenerarRecorrido(Productos).subscribe(Pasos => {
      this.Pasos = Pasos.Pasos;
      console.log(Pasos)
    },
      error => {
        console.log(error);
      })
  }

  ATiendas() {
    window.location.href = "/Tiendas/" + this.Usuario.Dpi
  }

Avanzar(){
  if (this.NPaso+1<this.Pasos.length){
    this.NPaso++
    this.Pendientes=this.Pasos[this.NPaso].Pendientes
    this.Recogidos=this.Pasos[this.NPaso].Recogidos
    this.Recorrido=this.Pasos[this.NPaso].Recorrido
    this.Distancia=this.Pasos[this.NPaso].Distancia
    this.Grafo="Paso"+this.NPaso
  }
}

Retroceder(){
  if (this.NPaso-1>=0){
    this.NPaso--
    this.Pendientes=this.Pasos[this.NPaso].Pendientes
    this.Recogidos=this.Pasos[this.NPaso].Recogidos
    this.Recorrido=this.Pasos[this.NPaso].Recorrido
    this.Distancia=this.Pasos[this.NPaso].Distancia
    this.Grafo="Paso"+this.NPaso
  }
}

GrafoInicial(){
  this.TipoGrafo="Grafo Inicial"
  this.Grafo="GrafoInicial"
}

Pasoss(){
  this.NPaso=0;
  this.Pendientes=this.Pasos[this.NPaso].Pendientes
  this.Recogidos=this.Pasos[this.NPaso].Recogidos
  this.Recorrido=this.Pasos[this.NPaso].Recorrido
  this.Distancia=this.Pasos[this.NPaso].Distancia
  this.TipoGrafo="Pasos"
  this.Grafo="Paso"+this.NPaso
}

RecorridoCompleto(){
  this.NPaso=this.Pasos.length-1
  this.Pendientes=this.Pasos[this.NPaso].Pendientes
  this.Recogidos=this.Pasos[this.NPaso].Recogidos
  this.Recorrido=this.Pasos[this.NPaso].Recorrido
  this.Distancia=this.Pasos[this.NPaso].Distancia
  this.TipoGrafo="Recorrido Completo"
  this.Grafo="RecorridoCompleto"
}




  Prueba(){

  }

}