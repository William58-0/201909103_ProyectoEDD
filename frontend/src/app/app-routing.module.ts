import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { InicioComponent } from "./components/inicio/inicio.component";
import { TiendasComponent } from "./components/Tiendas/Tiendas.component";
import { CargarComponent } from './components/Cargar/Cargar.component';
import { CarritoComponent } from './components/Carrito/Carrito.component';
import { ComprarComponent } from './components/Comprar/Comprar.component';

const routes: Routes = [
  {
    path: '',
    component: InicioComponent,
  },
  {
    path: 'Tiendas',
    component: TiendasComponent,
  },
  {
    path: 'Cargar',
    component: CargarComponent,
  },
  {
    path: 'Carrito',
    component: CarritoComponent,
  },
  {
    path: 'Comprar',
    component: ComprarComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
