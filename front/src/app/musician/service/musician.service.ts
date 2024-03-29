import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MusicianService {


  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });


  constructor(private http: HttpClient) {
   }

   search(query: any): Observable<any> {

    let queryParams = {};

    queryParams = {
      headers: this.headers,
      observe: 'response',
      params: new HttpParams()
        .set('query', query)
    };
    return this.http.get<HttpResponse<any>>(
      'http://localhost:3000/musician/search',
      queryParams
    );
  }

  getById(id:any): Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/musician/${id}`,
    );
  }
  submit(value: any) {
    return this.http.post<any>('http://localhost:3000/musician/', value, {
      headers: this.headers,
      responseType: 'json',
    });
  }

}
