import { Component, OnInit } from '@angular/core';
import { DatosService } from "../../services/Datos/Datos.service";
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-Comprar',
  templateUrl: './Comprar.component.html',
  styleUrls: ['./Comprar.component.css']
})
export class ComprarComponent implements OnInit {

  mostrarMensajeError=false
  mostrarMensaje=false
  mensajeError = ''

  constructor(private DatosService:DatosService) {

  }

  ngOnInit(): void {
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  mensaje(carnet){
    console.log("carnet "+carnet)
  }

}
