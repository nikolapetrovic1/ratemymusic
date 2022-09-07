import { HttpHeaders, HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CommentService {




  private headers = new HttpHeaders({ 'Content-Type': 'application/json'});


  constructor(private http: HttpClient) {
  }



  loadComments(id: any):Observable<any> {
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/comment/review/${id}`,
    );
  }
  submitComment(comment:any):Observable<any> {
    return this.http.post<any>('http://localhost:3000/comment/', comment, {
      headers: this.headers,
    });
  }
  reportComment(id:any): Observable<any> {
    return this.http.get<any>(`http://localhost:3000/comment/report/${id}`)
  }
  getAllByReport(): Observable<any> {
    return this.http.get<any>(`http://localhost:3000/comment/report/all`)
  }

  deleteComment(id: any): Observable<any> {
    return this.http.get<any>(`http://localhost:3000/comment/delete/${id}`)
  }
  getByUser(id: any) {
    return this.http.get<any>(`http://localhost:3000/comment/user/${id}`)
  }

}
