import { Injectable } from "@angular/core";
import { Router, CanActivate } from "@angular/router";
import { AuthService } from "../../service/auth.service";

@Injectable({
  providedIn: "root",
})
export class LoginGuard implements CanActivate {
  constructor(public auth: AuthService, public router: Router) {}

  canActivate(): boolean {
    if (this.auth.isLoggedIn()) {
      this.router.navigate(["/song"]);
      return false;
    }
    return true;
  }
}
