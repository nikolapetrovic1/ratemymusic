import { Component, Input, OnInit } from '@angular/core';
import { SongService } from '../../service/song.service';

@Component({
  selector: 'app-song-player',
  templateUrl: './song-player.component.html',
  styleUrls: ['./song-player.component.scss']
})
export class SongPlayerComponent implements OnInit {

  @Input() songId: any;
  constructor(private songService: SongService) {
  }

  
  ngOnInit(): void {
  }


}
