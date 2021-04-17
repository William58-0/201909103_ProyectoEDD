import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";
import { Usuario } from "../../models/Usuario/Usuario";
import { Paso } from "../../models/Paso/Paso";

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
  Usuario: Usuario;
  Nombre: string;
  Clave: string;
  Valido: boolean;
  //Para matriz linealizada
  NVec:number=0;

  //Para grafos
  Grafo: string;
  NumeroPaso: number;
  TiposGrafo = ["Grafo Inicial", "Pasos", "Recorrido Completo"]
  TipoGrafo = "Grafo Inicial"
  Pasos: Paso[]=[]
  NPaso:number=0
  Pendientes:Producto[]=[]
  Recogidos:Producto[]=[]
  Recorrido:string=""
  Distancia:number=0


  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {
    this.DatosService.GetFechas().subscribe((dataList: any) => {
      this.Fechas = dataList.Fechas
      console.log(dataList)
      this.Calendario = this.Fechas[0]
      this.Estado = "Grafos"
      this.Valido = false
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
    this.DatosService.GetRecorrido().subscribe((dataList: any) => {
      this.Pasos = dataList.Pasos
      console.log(dataList)
      //console.log(this.Productos[0])>
    }, (err) => {
      console.log("error")
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

  Agrafos() {
    this.Estado = "Grafos"
    this.Grafo = "GrafoInicial"
  }

  Aarbol() {
    this.Estado = "Arbol"
  }

  Aarbolcuentas() {
    this.Estado = "ArbolCuentas"
  }

  Amatrizlinealizada() {
    this.Estado = "MatrizLinealizada"
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

  AvanzarVec(){
      this.NVec++

  }

  RetrocederVec(){
    if (this.NVec-1>=0){
      this.NVec--
    }
  }

  ToYear(date: string) {
    var year: string
    year = date.split("-")[0]
    return year
  }

  ngOnInit(): void {
    this.GetUsuario(this.route.snapshot.paramMap.get('Dpi'))
  }

  GetUsuario(Dpi) {
    var Busqueda = {
      Dpi: Dpi
    }
    this.DatosService.GetUsuario(Busqueda).subscribe(data => {
      this.Usuario = data;
      console.log(data)
      this.Nombre = data.Nombre;
      this.Valido = false
    },
      error => {
        console.log(error);
      });
  }

  VerificarLlave() {
    if (this.Clave == "tercerafaseeddfechadeentregaelsa") {
      this.Valido = true;
    } else {
      this.Valido = false;
      alert("Clave Incorrecta")
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

  CambiarTipoGrafo(grafo: string) {
    if (grafo == "Grafo Inicial") {
      this.Grafo = "GrafoInicial"
    }else if(grafo == "Pasos"){
      this.Grafo="Paso"+this.NPaso
    }else{
      this.Grafo="RecorridoCompleto"
    }

  }

  ADatosCuentas() {
    window.location.href = "/Cargar/" + this.Usuario.Dpi
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

  Prueba(){
    this.DatosService.GetRecorrido().subscribe((dataList: any) => {
      this.Pasos = dataList.Pasos
      console.log(dataList)
      //console.log(this.Productos[0])>
    }, (err) => {
      console.log("error")
    })
  }


}
