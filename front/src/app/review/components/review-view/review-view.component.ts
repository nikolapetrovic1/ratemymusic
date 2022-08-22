import { Component, Input, OnInit } from '@angular/core';
import { AuthService } from 'src/app/auth/service/auth.service';
import { RatingService } from 'src/app/song/service/rating.service';
import { CommentService } from '../../service/comment.service';
import { ReviewService } from '../../service/review.service';

@Component({
  selector: 'app-review-view',
  templateUrl: './review-view.component.html',
  styleUrls: ['./review-view.component.scss']
})
export class ReviewViewComponent implements OnInit {


  @Input() review!: any;
  comments: any = [];
  comment: any = '';
  constructor(private reviewService:ReviewService,private commentService:CommentService,private authService:AuthService) { }

  ngOnInit(): void {
    this.loadComments(this.review.id);
    this.loadUserId(this.review.user_id);
  }
  loadUserId(id:any){
    this.authService.loadUserInfo(id).subscribe((res)=>{
      this.review.user = res;
    })
  }
  loadComments(id:any){
    this.commentService.loadComments(id).subscribe((res) =>{
      if(res){
        this.comments = res.comments;
      }
    })
  }
  submit(){
    var comment = {
      review_id : this.review.id,
      comment : this.comment
    }
    this.commentService.submitComment(comment).subscribe((res)=>{
      this.comment = '';
      this.loadComments(this.review.id);
    })
  }

}
