import { Component, Input, OnInit } from '@angular/core';
import { SongService } from '../../service/song.service';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
@Component({
  selector: 'app-song-player',
  templateUrl: './song-player.component.html',
  styleUrls: ['./song-player.component.scss']
})
export class SongPlayerComponent implements OnInit {

  @Input() songId: any;
  song!:SafeResourceUrl;
  constructor(private songService: SongService,private domSanitizer: DomSanitizer ) {
  }

  
  ngOnInit(): void {
    this.songService.streamSong(this.songId).subscribe((res)=>{
       this.song = this.domSanitizer.bypassSecurityTrustUrl("data:audio/mpeg;base64," + res);
    })
  }


}
