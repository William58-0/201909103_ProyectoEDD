import { ANALYZE_FOR_ENTRY_COMPONENTS, Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Usuario } from "../../models/Usuario/Usuario";

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

  constructor(private DatosService: DatosService) {
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

  //Para validar
  Buscar(Registro, Correo) {

  }
}

