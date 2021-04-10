import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TiendasComponent } from "./components/Tiendas/Tiendas.component";
import { CargarComponent } from './components/Cargar/Cargar.component';
import { CarritoComponent } from './components/Carrito/Carrito.component';
import { AdministrarComponent } from './components/Administrar/Administrar.component';
import { LoginComponent } from './components/Login/Login.component';

const routes: Routes = [
  {path: '',component: LoginComponent,},
  {path: 'Tiendas/:Dpi',component: TiendasComponent,},
  {path: 'Cargar/:Dpi',component: CargarComponent,},
  {path: 'Carrito/:Dpi',component: CarritoComponent,},
  {path: 'Administrar/:Dpi',component: AdministrarComponent,}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
