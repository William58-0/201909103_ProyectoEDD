import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";

@Component({
  selector: 'app-Administrar',
  templateUrl: './Administrar.component.html',
  styleUrls: ['./Administrar.component.css']
})
export class AdministrarComponent implements OnInit {

  Fechas: string[] = []
  Calendario: string
  mostrarMensajeError = false
  mostrarMensaje = false
  mensajeError = ''
  Arbol: string;
  Estado: string;

  constructor(private DatosService: DatosService) {
    this.DatosService.GetPedidos().subscribe((dataList: any) => {
      this.Fechas = dataList.Fechas
      console.log(dataList)
      //this.Estado = "Calendarios"
      this.Calendario = this.Fechas[0]
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo cargar la lista de pedidos'
    })
  }

  change(date: string) {
    this.Calendario = date
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

  ToYear(date : string){
    var year:string
    year=date.split("-")[0]
    return year
  }

  ngOnInit(): void {
  }

  Acalendarios() {
    this.Estado = "Calendarios"
  }

  desactivarMensaje() {
    this.mostrarMensaje = false
    this.mostrarMensajeError = false
  }

  mensaje(carnet) {
    console.log("carnet " + carnet)
  }

}

