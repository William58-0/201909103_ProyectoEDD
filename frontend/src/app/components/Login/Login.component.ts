
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
  Dpi: number;
  Password: string;
  NuevoDpi:number;
  NuevoNombre:string;
  NuevoPassword:string;
  NuevoCorreo:string;
  Comentarios1 = [
    { Autor: 285, Texto: "Sistemas operativos 2", Subcomentarios: [], },
    { Autor: 975, Texto: "Redes de computadoras 2", Subcomentarios: [], }
    ];
  Comentarios = [
  { Autor: 285, Texto: "Sistemas operativos 2", Subcomentarios: this.Comentarios1, },
  { Autor: 975, Texto: "Redes de computadoras 2", Subcomentarios: 4, }
  ];

  constructor(private DatosService: DatosService,
    private sanitizer:DomSanitizer) {
    this.Estado = "Login"
  }

  IniciarSesion() {
    var user:Usuario = {
      Dpi: this.Dpi,
      Nombre:"",
      Correo:"",
      Password: this.Password,
      Cuenta:"",
    }
    console.log(user);
    this.DatosService.IniciarSesion(user).subscribe((Usuario: Usuario) => {
      console.log(Usuario)
      if(Usuario.Dpi===0 &&
        Usuario.Nombre==="" &&
        Usuario.Correo==="" &&
        Usuario.Password==="" &&
        Usuario.Cuenta===""){
          alert("Este usuario no existe")
      }else{
        alert("Bienvenido "+Usuario.Nombre)
        if(Usuario.Cuenta==="Admin"){
          window.location.href="/Cargar/"+Usuario.Dpi
        }else if(Usuario.Cuenta==="Usuario"){
          window.location.href="/Tiendas/"+Usuario.Dpi
        }
      }
    }, (err) => {
      console.log("No se pudo cargar usuario")
    })
  }

  ngOnInit(): void {

  }

  Registrar(){
    var NuevoUsuario:Usuario={
      Dpi:this.NuevoDpi,
      Nombre:this.NuevoNombre,
      Correo:this.NuevoCorreo,
      Password:this.NuevoPassword,
      Cuenta:"Usuario"
    }
    this.DatosService.Registrar(NuevoUsuario).subscribe(() => {

    }, (err) => {
      console.log("Ocurrio un error")
    })

  }

  Regresar() {
    this.Estado = "Login"
  }

  ARegistrar(){
    this.NuevoDpi=null
    this.NuevoNombre=""
    this.NuevoCorreo=""
    this.NuevoPassword=""
    this.Estado="Registrar"
  }

}
