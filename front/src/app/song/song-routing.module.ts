import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllSongsComponent } from './components/all-songs/all-songs.component';
import { SongProfileComponent } from './components/song-profile/song-profile.component';

const routes: Routes = [

  {
    path: "",
    pathMatch: "full",
    component: AllSongsComponent,
  },
  {
    path: ":id",
    pathMatch: "full",
    component: SongProfileComponent,
  },

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SongRoutingModule { }
