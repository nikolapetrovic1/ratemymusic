import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllMusiciansComponent } from './components/all-musicians/all-musicians.component';

const routes: Routes = [
  {
    path: "",
    pathMatch: "full",
    component: AllMusiciansComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MusicianRoutingModule { }
