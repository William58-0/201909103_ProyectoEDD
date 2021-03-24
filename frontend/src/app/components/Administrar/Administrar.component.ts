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
  MostrarP: Producto[] = []
  Fechas: string[] = []
  Calendario: string
  Arbol: string;
  Estado: string;
  //Para los pedidos
  Dates: string[] = []
  Departamentos: string[] = []
  Tiendas:string[]=[]
  Fecha: string = "Cualquiera"
  Tienda: string = "Cualquiera"
  Departamento: string = "Cualquiera"

  constructor(private DatosService: DatosService) {
    this.DatosService.GetFechas().subscribe((dataList: any) => {
      this.Fechas = dataList.Fechas
      console.log(dataList)
      this.Calendario = this.Fechas[0]
      this.Estado = "Pedidos"
    }, (err) => {
      console.log('No se pudo cargar la lista de meses')
    })
    this.DatosService.GetPedidos().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      this.Dates = []
      this.Departamentos = []
      this.Tiendas=[]
      this.AgregarFechas(this.Productos)
      this.AgregarDepartamentos(this.Productos)
      this.AgregarTiendas(this.Productos)
      this.Calendario = this.Fechas[0]
      this.Mostrar = this.Filtrar(this.Productos)
      this.MostrarP = this.FiltrarProductos(this.Fecha, this.Tienda, this.Departamento)
    }, (err) => {
      console.log('No se pudo cargar la lista de meses')
    })
  }

  FiltrarProductos(Fecha: string, Tienda: string, Departamento: string) {
    var nuevo: Producto[] = []
    var ignorados: Producto[]=[]
    //Filtrar por fecha
    if (this.Fecha != "Cualquiera" && this.Fecha != "" && this.Fecha != null) {
      for (let i = 0; i < this.Productos.length; i++) {
        if (this.Productos[i].Fecha === Fecha && !nuevo.includes(this.Productos[i]) && !ignorados.includes(this.Productos[i])) {
          nuevo.push(this.Productos[i])
        }else if(!ignorados.includes(this.Productos[i])){
          ignorados.push(this.Productos[i])
        }
      }
    }
    //Filtrar por Tienda
    if (this.Tienda != "Cualquiera" && this.Tienda != "" && this.Tienda != null) {
      for (let i = 0; i < this.Productos.length; i++) {
        if (this.Productos[i].Tienda === Tienda && !nuevo.includes(this.Productos[i]) && !ignorados.includes(this.Productos[i])) {
          nuevo.push(this.Productos[i])
        }else if(!ignorados.includes(this.Productos[i])){
          ignorados.push(this.Productos[i])
        }
      }
    }
    //Filtrar por Departamento
    if (this.Departamento != "Cualquiera" && this.Departamento != "" && this.Departamento != null) {
      for (let i = 0; i < this.Productos.length; i++) {
        if (this.Productos[i].Departamento === Departamento && !nuevo.includes(this.Productos[i]) && !ignorados.includes(this.Productos[i])) {
          nuevo.push(this.Productos[i])
        }else if(!ignorados.includes(this.Productos[i])){
          ignorados.push(this.Productos[i])
        }
      }
    }
    if (Fecha === "Cualquiera" && Tienda === "Cualquiera" && Departamento === "Cualquiera") {
      nuevo = this.Productos
    }
    return nuevo
  }

  changeCalendar(date: string) {
    this.Calendario = date
    this.Mostrar = this.Filtrar(this.Productos)
  }

  changeDate(date: string) {
    this.Fecha = date
    this.MostrarP = this.FiltrarProductos(this.Fecha,this.Tienda,this.Departamento)
  }

  changeTienda(Tienda: string) {
    this.Tienda = Tienda
    this.MostrarP = this.FiltrarProductos(this.Fecha,this.Tienda,this.Departamento)
  }

  changeDep(Dep: string) {
    this.Departamento = Dep
    this.MostrarP = this.FiltrarProductos(this.Fecha,this.Tienda,this.Departamento)
  }

  Apedidos() {
    this.Estado = "Pedidos"
  }

  Acalendarios() {
    this.Estado = "Calendarios"
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

  info(dato) {
    console.log(dato)
  }

  AgregarFechas(Productos: Producto[]) {
    for (let i = 0; i < Productos.length; i++) {
      if (!this.Dates.includes(Productos[i].Fecha)) {
        this.Dates.push(Productos[i].Fecha)
      }
    }
    var j: number
    var aux: string
    var n = this.Dates.length
    for (let i = 1; i < n; i++) {
      j = i
      aux = this.Dates[i]
      while (j > 0 && aux.split("-")[2] < this.Dates[j - 1].split("-")[2]) {
        this.Dates[j] = this.Dates[j - 1]
        j--
      }
      this.Dates[j] = aux
    }
  }

  AgregarDepartamentos(Productos: Producto[]) {
    for (let i = 0; i < Productos.length; i++) {
      if (!this.Departamentos.includes(Productos[i].Departamento)) {
        this.Departamentos.push(Productos[i].Departamento)
      }
    }
  }

  AgregarTiendas(Productos: Producto[]) {
    for (let i = 0; i < Productos.length; i++) {
      if (!this.Tiendas.includes(Productos[i].Tienda)) {
        this.Tiendas.push(Productos[i].Tienda)
      }
    }
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
