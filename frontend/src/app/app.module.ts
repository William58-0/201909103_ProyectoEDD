import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { InicioComponent } from './components/inicio/inicio.component';
import { TiendasComponent } from './components/Tiendas/Tiendas.component';
import { CarritoComponent } from './components/Carrito/Carrito.component';
import { AdministrarComponent } from './components/Administrar/Administrar.component';

@NgModule({
  declarations: [
    AppComponent,
    InicioComponent,
    TiendasComponent,
    CarritoComponent,
    AdministrarComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
