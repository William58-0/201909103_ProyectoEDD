import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";

@Component({
  selector: 'app-Carrito',
  templateUrl: './Carrito.component.html',
  styleUrls: ['./Carrito.component.css']
})
export class CarritoComponent implements OnInit {

  Productos: Producto[] = []

  constructor(private DatosService: DatosService) {
    this.DatosService.CargarCarro().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      //console.log(this.Productos[0])>
    }, (err) => {
      console.log("error")
    })
  }

  ngOnInit(): void {
  }

  Devolver(Producto: Producto) {
    this.DatosService.Devolver(Producto).subscribe((res: any) => {
      Producto.Cantidad++
      this.Productos=this.removeItemFromArr(this.Productos, Producto)
    }, (err) => {
      console.log("Error")
    })
  }
/*
  removeItemFromArr(arr: Producto[], item: Producto) {
    var Eliminado = false
    var nuevo:Producto[]=[]
    for (let i = 0; i < arr.length; i++) {
      if (arr[i].Nombre === item.Nombre && arr[i].Tienda === item.Tienda &&
         arr[i].Departamento === item.Departamento && !Eliminado) {
          Eliminado=true
      }else{
        nuevo.push(arr[i])
      }
    }
    return nuevo
  }
  */

  removeItemFromArr(arr: Producto[], item: Producto){
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

}