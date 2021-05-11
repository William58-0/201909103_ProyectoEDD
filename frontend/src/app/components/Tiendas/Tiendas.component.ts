import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormControl } from '@angular/forms';
import { DatosService } from "../../services/Datos/Datos.service";
import { Tienda } from "../../models/Tienda/Tienda";
import { Producto } from "../../models/Producto/Producto";
import { Usuario } from "../../models/Usuario/Usuario";

@Component({
  selector: 'app-Tiendas',
  templateUrl: './Tiendas.component.html',
  styleUrls: ['./Tiendas.component.css']
})
export class TiendasComponent implements OnInit {

  Tienda: string;
  Departamento: string;
  Tiendas: Tienda[] = []
  Productos: Producto[] = []
  Estado: string;
  Arbol: string;
  Comprados: Producto[] = []
  Usuario: Usuario;
  NuevoDpi: number;
  NuevoNombre: string;
  NuevoPassword: string;
  NuevoCorreo: string;
  Contrasenia: string;
  Nombre: string;
  Id: string;
  NuevoMensaje: string

  //Para los comentarios
  Anterior = {}
  Ruta = ""
  Comentarios = [];
  Totales = this.Comentarios
  Previo: string;


  pilaComentarios = []
  pilaRespuestas = []

  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {
    this.DatosService.GetTiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      this.Totales = this.Comentarios
      console.log(dataList)
      console.log(this.Tiendas[0].Nombre)
      this.Estado = "Tiendas"
      this.DatosService.GetComentarios().subscribe((dataList1: any) => {
        this.CrearComentarios(dataList1.Comentarios)
        console.log(dataList1)
      }, (err) => {
        console.log("no se pudo")
      })
    }, (err) => {
      console.log("no hay tiendas")
    })
    
  }

  ngOnInit(): void {
    this.GetUsuario(this.route.snapshot.paramMap.get('Dpi'))
  }

  //Id: id, Ruta: Ruta, Usuario: String(this.Usuario.Dpi), Nombre: this.Nombre, Mensaje: Mensaje, Respuesta: new FormControl(''),

  GetUsuario(Dpi) {
    var Busqueda = {
      Dpi: Dpi
    }
    this.DatosService.GetUsuario(Busqueda).subscribe(Usuario => {
      this.Usuario = Usuario;
      this.Nombre = Usuario.Nombre
    },
      error => {
        console.log(error);
      });
  }

  Comprar(Producto: Producto) {
    this.DatosService.Comprar(Producto).subscribe(() => {
      if (Producto.Cantidad > 0) {
        Producto.Cantidad--
      }
    }, (err) => {
      console.log("Ocurrio un error")
    })
  }

  Regresar() {
    this.pilaComentarios = []
    this.pilaRespuestas = []
    if(this.Previo=="Tiendas"){
      this.Estado = "Tiendas"
      this.Previo=""
    }else{
      this.Estado="Productos"
      this.Previo="Tiendas"
    }
  }

  INV(Tienda: Tienda) {
    this.Tienda = Tienda.Nombre
    this.Departamento = Tienda.Departamento
    var Busqueda = {
      Tienda: Tienda.Nombre,
      Departamento: Tienda.Departamento,
      Calificacion: Tienda.Calificacion
    }
    this.Arbol = Tienda.Nombre + "---" + Tienda.Departamento + "---" + Tienda.Calificacion + ".png"
    this.DatosService.GetInventario(Busqueda).subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      this.Estado = "Productos"
      this.Previo="Tiendas"
    }, (err) => {
      console.log("No se pudo cargar inventario")
    })
  }

  AComentario(Tienda: Tienda) {
    if(this.Estado=="Tiendas"){
      this.Previo="Tiendas"
    }else{
      this.Previo="Productos"
    }
    this.Estado = "Comentarios"
    this.Id = Tienda.Nombre + "&&&" + Tienda.Calificacion + "&&&" + Tienda.Departamento
    this.pilaComentarios.push({
      Id: this.Id, Ruta: "", Usuario: "", Nombre: "", Mensaje: ""
    })
    this.Filtrar()
  }

  AComentarioP(Producto: Producto) {
    if(this.Estado=="Tiendas"){
      this.Previo="Tiendas"
    }else{
      this.Previo="Productos"
    }
    this.Estado = "Comentarios"
    this.Id = String(Producto.Codigo)
    this.pilaComentarios.push({
      Id: this.Id, Ruta: "", Usuario: "", Nombre: "", Mensaje: ""
    })
    this.Filtrar()
  }

  CrearComentarios(Comments: any) {
    var nuevoo = []
    for (let i = 0; i < Comments.length; i++) {
      var nuevo = {
        Id: Comments[i].Id,
        Ruta: Comments[i].Ruta,
        Usuario: String(Comments[i].Usuario),
        Nombre: Comments[i].Nombre,
        Mensaje: Comments[i].Mensaje,
        Respuesta: new FormControl('')
      }
      nuevoo.push(nuevo)
    }
    this.Totales = nuevoo
  }

  TurnBack() {
    this.Comentarios = this.pilaRespuestas.pop()
    this.pilaComentarios.pop()
    this.Anterior = this.pilaComentarios[this.pilaComentarios.length - 1]
    this.Filtrar()
  }

  //Para comentar sobre la tienda
  NuevoComentario(Mensaje: string) {
    var Ruta = this.HacerCadenaRuta();
    var d = new Date();
    var id = String(d.getHours()) + String(d.getMinutes()) +
      String(d.getFullYear()) + String(d.getMilliseconds())
    var Comment = {
      Id: id,
      Ruta: Ruta,
      Usuario: Number(this.Usuario.Dpi),
      Nombre: this.Nombre,
      Mensaje: Mensaje
    }
    this.DatosService.SendComentario(Comment).subscribe(() => {
      this.Totales.push({
        Id: id,
        Ruta: Ruta,
        Usuario: String(this.Usuario.Dpi),
        Nombre: this.Nombre,
        Mensaje: Mensaje,
        Respuesta: new FormControl('')
      });
      this.Comentarios.push({
        Id: id, Ruta: Ruta,
        Usuario: String(this.Usuario.Dpi),
        Nombre: this.Nombre, Mensaje: Mensaje,
        Respuesta: new FormControl('')
      });
    }, (err) => {
      console.log("Ocurrio un error")
    })
  }

  //Para responder un comentario existente
  Responder(Comentario: any, Mensaje: string) {
    var Ruta = this.HacerCadenaRuta() + "--->" + String(Comentario.Id)
    var d = new Date();
    var id = String(d.getHours()) + String(d.getMinutes()) +
      String(d.getFullYear()) + String(d.getMilliseconds())
    var Comment = {
      Id: id,
      Ruta: Ruta,
      Usuario: Number(this.Usuario.Dpi),
      Nombre: this.Nombre,
      Mensaje: Mensaje
    }
    this.DatosService.SendComentario(Comment).subscribe(() => {
      this.Totales.push({
        Id: id,
        Ruta: Ruta,
        Usuario: String(this.Usuario.Dpi),
        Nombre: this.Nombre,
        Mensaje: Mensaje,
        Respuesta: new FormControl(''),
      })
      if (this.Anterior == Comentario) {
        this.Comentarios.push({
          Id: id,
          Ruta: Ruta,
          Usuario: String(this.Usuario.Dpi),
          Nombre: this.Nombre,
          Mensaje: Mensaje,
          Respuesta: new FormControl(''),
        })
      }
    }, (err) => {
      console.log("Ocurrio un error")
    })
  }


  SendComentario(Comment: any) {
    var user = Number(Comment.Usuario)
    var Comentario = {
      Id: Comment.Id,
      Ruta: Comment.Ruta,
      Usuario: user,
      Nombre: Comment.Nombre,
      Mensaje: Comment.Mensaje
    }
    this.DatosService.SendComentario(Comentario).subscribe(() => {
      alert("si se pudo")
    }, (err) => {
      console.log("Ocurrio un error")
    })
  }


  VerRespuestas(Anterior, Anteriores) {
    this.Anterior = Anterior
    this.pilaComentarios.push(Anterior)
    this.pilaRespuestas.push(Anteriores)
    this.Filtrar()
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
    return nuevo
  }

  Filtrar() {
    var Mostrar = []
    if (this.Totales.length == 1) {
      for (let i = 0; i < this.Totales.length; i++) {
        if (this.Id == this.Totales[i].Ruta) {
          Mostrar.push(this.Totales[i])
        }
      }
    } else {
      for (let i = 0; i < this.Totales.length; i++) {
        if (this.pilaComentarios[this.pilaComentarios.length - 1].Id ==
          this.Totales[i].Ruta.split("--->")[this.Totales[i].Ruta.split("--->").length - 1]) {
          Mostrar.push(this.Totales[i])
        }
      }
    }
    this.Comentarios = Mostrar
  }

  probar() {
    this.DatosService.GetComentarios().subscribe((dataList: any) => {
      this.CrearComentarios(dataList.Comentarios)
      console.log(dataList)
    }, (err) => {
      console.log("no se pudo")
    })
  }


  removeItemFromArr(arr: Producto[], item: Producto) {
    var i = arr.indexOf(item);
    if (i !== -1) {
      arr.splice(i, 1);
    }
    return arr
  }

  GenerarPedido(Productos: Producto[]) {
    this.DatosService.GenerarPedido(Productos).subscribe((res: any) => {
      this.Comprados = null
    }, (err) => {
      console.log("Error")
    })
  }

  ACuentas() {
    this.NuevoDpi = null
    this.NuevoNombre = ""
    this.NuevoCorreo = ""
    this.NuevoPassword = ""
    this.Previo="Tiendas"
    this.Estado = "Cuentas"
  }

  ACarrito() {
    window.location.href = "/Carrito/" + this.Usuario.Dpi
  }

  Eliminar() {
    if (this.Contrasenia === this.Usuario.Password) {
      alert("Contraseña correcta")
      var Usuario: Usuario = {
        Dpi: this.Usuario.Dpi,
        Nombre: this.Usuario.Nombre,
        Correo: this.Usuario.Correo,
        Password: this.Usuario.Password,
        Cuenta: "Usuario"
      }
      this.DatosService.Eliminar(Usuario).subscribe(() => {
        window.location.href = "/"
      }, (err) => {
        console.log("Ocurrio un error")
      })
    }
  }

  /*
   //Para responder un comentario existente
   Responder(Comentario: any, Mensaje: string) {
     var Ruta = this.HacerCadenaRuta() + "--->" + String(Comentario.Id)
     alert(Ruta)
     alert(Mensaje)
     var d = new Date();
     var id = String(d.getHours()) + String(d.getMinutes()) +
       String(d.getFullYear()) + String(d.getMilliseconds())
     alert(id)
     this.Totales.push({
       Id: id, Ruta: Ruta, Usuario: String(this.Usuario.Dpi), Nombre: this.Nombre, Mensaje: Mensaje, Respuesta: new FormControl(''),
     })
     if (this.Anterior == Comentario) {
       this.Comentarios.push({
         Id: id, Ruta: Ruta, Usuario: String(this.Usuario.Dpi), Nombre: this.Nombre, Mensaje: Mensaje, Respuesta: new FormControl(''),
       })
     }
   }
   */

}

