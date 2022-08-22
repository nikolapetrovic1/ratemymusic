import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SongRoutingModule } from './song-routing.module';
import { AllSongsComponent } from './components/all-songs/all-songs.component';
import { SongCardComponent } from './components/song-card/song-card.component';
import { SharedModule } from '../shared/shared.module';
import { SongProfileComponent } from './components/song-profile/song-profile.component';
import { NgbRatingModule } from '@ng-bootstrap/ng-bootstrap';
import { ReviewModule } from '../review/review.module';
import { CreateSongComponent } from './components/create-song/create-song.component';
@NgModule({
  declarations: [
    AllSongsComponent,
    SongCardComponent,
    SongProfileComponent,
    CreateSongComponent,
    
  ],
  imports: [
    CommonModule,
    SongRoutingModule,
    SharedModule,
    NgbRatingModule,
    ReviewModule
    
  ]
})
export class SongModule { }
