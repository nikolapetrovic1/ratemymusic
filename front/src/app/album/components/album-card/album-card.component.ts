import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { MusicianService } from 'src/app/musician/service/musician.service';
import { AlbumService } from '../../service/album.service';

@Component({
  selector: 'app-album-card',
  templateUrl: './album-card.component.html',
  styleUrls: ['./album-card.component.scss']
})
export class AlbumCardComponent implements OnInit {

  @Input() album!: any;
  musician!: any;
  constructor(private albumService:AlbumService,private musicianService:MusicianService,private _snackbar:MatSnackBar,private router:Router) { }

  ngOnInit(): void {
    // if(this.album.musicianID){
    //   this.loadMusicianInfo(this.album.musicianID)
    // }

  }

  loadMusicianInfo(id:any) {
    this.musicianService.getById(id).subscribe((res)=>{
      this.musician = res.data;
    })
  }
  albumPage(id:any){
    this.router.navigate([`/album/${id}`])
  }

}
