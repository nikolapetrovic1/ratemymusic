import { HttpHeaders, HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RatingService {


  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });


  constructor(private http: HttpClient) {
   }

   rate(rating:any,type:any): Observable<any> {
    return this.http.post<any>(`http://localhost:3000/rating/${type}`, rating, {
      headers: this.headers,
      responseType: 'json',
    });
  }
  getUserRating(id: string,type:any):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/rating/user/${type}/${id}`,
    );
  }

  getAllSongRatings(id: string,type:any):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/rating/${type}/${id}`,
    );
  }
  findByUser(id: any):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/rating/user/${id}`,
    );
  }
}
