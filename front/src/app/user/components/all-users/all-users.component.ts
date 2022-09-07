import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/auth/service/auth.service';

@Component({
  selector: 'app-all-users',
  templateUrl: './all-users.component.html',
  styleUrls: ['./all-users.component.scss']
})
export class AllUsersComponent implements OnInit {


  users!: any;
  displayedColumns: string[] = ['id', 'name',"ban"];
  constructor(private authService:AuthService) { }

  ngOnInit(): void {
    this.loadUsers();
  }
  loadUsers(){
    this.authService.all().subscribe((res)=>{
      this.users = res.users;
    })
  }
  banUser(id:any){
    this.authService.banUser(id).subscribe((res)=>{
      this.users = res.users;
    })
  }
  unBanUser(id:any){
    this.authService.unBanUser(id).subscribe((res)=>{
      this.users = res.users;
    })
  }
  deleteUser(id:any){
    this.authService.delete(id).subscribe((res) =>{
      this.loadUsers();
    })
  }
}
