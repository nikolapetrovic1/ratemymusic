import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { Token } from '../models/Token';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const item = localStorage.getItem('user');
    if (item) {
      const decodedItem: Token = JSON.parse(item);
      const cloned = req.clone({
        headers: req.headers.set(
          'Authorization',
          `Bearer ${decodedItem.accessToken}`
        ),
      });
      return next.handle(cloned);
    } else {
      return next.handle(req);
    }
  }
}
