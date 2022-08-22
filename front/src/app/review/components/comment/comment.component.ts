import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AuthService } from 'src/app/auth/service/auth.service';
import { CommentService } from '../../service/comment.service';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.scss']
})
export class CommentComponent implements OnInit {

  @Input() comment!: any;

  constructor(private commentService:CommentService,private authService:AuthService,private _snackBar:MatSnackBar) { }

  ngOnInit(): void {
    this.getUser(this.comment.user_id);
  }
  report(){
    this.commentService.reportComment(this.comment.id).subscribe((res)=>{
      this.comment.user = res.email;
      this._snackBar.open('Comment reported!', 'Close', {
        duration: 3000
      });
    })
  }
  getUser(id:any){
    this.authService.loadUserInfo(id).subscribe((res)=>{
      this.comment.user = res.email;
    })
  } 
}
