import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Tienda } from "../../models/Tienda/Tienda";
import { Producto } from "../../models/Producto/Producto";
import { Busqueda } from "../../models/Busqueda/Busqueda";

@Component({
  selector: 'app-Tiendas',
  templateUrl: './Tiendas.component.html',
  styleUrls: ['./Tiendas.component.css']
})
export class TiendasComponent implements OnInit {

  Tiendas: Tienda[] = []
  Productos: Producto[] = []
  Carrito: Producto[]=[]
  mostrarMensajeError = false
  mostrarMensaje = false
  mensajeError = ''
  Estado: string;
  Arbol: string;

  constructor(private DatosService: DatosService) {
    this.DatosService.cargartiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      console.log(dataList)
      console.log(this.Tiendas[0].Nombre)
      this.Estado = "Tiendas"
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo cargar la lista de tiendas'
    })
  }

  ngOnInit(): void {
  }

  cargartiendas() {
    this.DatosService.cargartiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      console.log(dataList)
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo cargar la lista de tiendas'
    })
  }

  guardarCurso(Producto: Producto){
    this.DatosService.Comprar(Producto).subscribe((res:any)=>{
      //this.mostrarMensaje=true
    }, (err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo guardar el curso aprobado'
    })
  }

  INV(Tienda: Tienda) {
    var Busqueda: Busqueda = {
      Tienda: Tienda.Nombre,
      Departamento: Tienda.Departamento,
      Calificacion: Tienda.Calificacion
    }
    this.Arbol = Tienda.Nombre + "---" + Tienda.Departamento + "---" + Tienda.Calificacion+".png"
    console.log(Busqueda);
    this.DatosService.getinv(Busqueda).subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      console.log(this.Productos[0].Nombre)
      this.Estado = "Productos"
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo guardar el curso aprobado'
    })
  }

  Comprar(Producto: Producto){
    this.DatosService.Comprar(Producto).subscribe((res:any)=>{
      this.mostrarMensaje=true
      if(Producto.Cantidad>0){
        Producto.Cantidad--
      }
    }, (err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo guardar el curso aprobado'
    })
  }

  Regresar() {
    this.Estado = "Tiendas"
  }

  desactivarMensaje() {
    this.mostrarMensaje = false
    this.mostrarMensajeError = false
  }

  mensaje(carnet) {
    console.log("carnet " + carnet)
  }

}

