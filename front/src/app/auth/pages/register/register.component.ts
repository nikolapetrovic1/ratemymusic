import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../service/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

  registerForm!: FormGroup;
  constructor(private fb: FormBuilder,private authService: AuthService,private router:Router) { 
    this.registerForm = this.fb.group({
      email : [''],
      password: ['']
    })
  }

  ngOnInit(): void {
  }

  onSubmit(){
    this.authService.register(this.registerForm.value).subscribe((res)=> {
      this.router.navigate(['/login'])
    })

  }

}
