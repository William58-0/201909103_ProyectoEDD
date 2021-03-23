import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { Tienda } from "../../models/Tienda/Tienda";

@Component({
  selector: 'app-Cargar',
  templateUrl: './Cargar.component.html',
  styleUrls: ['./Cargar.component.css']
})
export class CargarComponent implements OnInit {

  Tiendas: Tienda[]=[]
  FileTiendas:File
  FileInventario:File
  FilePedidos:File

  constructor(private DatosService: DatosService) {

  }

  LoadTiendas(event: any){
    this.FileTiendas =event.target.files[0];
    const reader=new FileReader();
    var data1
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      data1=data
      console.log(data)
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
    var data1
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      data1=data
      console.log(data)
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
    var data1
    reader.onload=(e)=>{
      const data =reader.result!.toString().trim();
      data1=data
      console.log(data)
    }
    reader.readAsText(this.FilePedidos)
    this.DatosService.LoadPedidos(this.FilePedidos).subscribe(() => {
    }, (err) => {
      console.log("no se pudo cargar")
    })
  }

  ngOnInit(): void {
  }


}

