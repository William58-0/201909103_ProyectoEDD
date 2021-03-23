import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";

@Component({
  selector: 'app-Carrito',
  templateUrl: './Carrito.component.html',
  styleUrls: ['./Carrito.component.css']
})
export class CarritoComponent implements OnInit {

  Productos: Producto[]=[]
  mostrarMensajeError = false
  mostrarMensaje = false
  mensajeError = ''

  constructor(private DatosService: DatosService) {
    this.DatosService.CargarCarro().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      console.log(this.Productos[0])
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo guardar el curso aprobado'
    })
  }

  ngOnInit(): void {
  }

  Devolver(Producto:Producto){
    this.DatosService.Devolver(Producto).subscribe((res:any)=>{
      this.mostrarMensaje=true
        Producto.Cantidad++
      this.removeItemFromArr(this.Productos,Producto)
    }, (err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo guardar el curso aprobado'
    })
  }

 removeItemFromArr ( arr, item ) {
    var i = arr.indexOf( item );
    if ( i !== -1 ) {
        arr.splice( i, 1 );
    }
}

  cargar() {
    this.DatosService.CargarCarro().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      console.log(this.Productos[0].Nombre)
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo cargar la lista de tiendas'
    })
  }

  desactivarMensaje() {
    this.mostrarMensaje = false
    this.mostrarMensajeError = false
  }


}