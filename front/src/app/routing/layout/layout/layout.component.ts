import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { HeaderComponent } from 'src/app/shared/components/header/header.component';
@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent implements OnInit {
  role: any;

  constructor() { this.role = "";}

  ngOnInit(): void {
  }

  checkRole() {
    const item = localStorage.getItem("user");

    if (item) {
      const jwt: JwtHelperService = new JwtHelperService();
      this.role = jwt.decodeToken(item).role[0].authority;
      console.log(this.role)
    }
  }
}
