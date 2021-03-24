import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";

@Component({
  selector: 'app-Administrar',
  templateUrl: './Administrar.component.html',
  styleUrls: ['./Administrar.component.css']
})


export class AdministrarComponent implements OnInit {

  Productos: Producto[] = []
  Mostrar: Producto[] = []
  Fechas: string[] = []
  Calendario: string
  Arbol: string;
  Estado: string;


  constructor(private DatosService: DatosService) {
    this.DatosService.GetFechas().subscribe((dataList: any) => {
      this.Fechas = dataList.Fechas
      console.log(dataList)
      this.Calendario = this.Fechas[0]
      this.Estado = "Calendarios"
    }, (err) => {
      console.log('No se pudo cargar la lista de fechas')
    })
    this.DatosService.GetPedidos().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      this.Calendario = this.Fechas[0]
      this.Mostrar = this.Filtrar(this.Productos)
    }, (err) => {
      console.log('No se pudieron cargar los pedidos')
    })
  }

  changeCalendar(date: string) {
    this.Calendario = date
    this.Mostrar = this.Filtrar(this.Productos)
  }

  Apedidos() {
    this.Estado = "Pedidos"
  }

  Acalendarios() {
    this.Estado = "Calendarios"
  }

  Aarbol() {
    this.Estado = "Arbol"
  }

  ToMes(date: string) {
    var month: string;
    month = date.split("-")[1]
    switch (month) {
      case "01":
        return "Enero"
      case "02":
        return "Febrero"
      case "03":
        return "Marzo"
      case "04":
        return "Abril"
      case "05":
        return "Mayo"
      case "06":
        return "Junio"
      case "07":
        return "Julio"
      case "08":
        return "Agosto"
      case "09":
        return "Septiembre"
      case "10":
        return "Octubre"
      case "11":
        return "Noviembre"
      case "12":
        return "Diciembre"
    }
    return date
  }

  ToYear(date: string) {
    var year: string
    year = date.split("-")[0]
    return year
  }

  ngOnInit(): void {
  }

  Filtrar(Productos: Producto[]) {
    var nuevo: Producto[] = []
    for (let i = 0; i < Productos.length; i++) {
      if (Productos[i].Fecha.split("-")[2] + "-" + Productos[i].Fecha.split("-")[1] === this.Calendario) {
        nuevo.push(Productos[i])
      }
    }
    return nuevo
  }

}
