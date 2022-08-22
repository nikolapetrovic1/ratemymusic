import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { MusicianService } from 'src/app/musician/service/musician.service';
import { RatingService } from 'src/app/song/service/rating.service';
import { AlbumService } from '../../service/album.service';

@Component({
  selector: 'app-album-profile',
  templateUrl: './album-profile.component.html',
  styleUrls: ['./album-profile.component.scss']
})
export class AlbumProfileComponent implements OnInit {

  query!: string;
  album!: any;
  songs!: any;
  albumId!:any;
  musician!:any;
  ratings!:any;
  ratingValue = new FormControl('');
  averageRating! : number;
  constructor(private albumService:AlbumService,
    private router:ActivatedRoute,
    private musicianService:MusicianService,
    private ratingService:RatingService) { 

  }

  ngOnInit(): void {
    this.albumId =JSON.parse(this.router.snapshot.paramMap.get('id')!);
    this.getSongs(this.albumId);
    this.loadRatings();
  }
  getAlbum(id:any){
    this.albumService.getAlbum(id).subscribe((res)=>{
      this.album = res.data;
      this.getMusician(this.album.musicianID)
    })
  }
  getSongs(id:any){
      this.albumService.getSongs(id).subscribe((res)=>{
        this.songs = res.songs;
      })
  }
  getMusician(id:any){
    this.musicianService.getById(id).subscribe((res)=>{
      this.musician = res.data
    })
  }

  rateAlbum(){
    var rating : any = {
      rating_value : this.ratingValue.value,
      album_id : this.album.id,
    }
    this.ratingService.rate(rating,"album").subscribe((res)=>{
      this.loadRatings();
    })
  }
  calculateAverage(){
    this.averageRating = 0;
    this.ratings.forEach((rating: { rating_value: number; }) => { this.averageRating += rating.rating_value });
    this.averageRating = this.averageRating/this.ratings.length
  }
  getUserRating(id:string){
    this.
    ratingService.getUserRating(id,'album').subscribe((res)=>{
      this.ratingValue.setValue(res.rating_value);
    })
  }
  getAllAlbumRatings(id:string){
    this.ratingService.getAllSongRatings(id,'album').subscribe((res)=>{
      this.ratings = res.rating_data;
      if(this.ratings){
        this.calculateAverage();
      }

    })
  }
  loadRatings(){
    this.getAlbum(this.albumId)
    this.getUserRating(this.albumId);
    this.getAllAlbumRatings(this.albumId);
  }

}
