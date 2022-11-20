import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MusicianService } from 'src/app/musician/service/musician.service';

@Component({
  selector: 'app-song-card',
  templateUrl: './song-card.component.html',
  styleUrls: ['./song-card.component.scss']
})
export class SongCardComponent implements OnInit {

  @Input() song: any;
  musician: any;
  constructor(private router:Router,private musicianService: MusicianService) { }

  ngOnInit(): void {
    this.musicianService.getById(this.song.musicianID).subscribe((res)=>{
      this.musician = res.data;
    })
  }
  songPage(id:any){
    this.router.navigate([`/song/${id}`]);
  }
}
