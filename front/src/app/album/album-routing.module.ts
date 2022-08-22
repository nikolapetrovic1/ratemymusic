import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AlbumProfileComponent } from './components/album-profile/album-profile.component';
import { AllAlbumComponent } from './components/all-album/all-album.component';
import { CreateAlbumComponent } from './components/create-album/create-album.component';

const routes: Routes = [
  {
    path: "",
    pathMatch: "full",
    component: AllAlbumComponent,
  },
  {
    path: "create",
    pathMatch: "full",
    component: CreateAlbumComponent,
  },
  {
    path: ":id",
    pathMatch: "full",
    component: AlbumProfileComponent,
  },

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AlbumRoutingModule { }
