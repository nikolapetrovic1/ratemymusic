import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LayoutComponent } from './routing/layout/layout/layout.component';
import { SharedModule } from './shared/shared.module';

const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    loadChildren: () => import('./auth/auth.module').then((m) => m.AuthModule),
  },
  {
    path: 'musician',
    component: LayoutComponent,
    loadChildren: () => import('./musician/musician.module').then((m) => m.MusicianModule),
  },
  {
    path: 'song',
    component: LayoutComponent,
    loadChildren: () => import('./song/song.module').then((m) => m.SongModule),
  },
  {
    path: 'album',
    component: LayoutComponent,
    loadChildren: () => import('./album/album.module').then((m) => m.AlbumModule),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes),SharedModule],
  exports: [RouterModule]
})
export class AppRoutingModule { }
