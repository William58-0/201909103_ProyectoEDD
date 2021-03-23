import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class DatosService {
  
  constructor(private http: HttpClient) { }

  LoadTiendas(data):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'LoadTiendas', data, httpOptions);
  }

  GetTiendas():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'GetTiendas', httpOptions);
  }

  LoadInventario(data):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'LoadInventario', data, httpOptions);
  }

  GetInventario(Busqueda):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'GetInventario', Busqueda, httpOptions);
  }

  LoadPedidos(data):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'LoadPedidos', data, httpOptions);
  }

  GetPedidos():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'GetPedidos', httpOptions);
  }

  Comprar(Producto):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'Comprar', Producto, httpOptions);
  }

  Devolver(Producto):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'Devolver', Producto, httpOptions);
  }

  CargarCarro():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'CargarCarro', httpOptions);
  }




  mandarInventarios(json: string):Observable<any>{
    return this.http.post<any>('http://localhost:3000/cargarInventarios',json)
  }

}
