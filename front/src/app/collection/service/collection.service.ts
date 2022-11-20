import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CollectionService {


  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });


  constructor(private http: HttpClient) {}

  favorite(favorite:any): Observable<any> {
    return this.http.post<any>(`http://localhost:3000/collection/favorite`, favorite ,{
      headers: this.headers,
      responseType: 'json',
    });
  }

  unFavorite(id:any): Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/collection/unfavorite/${id}`,
    );
  }

  getFavoriteUserKindId(kind:string,kindId:string):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/collection/user/${kind}/${kindId}`,
    );
  }

}
