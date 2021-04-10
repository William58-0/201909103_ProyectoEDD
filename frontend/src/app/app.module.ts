import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TiendasComponent } from './components/Tiendas/Tiendas.component';
import { AdministrarComponent } from './components/Administrar/Administrar.component';
import { LoginComponent } from './components/Login/Login.component';
import { CargarComponent } from './components/Cargar/Cargar.component';
import { CarritoComponent } from './components/Carrito/Carrito.component';

@NgModule({
  declarations: [
    AppComponent,
    TiendasComponent,
    CargarComponent,
    CarritoComponent,
    AdministrarComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
