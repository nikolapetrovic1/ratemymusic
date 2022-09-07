import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserRoutingModule } from './user-routing.module';
import { UserProfileComponent } from './components/user-profile/user-profile.component';
import { AllUsersComponent } from './components/all-users/all-users.component';
import { SharedModule } from '../shared/shared.module';


@NgModule({
  declarations: [
    UserProfileComponent,
    AllUsersComponent
  ],
  imports: [
    CommonModule,
    UserRoutingModule,
    SharedModule
  ]
})
export class UserModule { }
