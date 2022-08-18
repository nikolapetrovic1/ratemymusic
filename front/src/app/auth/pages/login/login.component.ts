import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { AuthService } from '../../service/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {


  loginForm!: FormGroup;
  constructor(private fb: FormBuilder,private authService: AuthService) { 
    this.loginForm = this.fb.group({
      email : [''],
      password: ['']
    })
  }

  ngOnInit(): void {
  }

  onSubmit(){
    this.authService.login(this.loginForm.value).subscribe((res)=> {
      alert(res);
    })

  }
}
