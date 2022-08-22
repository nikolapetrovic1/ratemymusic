import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SongService {


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
      'http://localhost:3000/song/search',
      queryParams
    );
  }
  get(id:string): Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/song/${id}`,
    );
  }
  getByMusician(id:string): Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/song/musician/${id}`,
    );
  }
  submit(value: any) {
    return this.http.post<any>('http://localhost:3000/song/', value, {
      headers: this.headers,
      responseType: 'json',
    });
  }
}
