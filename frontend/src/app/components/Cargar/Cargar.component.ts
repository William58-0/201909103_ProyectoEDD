import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DatosService } from "../../services/Datos/Datos.service";
import { Usuario } from "../../models/Usuario/Usuario";

@Component({
  selector: 'app-Cargar',
  templateUrl: './Cargar.component.html',
  styleUrls: ['./Cargar.component.css']
})
export class CargarComponent implements OnInit {

  FileTiendas:File
  FileInventario:File
  FilePedidos:File
  FileUsuarios:File
  FileGrafo:File
  Usuario:Usuario;
  Nombre:string;
  Estado:string
  NuevoUsuario:Usuario

  constructor(private DatosService: DatosService,
    private route: ActivatedRoute,
    private router: Router) {   
      this.Estado="Cargar"
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
      this.Nombre=data.Nombre
    },
      error => {
        console.log(error);
      });
  }

  LoadTiendas(event: any){
    this.FileTiendas =event.target.files[0];
    const reader=new FileReader();
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      //console.log(data)
    }
    reader.readAsText(this.FileTiendas)
    this.DatosService.LoadTiendas(this.FileTiendas).subscribe(() => {
    }, (err) => {
      console.log("no se pudo cargar")
    })
  }

  LoadInventario(event: any){
    this.FileInventario =event.target.files[0];
    const reader=new FileReader();
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      //console.log(data)
    }
    reader.readAsText(this.FileInventario)
    this.DatosService.LoadInventario(this.FileInventario).subscribe(() => {
    }, (err) => {
      console.log("no se pudo cargar")
    })
  }

  LoadPedidos(event: any){
    this.FilePedidos =event.target.files[0];
    const reader=new FileReader();
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      //console.log(data)
    }
    reader.readAsText(this.FilePedidos)
    this.DatosService.LoadFechas(this.FilePedidos).subscribe(() => {
    }, (err) => {
      console.log("no se pudo cargar")
    })
  }

  LoadUsuarios(event: any){
    this.FileUsuarios =event.target.files[0];
    const reader=new FileReader();
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      //console.log(data)
    }
    reader.readAsText(this.FileUsuarios)
    this.DatosService.LoadUsuarios(this.FileUsuarios).subscribe(() => {
    }, (err) => {
      console.log("no se pudo cargar")
    })
  }

  LoadGrafo(event: any){
    this.FileGrafo =event.target.files[0];
    const reader=new FileReader();
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      //console.log(data)
    }
    reader.readAsText(this.FileGrafo)
    this.DatosService.LoadGrafo(this.FileGrafo).subscribe(() => {
    }, (err) => {
      console.log("no se pudo cargar")
    })
  }

  ACuentas(){
    this.NuevoUsuario.Dpi=null
    this.NuevoUsuario.Nombre=null
    this.NuevoUsuario.Correo=null
    this.NuevoUsuario.Password=null
    this.NuevoUsuario.Cuenta="Admin"
    this.Estado="Cuentas"
  }

  AReportes(){
    window.location.href="/Administrar/"+this.Usuario.Dpi
  }

}

