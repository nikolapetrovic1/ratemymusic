import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CommentRoutingModule } from './comment-routing.module';
import { AllCommentsAdminComponent } from './components/all-comments-admin/all-comments-admin.component';
import { SharedModule } from '../shared/shared.module';


@NgModule({
  declarations: [
    AllCommentsAdminComponent
  ],
  imports: [
    CommonModule,
    CommentRoutingModule,
    SharedModule
  ]
})
export class CommentModule { }
