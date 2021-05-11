
import { ANALYZE_FOR_ENTRY_COMPONENTS, Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Usuario } from "../../models/Usuario/Usuario";
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-Login',
  templateUrl: './Login.component.html',
  styleUrls: ['./Login.component.css']
})

export class LoginComponent implements OnInit {
  Estado: string
  public static Usuario: Usuario;
  Anterior = {}
  Ruta = ""

  Comentarios = [
    {
      Id: "Primero", Ruta: "Tienda", Usuario: "nuevo", Nombre: "Uno", Mensaje: "Sistemas operativos 2"
    },
    {
      Id: "Segundo", Ruta: "Tienda", Usuario: "nuevo", Nombre: "dos", Mensaje: "Sistemas operativos 2"
    },
    {
      Id: "Tercero", Ruta: "Tienda--->Primero", Usuario: "nuevo", Nombre: "tres", Mensaje: "Sistemas operativos 2"
    },
    {
      Id: "Cuarto", Ruta: "Tienda--->Segundo", Usuario: "nuevo", Nombre: "cuatro", Mensaje: "Sistemas operativos 2"
    },
    {
      Id: "Quinto", Ruta: "Tienda--->Segundo--->Cuarto", Usuario: "nuevo", Nombre: "quinto", Mensaje: "Sistemas operativos 2"
    },
  ];
  Totales = this.Comentarios


  pilaComentarios = []
  pilaRespuestas = []


  constructor(private DatosService: DatosService,
    private sanitizer: DomSanitizer) {
    this.Estado = "Login"
    this.pilaComentarios.push({
      Id: "Tienda", Ruta: "", Usuario: "", Nombre: "", Mensaje: ""
    })
    this.Totales = this.Comentarios
    this.Filtrar()
  }

  ngOnInit(): void {
    this.Filtrar()
  }

  Regresar() {
    this.Comentarios = this.pilaRespuestas.pop()
    this.pilaComentarios.pop()
    this.Anterior = this.pilaComentarios[this.pilaComentarios.length - 1]
    this.Filtrar()
  }

  VerRespuestas(Anterior, Anteriores) {
    this.Anterior = Anterior
    this.pilaComentarios.push(Anterior)
    this.pilaRespuestas.push(Anteriores)
    this.Filtrar()
    //this.Comentarios = Nuevos
  }

  HacerCadenaRuta() {
    var nuevo = ""
    for (let i = 0; i < this.pilaComentarios.length; i++) {
      if (i < this.pilaComentarios.length - 1) {
        nuevo += this.pilaComentarios[i].Id + "--->"
      } else {
        nuevo += this.pilaComentarios[i].Id
      }
    }
    alert(nuevo)
  }

  Filtrar() {
    var Mostrar = []
    if (this.Totales.length == 1) {
      for (let i = 0; i < this.Totales.length; i++) {
        if ("Tienda" == this.Totales[i].Ruta) {
          Mostrar.push(this.Totales[i])
        }
      }
    } else {
      for (let i = 0; i < this.Totales.length; i++) {
        if (this.pilaComentarios[this.pilaComentarios.length - 1].Id == 
          this.Totales[i].Ruta.split("--->")[this.Totales[i].Ruta.split("--->").length-1]) {
          Mostrar.push(this.Totales[i])
        }
      }
    }
    this.Comentarios = Mostrar
  }

  Probar() {
    console.log(this.pilaComentarios)
    var d = new Date();
    alert(d.getHours());
    alert(d.getMinutes());
    alert(d.getFullYear());
    alert(d.getMilliseconds());
  }

}

