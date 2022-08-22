import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllCommentsAdminComponent } from './components/all-comments-admin/all-comments-admin.component';

const routes: Routes = [

  {
    path: "all",
    pathMatch: "full",
    component: AllCommentsAdminComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CommentRoutingModule { }
