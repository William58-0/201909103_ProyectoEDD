import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DatosService } from "../../services/Datos/Datos.service";
import { Producto } from "../../models/Producto/Producto";
import { Usuario } from "../../models/Usuario/Usuario";
import { Paso } from "../../models/Paso/Paso";
import { Tienda } from "../../models/Tienda/Tienda";

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
  NVec: number = 0;

  //Para grafos
  Grafo: string = "GrafoInicial";
  NumeroPaso: number;
  TiposGrafo = ["Grafo Inicial", "Pasos", "Recorrido Completo"]
  TipoGrafo = "Grafo Inicial"
  Pasos: Paso[] = []
  NPaso: number = 0
  Pendientes: Producto[] = []
  Recogidos: Producto[] = []
  Recorrido: string = ""
  Distancia: number = 0
  Usuarios: number[] = []

  //Para AVL
  Tiendas: Tienda[] = []


  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {
    this.DatosService.GetFechas().subscribe((dataList: any) => {
      this.Fechas = dataList.Fechas
      this.Calendario = this.Fechas[0]
      this.Estado = "AVL"
      this.Valido = false
    }, (err) => {
      console.log('No se pudo cargar la lista de fechas')
    })
    this.DatosService.GetPedidos().subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      this.Calendario = this.Fechas[0]
      this.Mostrar = this.Filtrar(this.Productos)
    }, (err) => {
      console.log('No se pudieron cargar los pedidos')
    })
    this.DatosService.GetRecorrido().subscribe((dataList: any) => {
      this.Pasos = dataList.Pasos
    }, (err) => {
      console.log("error")
    })
    this.DatosService.GetUsuarios().subscribe((dataList: any) => {
      this.Usuarios = dataList
    }, (err) => {
      console.log("error")
    })
    this.DatosService.GetTiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      console.log(dataList)
      console.log(this.Tiendas[0].Nombre)
    }, (err) => {
      console.log("no hay tiendas")
    })

  }

  ngOnInit(): void {
    this.GetUsuario(this.route.snapshot.paramMap.get('Dpi'))
    for (let i = 0; i < this.Pasos.length; i++) {
      for (let j = 0; j < this.Pasos[i].Pendientes.length; j++) {
        if (this.VerSiExiste(this.Pasos[i].Pendientes[j].Cliente)) {
          this.Pasos[i].Pendientes[j].Cliente = 0
        }
      }
      for (let j = 0; j < this.Pasos[i].Recogidos.length; j++) {
        if (this.VerSiExiste(this.Pasos[i].Recogidos[j].Cliente)) {
          this.Pasos[i].Recogidos[j].Cliente = 0
        }
      }
    }
  }

  GetUsuario(Dpi) {
    var Busqueda = {
      Dpi: Dpi
    }
    this.DatosService.GetUsuario(Busqueda).subscribe(data => {
      this.Usuario = data;
      this.Nombre = data.Nombre;
      this.Valido = false
    },
      error => {
        console.log(error);
      });
  }

  VerSiExiste(Dpi) {
    var respuesta=false
    if (this.Usuarios.includes(Dpi)){
      respuesta=true
      return true
    }
      return respuesta
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

  AMerkle() {
    this.Estado = "Merkle"
  }

  AAVL() {
    this.Estado = "AVL"
  }

  Aarbolcuentas() {
    this.Estado = "ArbolCuentas"
    this.Valido=false;

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

  AvanzarVec() {
    this.NVec++
  }

  RetrocederVec() {
    if (this.NVec - 1 >= 0) {
      this.NVec--
    }
  }

  ToYear(date: string) {
    var year: string
    year = date.split("-")[0]
    return year
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

  Avanzar() {
    if (this.NPaso + 1 < this.Pasos.length) {
      this.NPaso++
      this.Pendientes = this.Pasos[this.NPaso].Pendientes
      this.Recogidos = this.Pasos[this.NPaso].Recogidos
      this.Recorrido = this.Pasos[this.NPaso].Recorrido
      this.Distancia = this.Pasos[this.NPaso].Distancia
      this.Grafo = "Paso" + this.NPaso
    }
  }

  Retroceder() {
    if (this.NPaso - 1 >= 0) {
      this.NPaso--
      this.Pendientes = this.Pasos[this.NPaso].Pendientes
      this.Recogidos = this.Pasos[this.NPaso].Recogidos
      this.Recorrido = this.Pasos[this.NPaso].Recorrido
      this.Distancia = this.Pasos[this.NPaso].Distancia
      this.Grafo = "Paso" + this.NPaso
    }
  }

  CambiarTipoGrafo(grafo: string) {
    if (grafo == "Grafo Inicial") {
      this.Grafo = "GrafoInicial"
    } else if (grafo == "Pasos") {
      this.NPaso = 0
      this.Pendientes = this.Pasos[this.NPaso].Pendientes
      this.Recogidos = this.Pasos[this.NPaso].Recogidos
      this.Recorrido = this.Pasos[this.NPaso].Recorrido
      this.Distancia = this.Pasos[this.NPaso].Distancia
      this.Grafo = "Paso" + this.NPaso
    } else {
      this.NPaso = this.Pasos.length - 1
      this.Pendientes = this.Pasos[this.NPaso].Pendientes
      this.Recogidos = this.Pasos[this.NPaso].Recogidos
      this.Recorrido = this.Pasos[this.NPaso].Recorrido
      this.Distancia = this.Pasos[this.NPaso].Distancia
      this.Grafo = "RecorridoCompleto"
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
    },
      error => {
        console.log(error);
      })
  }



}
