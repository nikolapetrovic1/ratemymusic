import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ReviewRoutingModule } from './review-routing.module';
import { AllReviewsComponent } from './components/all-reviews/all-reviews.component';
import { ReviewComponent } from './components/review/review.component';
import { SharedModule } from '../shared/shared.module';
import { ReviewViewComponent } from './components/review-view/review-view.component';
import { CommentComponent } from './components/comment/comment.component';


@NgModule({
  declarations: [
    AllReviewsComponent,
    ReviewComponent,
    ReviewViewComponent,
    CommentComponent
  ],
  imports: [
    CommonModule,
    ReviewRoutingModule,
    SharedModule
  ],
  exports:[AllReviewsComponent,ReviewComponent]
})
export class ReviewModule { }
