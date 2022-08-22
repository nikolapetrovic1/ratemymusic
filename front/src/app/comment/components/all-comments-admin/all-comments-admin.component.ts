import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { CommentService } from 'src/app/review/service/comment.service';

@Component({
  selector: 'app-all-comments-admin',
  templateUrl: './all-comments-admin.component.html',
  styleUrls: ['./all-comments-admin.component.scss']
})
export class AllCommentsAdminComponent implements OnInit {

  comments!: MatTableDataSource<any>;
  displayedColumns: string[] = ['id', 'comment',"user_id","review_id","reports","delete"];
  constructor(private commentService:CommentService) {
    this.loadComments();
   }

  ngOnInit(): void {
  }

  loadComments(){
    this.commentService.getAllByReport().subscribe((res)=>{
      this.comments = res.comments;
    })
  }
  deleteComment(id:any){
    this.commentService.deleteComment(id).subscribe((res)=>{
      console.log(res);
    })
  }

}
