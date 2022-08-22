import { Component, OnInit } from '@angular/core';
import { SongService } from '../../service/song.service';

@Component({
  selector: 'app-all-songs',
  templateUrl: './all-songs.component.html',
  styleUrls: ['./all-songs.component.scss']
})
export class AllSongsComponent implements OnInit {

  query: string = '';
  songs: any = [];
  currentRate = 1;
  constructor(private songService:SongService) { }

  ngOnInit(): void {
    this.search();
  }

  search(){
    this.songService.search(this.query).subscribe((res)=>{
      this.songs = res.body.songs;
    })
  }
}
