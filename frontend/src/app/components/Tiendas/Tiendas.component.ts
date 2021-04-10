import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
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

  Tienda:string;
  Departamento: string;
  Tiendas: Tienda[] = []
  Productos: Producto[] = []
  Estado: string;
  Arbol: string;
  Comprados: Producto[] = []
  Usuario:Usuario;
  NuevoDpi:number;
  NuevoNombre:string;
  NuevoPassword:string;
  NuevoCorreo:string;
  Contrasenia:string;

  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {
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
    this.GetUsuario(this.route.snapshot.paramMap.get('Dpi'))
  }

  GetUsuario(Dpi) {
    var Busqueda = {
      Dpi: Dpi
    }
    this.DatosService.GetUsuario(Busqueda).subscribe(Usuario => {
      this.Usuario = Usuario;
      console.log("Datos")
      console.log(Usuario.Dpi)
      //this.Carro="/Carrito/"+this.Usuario.Dpi
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

  INV(Tienda: Tienda) {
    this.Tienda=Tienda.Nombre
    this.Departamento=Tienda.Departamento
    var Busqueda = {
      Tienda: Tienda.Nombre,
      Departamento: Tienda.Departamento,
      Calificacion: Tienda.Calificacion
    }
    this.Arbol = Tienda.Nombre + "---" + Tienda.Departamento + "---" + Tienda.Calificacion + ".png"
    console.log(Busqueda);
    this.DatosService.GetInventario(Busqueda).subscribe((dataList: any) => {
      this.Productos = dataList.Productos
      console.log(dataList)
      this.Estado = "Productos"
    }, (err) => {
      console.log("No se pudo cargar inventario")
    })
  }

  Registrar(){
    var NuevoUsuario:Usuario={
      Dpi:this.NuevoDpi,
      Nombre:this.NuevoNombre,
      Correo:this.NuevoCorreo,
      Password:this.NuevoPassword,
      Cuenta:"Usuario"
    }
    this.DatosService.Registrar(NuevoUsuario).subscribe((res: any) => {

    }, (err) => {
      console.log("Ocurrio un error")
    })

  }

  Regresar() {
    this.Estado = "Tiendas"
  }

  removeItemFromArr(arr: Producto[], item: Producto){
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

  ACuentas(){
    this.NuevoDpi=null
    this.NuevoNombre=""
    this.NuevoCorreo=""
    this.NuevoPassword=""
    this.Estado="Cuentas"
  }

  ACarrito(){
    window.location.href="/Carrito/"+this.Usuario.Dpi
  }

  Eliminar(){
    if(this.Contrasenia===this.Usuario.Password){
      alert("son iguales")
      var Usuario:Usuario={
        Dpi:this.Usuario.Dpi,
        Nombre:this.Usuario.Nombre,
        Correo:this.Usuario.Correo,
        Password:this.Usuario.Password,
        Cuenta:"Usuario"
      }
      this.DatosService.Eliminar(Usuario).subscribe(() => {
        window.location.href="/Cargar/"+Usuario.Dpi
      }, (err) => {
        console.log("Ocurrio un error")
      })
    }
  }

}

