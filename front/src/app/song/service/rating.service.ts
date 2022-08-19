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

   rate(rating:any): Observable<any> {
    return this.http.post<any>('http://localhost:3000/rating/song', rating, {
      headers: this.headers,
      responseType: 'json',
    });
  }
  getUserRating(songId: string):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/rating/user/song/${songId}`,
    );
  }

  getAllSongRatings(songId: string):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/rating/song/${songId}`,
    );
  }
}
