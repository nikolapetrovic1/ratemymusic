import { Component, Input, OnInit } from '@angular/core';
import { ReviewService } from '../../service/review.service';

@Component({
  selector: 'app-all-reviews',
  templateUrl: './all-reviews.component.html',
  styleUrls: ['./all-reviews.component.scss']
})
export class AllReviewsComponent implements OnInit {


  @Input() id:any;
  @Input() type!:string;
  reviews!:any;
  constructor(private reviewService:ReviewService) { }

  ngOnInit(): void {
    this.loadReviews();
  }

  loadReviews(){
    this.reviewService.getAllReview(this.id,this.type).subscribe((res)=>{
      this.reviews = res.reviews;
    })
  }

}
