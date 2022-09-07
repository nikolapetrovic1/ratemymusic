import { HttpHeaders, HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {



  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });


  constructor(private http: HttpClient) {
  }

   getAllReview(id:any,type:string):Observable<any>{
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/review/${type}/${id}`,
    );
   }
   getUserReview(id:any,type:string):Observable<any>{
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/review/user/${type}/${id}`,
    );
   }
   submitReview(review:any):Observable<any>{
    return this.http.post<any>('http://localhost:3000/review/', review, {
      headers: this.headers,
    });
   }
   findByUser(id: any):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/review/user/${id}`,
    );
  }

}
