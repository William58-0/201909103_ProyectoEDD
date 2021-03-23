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
  Estado: string;
  Arbol: string;

  constructor(private DatosService: DatosService) {
    this.DatosService.GetTiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      console.log(dataList)
      console.log(this.Tiendas[0].Nombre)
      this.Estado = "Tiendas"
    }, (err) => {
      console.log("no hay tiendas")
    })
  }

  ngOnInit(): void {
  }

  INV(Tienda: Tienda) {
    var Busqueda: Busqueda = {
      Tienda: Tienda.Nombre,
      Departamento: Tienda.Departamento,
      Calificacion: Tienda.Calificacion
    }
    this.Arbol = Tienda.Nombre + "---" + Tienda.Departamento + "---" + Tienda.Calificacion+".png"
    console.log(Busqueda);
    this.DatosService.GetInventario(Busqueda).subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      this.Estado = "Productos"
    }, (err) => {
console.log("No se pudo cargar inventario")
    })
  }

  Comprar(Producto: Producto){
    this.DatosService.Comprar(Producto).subscribe((res:any)=>{
      if(Producto.Cantidad>0){
        Producto.Cantidad--
      }
    }, (err)=>{
console.log("Ocurrio un error")
    })
  }

  Regresar() {
    this.Estado = "Tiendas"
  }

}

