import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/auth/service/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  loggedIn!: any;
  constructor(private authService:AuthService,private router:Router) { }

  ngOnInit(): void {
    this.isLoggedIn();
  }

  logout(){
    this.authService.logout();
    this.router.navigate(['/login'])
    this.isLoggedIn();
  }
  isLoggedIn(){
    this.loggedIn = this.authService.isLoggedIn()
  }
  navigate(location:string){
    this.router.navigate([`/${location}`])
  }
}
