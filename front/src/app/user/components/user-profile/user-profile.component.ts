import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from 'src/app/auth/service/auth.service';
import { CommentService } from 'src/app/review/service/comment.service';
import { ReviewService } from 'src/app/review/service/review.service';
import { RatingService } from 'src/app/song/service/rating.service';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent implements OnInit {
  userId!: any;
  user: any;
  comments!: any;
  reviews!: any;
  ratings!: any;
  constructor(private route:ActivatedRoute,private authService:AuthService,private commentService:CommentService,
    private reviewService:ReviewService,
    private ratingService:RatingService) { 
    this.userId = JSON.parse(this.route.snapshot.paramMap.get('id')!);
    this.loadData(this.userId);
  }

  ngOnInit(): void {
  }

  loadData(id:any){
    this.loadUser(id);
    this.loadComments(id);
    this.loadReviews(id);
    this.loadRatings(id);
  }
  loadUser(id:any){
  this.authService.loadUserInfo(id).subscribe((res) => {
    this.user = res;
  })
  }
  loadComments(id:any){
    this.commentService.getByUser(id).subscribe((res)=>{
      this.comments = res.comments;
    })
  }
  loadReviews(id:any){
    this.reviewService.findByUser(id).subscribe((res)=>{
      this.reviews = res.reviews;
    })
  }
  loadRatings(id:any){
    this.ratingService.findByUser(id).subscribe((res)=>{
      this.ratings = res.rating_data;
    })
  }
}
