import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginGuard } from './guards/login/login.guard';
import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';

const routes: Routes = [  
  {
  path: "login",
  pathMatch: "full",
  component: LoginComponent,
  canActivate: [LoginGuard],
},
{
  path: "register",
  pathMatch: "full",
  component: RegisterComponent,
  canActivate: [LoginGuard],
},
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
