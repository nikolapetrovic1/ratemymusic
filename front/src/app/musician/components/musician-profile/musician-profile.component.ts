import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AlbumService } from 'src/app/album/service/album.service';
import { SongService } from 'src/app/song/service/song.service';
import { MusicianService } from '../../service/musician.service';

@Component({
  selector: 'app-musician-profile',
  templateUrl: './musician-profile.component.html',
  styleUrls: ['./musician-profile.component.scss']
})
export class MusicianProfileComponent implements OnInit {


  musician!: any;
  musicianId!: any;
  songs!: any;
  albums!: any;
  constructor(private musicianService:MusicianService,private router:ActivatedRoute,private songService:SongService,private albumService:AlbumService) { }

  ngOnInit(): void {
    this.loadMusician();
  }

  loadMusician(){
    this.musicianId = JSON.parse(this.router.snapshot.paramMap.get('id')!);
    this.musicianService.getById(this.musicianId).subscribe((res)=>{
      this.musician = res.data;
      this.loadSongs(this.musician.id)
      this.loadAlbums(this.musician.id)
    });
  }
  loadSongs(id:any){
    this.songService.getByMusician(id).subscribe((res)=>{
      this.songs = res.songs
    })
  }
  loadAlbums(id:any){
    this.albumService.loadByMusician(id).subscribe((res)=>{
      this.albums = res.albums
    })
  }
}
