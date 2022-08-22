import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ReviewService } from '../../service/review.service';

@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.scss']
})
export class ReviewComponent implements OnInit {

  @Input() id: any;
  @Input() type: any;
  @Output() submitEvent = new EventEmitter<string>();
  text!: any;
  review!: any;
  constructor(private reviewService:ReviewService) { }

  ngOnInit(): void {
    this.getUserReview();

  }
  
  getUserReview(){
    this.reviewService.getUserReview(this.id,this.type).subscribe((res)=>{
      this.review = res;
      if(!this.review) {
        this.review.text = '';
      }
    })
    
  }

  submit(){
    var review = {
      id: this.review.id,
      text: this.review.text,
      type: this.type,
      type_id: this.id
    }
    this.reviewService.submitReview(review).subscribe((res)=>{
      console.log(review);
      this.submitEvent.emit();
    })
  }

}
