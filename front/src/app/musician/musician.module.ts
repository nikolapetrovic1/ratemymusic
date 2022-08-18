import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MusicianRoutingModule } from './musician-routing.module';
import { AllMusiciansComponent } from './components/all-musicians/all-musicians.component';
import { SharedModule } from '../shared/shared.module';
import { MusicianCardComponent } from './components/musician-card/musician-card.component';
import { MusicianProfileComponent } from './components/musician-profile/musician-profile.component';


@NgModule({
  declarations: [
    AllMusiciansComponent,
    MusicianCardComponent,
    MusicianProfileComponent
  ],
  imports: [
    CommonModule,
    MusicianRoutingModule,
    SharedModule
  ]
})
export class MusicianModule { }
