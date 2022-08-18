import { Component, OnInit } from '@angular/core';
import { SongService } from '../../service/song.service';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-song-profile',
  templateUrl: './song-profile.component.html',
  styleUrls: ['./song-profile.component.scss']
})
export class SongProfileComponent implements OnInit {
  song!: any;
  constructor(private router: ActivatedRoute,private songService:SongService) { }

  ngOnInit(): void {
    this.getSongId(JSON.parse(this.router.snapshot.paramMap.get('id')!))
  }

  getSongId(id:string){
    this.songService.get(id).subscribe((res)=>{
        res.body.song = this.song;
    })
  }

}
