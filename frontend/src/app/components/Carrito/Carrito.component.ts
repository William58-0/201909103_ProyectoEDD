import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";
import { Usuario } from "../../models/Usuario/Usuario";

@Component({
  selector: 'app-Carrito',
  templateUrl: './Carrito.component.html',
  styleUrls: ['./Carrito.component.css']
})
export class CarritoComponent implements OnInit {

  Productos: Producto[] = []
  Usuario:Usuario;

  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {
    this.DatosService.CargarCarro().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
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

  ATiendas(){
    window.location.href="/Tiendas/"+this.Usuario.Dpi
  }

}