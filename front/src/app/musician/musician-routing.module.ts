import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllMusiciansComponent } from './components/all-musicians/all-musicians.component';
import { CreateMusicianComponent } from './components/create-musician/create-musician.component';
import { MusicianProfileComponent } from './components/musician-profile/musician-profile.component';

const routes: Routes = [
  {
    path: "",
    pathMatch: "full",
    component: AllMusiciansComponent,
  },
  {
    path: "create",
    pathMatch: "full",
    component: CreateMusicianComponent,
  },
  {
    path: ":id",
    pathMatch: "full",
    component: MusicianProfileComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MusicianRoutingModule { }
