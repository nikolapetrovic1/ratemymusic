import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../service/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {


  loginForm!: FormGroup;
  constructor(private fb: FormBuilder,private authService: AuthService,private router:Router) { 
    this.loginForm = this.fb.group({
      email : [''],
      password: ['']
    })
  }

  ngOnInit(): void {
  }

  onSubmit(){
    this.authService.login(this.loginForm.value).subscribe((res)=> {
      localStorage.setItem('user', JSON.stringify(res));
      this.router.navigate(['/']);
    })

  }
}
