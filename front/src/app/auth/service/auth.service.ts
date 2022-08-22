import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Token } from '@angular/compiler';
@Injectable({
  providedIn: 'root'
})
export class AuthService {


  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });


  constructor(private http: HttpClient) {
   }


  login(auth: any): Observable<Token> {
    return this.http.post<Token>('http://localhost:3000/auth/login', auth, {
      headers: this.headers,
      responseType: 'json',
    });
  }
  register(register: any): Observable<any> {
    return this.http.post<any>('http://localhost:3000/auth/register', register, {
      headers: this.headers,
      responseType: 'json',
    });
  }
  logout() {
    localStorage.clear();
  }

  isLoggedIn(): boolean {
    if (!localStorage.getItem("user")) {
      return false;
    }
    return true;
  }
  loadUserInfo(id: any):Observable<any>{
    return this.http.get<HttpResponse<any>>(
      `http://localhost:3000/auth/${id}`,
    );
  }
}
