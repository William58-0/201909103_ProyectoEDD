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
  mostrarMensajeError = false
  mostrarMensaje = false
  mensajeError = ''

  constructor(private DatosService: DatosService) {
    this.DatosService.cargartiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      console.log(dataList)
      console.log(this.Tiendas[0].Nombre)
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo cargar la lista de tiendas'
    })
  }

  ngOnInit(): void {
  }

  //////////////////////////////de internet
 handleFileSelect(evt) {
    var files = evt.target.files; // FileList object

    // Loop through the FileList and render image files as thumbnails.
    for (var i = 0, f; f = files[i]; i++) {

      // Only process image files.
      if (!f.type.match('image.*')) {
        continue;
      }

      var reader = new FileReader();

      // Closure to capture the file information.
      reader.onload = (function(theFile) {
        return function(e) {
          // Render thumbnail.
          var span = document.createElement('span');
          span.innerHTML = ['<img class="thumb" src="', e.target.result,
                            '" title="', escape(theFile.name), '"/>'].join('');
          document.getElementById('list').insertBefore(span, null);
        };
      })(f);

      // Read in the image file as a data URL.
      reader.readAsDataURL(f);
      console.log(f)
    }
  }

  
  //////////////////////////////

  cargartiendas(){
    this.DatosService.cargartiendas().subscribe((dataList: any) => {
      this.Tiendas = dataList.Tiendas
      console.log(dataList)
    }, (err) => {
      this.mostrarMensajeError = true
      this.mensajeError = 'No se pudo cargar la lista de tiendas'
    })
  }

  desactivarMensaje() {
    this.mostrarMensaje = false
    this.mostrarMensajeError = false
  }

  mensaje(carnet) {
    console.log("carnet " + carnet)
  }

}

