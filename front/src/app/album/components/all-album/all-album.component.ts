import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AlbumService } from '../../service/album.service';

@Component({
  selector: 'app-all-album',
  templateUrl: './all-album.component.html',
  styleUrls: ['./all-album.component.scss']
})
export class AllAlbumComponent implements OnInit {
  query : string = '';
  albums: any  = [];
  constructor(private albumService:AlbumService,private router:ActivatedRoute) { }

  ngOnInit(): void {
    this.search();
  }

  search(){
    this.albumService.search(this.query).subscribe((res)=>{
      this.albums = res.body.albums;
    })
  }
}
