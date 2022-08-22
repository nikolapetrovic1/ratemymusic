import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AlbumRoutingModule } from './album-routing.module';
import { AlbumProfileComponent } from './components/album-profile/album-profile.component';
import { AllAlbumComponent } from './components/all-album/all-album.component';
import { SharedModule } from '../shared/shared.module';
import { AlbumCardComponent } from './components/album-card/album-card.component';
import { CreateAlbumComponent } from './components/create-album/create-album.component';
import { ReviewModule } from '../review/review.module';


@NgModule({
  declarations: [
    AlbumProfileComponent,
    AllAlbumComponent,
    AlbumCardComponent,
    CreateAlbumComponent
  ],
  imports: [
    CommonModule,
    AlbumRoutingModule,
    SharedModule,
    ReviewModule
  ]
})
export class AlbumModule { }
