import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SongRoutingModule } from './song-routing.module';
import { AllSongsComponent } from './components/all-songs/all-songs.component';
import { SongCardComponent } from './components/song-card/song-card.component';
import { SharedModule } from '../shared/shared.module';
import { SongProfileComponent } from './components/song-profile/song-profile.component';


@NgModule({
  declarations: [
    AllSongsComponent,
    SongCardComponent,
    SongProfileComponent
  ],
  imports: [
    CommonModule,
    SongRoutingModule,
    SharedModule
  ]
})
export class SongModule { }
